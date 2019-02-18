package parser

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"sync"
	"time"
	"wasm/core/IR"
	"wasm/types"
	"wasm/utils"
)

var (
	MagicNumber    = []byte{0x00, 0x61, 0x73, 0x6d}
	CurrentVersion = []byte{0x01, 0x00, 0x00, 0x00}
)

type Parser struct {
	Stream    io.Reader
	ChSection chan *Section
	ChErr     chan error
	ChQuit    chan struct{}
	ChDone    chan struct{}
	Wg        *sync.WaitGroup
	Module    *IR.Module
	Closed    bool

	*IR.DeferredCodeValidationState

	typeParsed            bool
	funcDeclarationParsed bool
}

type Section struct {
	Type            byte
	NumSectionBytes uint32
	Data            []byte
}

func (sec Section) String() string {
	return fmt.Sprintf("{Type: %d,NumSectionBytes: %d,Data: %v}", sec.Type, sec.NumSectionBytes, sec.Data)
}

func NewParser(stream io.Reader) (*Parser, error) {
	return &Parser{
		Stream:                      stream,
		ChSection:                   make(chan *Section),
		ChErr:                       make(chan error),
		ChQuit:                      make(chan struct{}),
		ChDone:                      make(chan struct{}),
		Wg:                          new(sync.WaitGroup),
		Module:                      IR.NewModule(),
		Closed:                      false,
		typeParsed:                  false,
		funcDeclarationParsed:       false,
		DeferredCodeValidationState: new(IR.DeferredCodeValidationState),
	}, nil
}

func (p *Parser) Parse2() (err error) {

	var (
		bufType = make([]byte, 1)
		bufData []byte
	)
	// check magic number
	err = checkConstant(p.Stream, MagicNumber, types.ErrMagicNumber)
	if err != nil {
		return err
	}
	// check version
	err = checkConstant(p.Stream, CurrentVersion, types.ErrVersion)
	if err != nil {
		return err
	}

	lastSectionType := IR.OrderUnknown

	// read loop
	// get section bytes,then send to loop()
	for {
		// get section type,if read 0 byte, more => false
		_, err := p.Stream.Read(bufType)
		if err == io.EOF {
			logrus.Info("End of file")
			break
		} else if err != nil {
			return err
		}
		rawSectionType := IR.RawSecType(bufType[0])
		orderSection, err := IR.SectionType2Order(rawSectionType)
		if err != nil {
			return nil
		}

		//check section order
		if orderSection != IR.OrderUser {
			if orderSection > lastSectionType {
				lastSectionType = orderSection
			} else {
				return fmt.Errorf(types.ErrIncorrectOrder)
			}
		}
		// get section num bytes
		var sectionNumBytes uint32
		_, err = utils.DecodeVarInt(p.Stream, 32, &sectionNumBytes)
		if err != nil {
			return err
		}

		// get section data
		bufData = make([]byte, sectionNumBytes)
		n, err := p.Stream.Read(bufData)
		if err != nil {
			return err
		}
		if uint32(n) != sectionNumBytes {
			return fmt.Errorf(types.ErrInsufficientBytes)
		}

		//make Section
		section := &Section{
			Type:            orderSection,
			NumSectionBytes: sectionNumBytes,
			Data:            bufData,
		}
		err = p.parseSection2(section)
		if err != nil {
			return err
		}
	}
	p.Post()
	return nil
}

func (p *Parser) Parse() error {
	go p.fileLoop()
	return p.eventLoop()
}

func (p *Parser) fileLoop() {
	var (
		bufType = make([]byte, 1)
		bufData []byte
	)
	// check magic number
	err := checkConstant(p.Stream, MagicNumber, types.ErrMagicNumber)
	if err != nil {
		p.NotifyError(err)
		return
	}
	// check version
	err = checkConstant(p.Stream, CurrentVersion, types.ErrVersion)
	if err != nil {
		p.NotifyError(err)
		return
	}

	lastSectionType := IR.OrderUnknown

	// read loop
	// get section bytes,then send to loop()
	for {
		// get section type,if read 0 byte, more => false
		_, err := p.Stream.Read(bufType)
		if err == io.EOF {
			logrus.Info("fileLoop(): End of file")
			close(p.ChDone)
			break
		} else if err != nil {
			p.NotifyError(err)
			break
		}
		rawSectionType := IR.RawSecType(bufType[0])
		orderSection, err := IR.SectionType2Order(rawSectionType)
		if err != nil {
			p.NotifyError(err)
			break
		}

		//check section order
		if orderSection != IR.OrderUser {
			if orderSection > lastSectionType {
				lastSectionType = orderSection
			} else {
				p.NotifyError(fmt.Errorf(types.ErrIncorrectOrder))
				break
			}
		}
		// get section num bytes
		var sectionNumBytes uint32
		_, err = utils.DecodeVarInt(p.Stream, 32, &sectionNumBytes)
		if err != nil {
			p.NotifyError(err)
			break
		}

		// get section data
		bufData = make([]byte, sectionNumBytes)
		n, err := p.Stream.Read(bufData)
		if err != nil {
			p.NotifyError(err)
			break
		}
		if uint32(n) != sectionNumBytes {
			p.NotifyError(fmt.Errorf(types.ErrInsufficientBytes))
			break
		}

		//make Section
		section := &Section{
			Type:            orderSection,
			NumSectionBytes: sectionNumBytes,
			Data:            bufData,
		}
		logrus.Infof("fileLoop(): Found new section: %v", section)
		p.Wg.Add(1)
		p.ChSection <- section
	}

}

func (p *Parser) eventLoop() error {
	err := fmt.Errorf("quit")
	for {
		select {
		case err = <-p.ChErr:
			logrus.Errorf("eventLoop(): error: %s", err.Error())
			p.Stop()

		case <-p.ChQuit:
			logrus.Infof("eventLoop(): quit.")
			return err

		case section := <-p.ChSection:
			logrus.Infof("eventLoop(): Got section: %v", section)
			go p.parseSection(section)

		case <-p.ChDone:
			p.Wg.Wait()
			p.Post()
			logrus.Infof("Parse done.")
			//TODO
			//p.validateDataSegments()
			return nil
		}
	}
}

func (p *Parser) parseSection2(sec *Section) (err error) {
	switch sec.Type {
	case IR.OrderType:
		err = p.typeSection(sec)
	case IR.OrderImport:
		err = p.importSection(sec)
	case IR.OrderFunctionDeclarations:
		err = p.functionDeclarationsSection(sec)
	case IR.OrderTable:
		err = p.tableSection(sec)
	case IR.OrderMemory:
		err = p.memorySection(sec)
	case IR.OrderGlobal:
		err = p.globalSection(sec)
	case IR.OrderExceptionTypes:
		err = p.exceptionTypesSection(sec)
	case IR.OrderExport:
		err = p.exportSection(sec)
	case IR.OrderStart:
		err = p.startSection(sec)
	case IR.OrderElem:
		err = p.elemSection(sec)
	case IR.OrderFunctionDefinitions:
		err = p.functionDefinitionsSection(sec)
	case IR.OrderData:
		err = p.dataSection(sec)
	case IR.OrderUser:
		err = p.userSection(sec)
	}
	return
}

func (p *Parser) parseSection(sec *Section) {
	var (
		err error
	)
	defer p.Wg.Done()
	switch sec.Type {
	case IR.OrderType:
		err = p.typeSection(sec)
	case IR.OrderImport:
		err = p.importSection(sec)
	case IR.OrderFunctionDeclarations:
		err = p.functionDeclarationsSection(sec)
	case IR.OrderTable:
		err = p.tableSection(sec)
	case IR.OrderMemory:
		err = p.memorySection(sec)
	case IR.OrderGlobal:
		err = p.globalSection(sec)
	case IR.OrderExceptionTypes:
		err = p.exceptionTypesSection(sec)
	case IR.OrderExport:
		err = p.exportSection(sec)
	case IR.OrderStart:
		err = p.startSection(sec)
	case IR.OrderElem:
		err = p.elemSection(sec)
	case IR.OrderFunctionDefinitions:
		err = p.functionDefinitionsSection(sec)
	case IR.OrderData:
		err = p.dataSection(sec)
	case IR.OrderUser:
		err = p.userSection(sec)
	}
	if err != nil {
		p.NotifyError(err)
	}
}

func (p *Parser) Stop() {
	if !p.Closed {
		close(p.ChQuit)
		p.Closed = true
	}
}

func (p *Parser) NotifyError(err error) {
	p.ChErr <- err
}

func (p *Parser) Post() {
	importFunLen := uint64(len(p.Module.Functions.Imports))
	for _, e := range p.Module.ExportFunctions {
		if e.Index >= importFunLen {
			p.Module.Functions.Defs[e.Index-importFunLen].Name = e.Name
		}
	}
}

func waitCondition(condition *bool, timeoutMsg string, timeout int) error {
	sleepCount := 0
	for !*condition {
		time.Sleep(time.Millisecond * 100)
		sleepCount += 1
		if sleepCount == timeout/100 {
			return fmt.Errorf(timeoutMsg)
		}
	}
	return nil
}
