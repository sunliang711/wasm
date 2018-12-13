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

const (
	ErrReadCount       = "Read count error"
	ErrReadSectionType = "Read section type error"
	ErrIncorrectOrder  = "Incorrect order for known section"

	ErrMagicNumber = "Magic number error"
	ErrVersion     = "Version number error"

	ErrInsufficientBytes = "Got insufficient bytes"

	ErrUnknownSection = "Error unknown section"
)

var (
	MagicNumber    = []byte{0x00, 0x61, 0x73, 0x6d}
	CurrentVersion = []byte{0x01, 0x00, 0x00, 0x00}
)

const (
	//section type (0x00-0x0b,0x7f)
	SectionUser = iota
	SectionType
	SectionImport
	SectionFunctionDeclarations
	SectionTable
	SectionMemory
	SectionGlobal
	SectionExport
	SectionStart
	SectionElem
	SectionFunctionDefinitions
	SectionData

	SectionExceptionTypes byte = 0x7f
)

const (
	//section type order
	OrderUnknown byte = iota
	OrderType
	OrderImport
	OrderFunctionDeclarations
	OrderTable
	OrderMemory
	OrderGlobal
	OrderExceptionTypes
	OrderExport
	OrderStart
	OrderElem
	OrderFunctionDefinitions
	OrderData
	OrderUser
)

var (
	sectionType2Order = map[byte]byte{
		SectionUser:                 OrderUser,
		SectionType:                 OrderType,
		SectionImport:               OrderImport,
		SectionFunctionDeclarations: OrderFunctionDeclarations,
		SectionTable:                OrderTable,
		SectionMemory:               OrderMemory,
		SectionGlobal:               OrderGlobal,
		SectionExport:               OrderExport,
		SectionStart:                OrderStart,
		SectionElem:                 OrderElem,
		SectionFunctionDefinitions:  OrderFunctionDefinitions,
		SectionData:                 OrderData,
		SectionExceptionTypes:       OrderExceptionTypes,
	}
)

type decodeBody func(rd io.Reader, numbytes uint32) error

type Parser struct {
	Stream    io.Reader
	ChSection chan *Section
	ChErr     chan error
	ChQuit    chan struct{}
	ChDone    chan struct{}
	Wg        *sync.WaitGroup
	Module    *Module
	Closed    bool
}
type Module struct {
	Types []types.FunctionType

	//TODO
	//[]types.Exports
	//[]types.Data
	//[]elementSegments
	//[]userSections

	StartFunctionIndex int
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
		Stream:    bytes.NewReader(content),
		ChSection: make(chan *Section),
		ChErr:     make(chan error),
		ChQuit:    make(chan struct{}),
		ChDone:    make(chan struct{}),
		Wg:        new(sync.WaitGroup),
		Module:    new(Module),
		Closed:    false,
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
	err := CheckConstant(p.Stream, MagicNumber, ErrMagicNumber)
	if err != nil {
		p.ChErr <- err
		return
	}
	// check version
	err = CheckConstant(p.Stream, CurrentVersion, ErrVersion)
	if err != nil {
		p.ChErr <- err
		return
	}

	lastSectionType := OrderUnknown

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
			p.ChErr <- err
			break
		}
		rawSectionType := bufType[0]
		orderSection := sectionType2Order[rawSectionType]
		if orderSection == OrderUnknown {
			p.ChErr <- fmt.Errorf(ErrUnknownSection)
			break
		}

		//check section order
		if orderSection != OrderUser {
			if orderSection > lastSectionType {
				lastSectionType = orderSection
			} else {
				p.ChErr <- fmt.Errorf(ErrIncorrectOrder)
				break
			}
		}
		// get section num bytes
		sectionNumBytes, err := utils.DecodeUInt32(p.Stream)
		if err != nil {
			p.ChErr <- err
			break
		}

		// get section data
		bufData = make([]byte, sectionNumBytes)
		n, err := p.Stream.Read(bufData)
		if err != nil {
			p.ChErr <- err
			break
		}
		if uint32(n) != sectionNumBytes {
			p.ChErr <- fmt.Errorf(ErrInsufficientBytes)
			break
		}

		//make Section
		section := &Section{
			Type:            orderSection,
			NumSectionBytes: sectionNumBytes,
			Data:            bufData,
		}
		logrus.Infof("fileLoop(): Found new section: %v",section)
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
	case OrderType:
		err = p.typeSection(sec)
	case OrderImport:
		err = p.importSection(sec)
	case OrderFunctionDeclarations:
		err = p.functionDeclarationsSection(sec)
	case OrderTable:
		err = p.tableSection(sec)
	case OrderMemory:
		err = p.memorySection(sec)
	case OrderGlobal:
		err = p.globalSection(sec)
	case OrderExceptionTypes:
		err = p.exceptionTypesSection(sec)
	case OrderExport:
		err = p.exportSection(sec)
	case OrderStart:
		err = p.startSection(sec)
	case OrderElem:
		err = p.elemSection(sec)
	case OrderFunctionDefinitions:
		err = p.functionDefinitionsSection(sec)
	case OrderData:
		err = p.dataSection(sec)
	case OrderUser:
		err = p.userSection(sec)
	}
	if err != nil {
		p.ChErr <- err
	}
}

func (p *Parser) Stop() {
	if !p.Closed {
		close(p.ChQuit)
		p.Closed = true
	}
}

func CheckConstant(rd io.Reader, constant []byte, errMsg string) error {
	numBytes := len(constant)
	buf := make([]byte, numBytes)
	nRead, err := rd.Read(buf)
	if err != nil {
		return err
	}
	if nRead != numBytes {
		return fmt.Errorf(ErrReadCount)
	}
	for i := 0; i < numBytes; i++ {
		if buf[i] != constant[i] {
			return fmt.Errorf(errMsg)
		}
	}
	return nil
}
