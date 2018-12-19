package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"sync"
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
	Module    *types.Module
	Closed    bool

	typeParsed chan struct{}
}

type Section struct {
	Type            byte
	NumSectionBytes uint32
	Data            []byte
}

func (sec Section) String() string {
	return fmt.Sprintf("{Type: %d,NumSectionBytes: %d,Data: %v}", sec.Type, sec.NumSectionBytes, sec.Data)
}

func New(filename string) (*Parser, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Parser{
		Stream:     bytes.NewReader(content),
		ChSection:  make(chan *Section),
		ChErr:      make(chan error),
		ChQuit:     make(chan struct{}),
		ChDone:     make(chan struct{}),
		Wg:         new(sync.WaitGroup),
		Module:     new(types.Module),
		Closed:     false,
		typeParsed: make(chan struct{}, 2),
	}, nil
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

	lastSectionType := types.OrderUnknown

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
		rawSectionType := types.RawSecType(bufType[0])
		orderSection, err := types.SectionType2Order(rawSectionType)
		if err != nil {
			p.NotifyError(err)
			break
		}

		//check section order
		if orderSection != types.OrderUser {
			if orderSection > lastSectionType {
				lastSectionType = orderSection
			} else {
				p.NotifyError(fmt.Errorf(types.ErrIncorrectOrder))
				break
			}
		}
		// get section num bytes
		var sectionNumBytes uint32
		err = utils.DecodeVarInt(p.Stream, 32, &sectionNumBytes)
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
			logrus.Infof("Parse done.")
			return nil
		}
	}
}

func (p *Parser) parseSection(sec *Section) {
	var (
		err error
	)
	defer p.Wg.Done()
	switch sec.Type {
	case types.OrderType:
		err = p.typeSection(sec)
	case types.OrderImport:
		err = p.importSection(sec)
	case types.OrderFunctionDeclarations:
		err = p.functionDeclarationsSection(sec)
	case types.OrderTable:
		err = p.tableSection(sec)
	case types.OrderMemory:
		err = p.memorySection(sec)
	case types.OrderGlobal:
		err = p.globalSection(sec)
	case types.OrderExceptionTypes:
		err = p.exceptionTypesSection(sec)
	case types.OrderExport:
		err = p.exportSection(sec)
	case types.OrderStart:
		err = p.startSection(sec)
	case types.OrderElem:
		err = p.elemSection(sec)
	case types.OrderFunctionDefinitions:
		err = p.functionDefinitionsSection(sec)
	case types.OrderData:
		err = p.dataSection(sec)
	case types.OrderUser:
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
