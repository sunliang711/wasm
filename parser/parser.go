package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"sync"
	"time"
	"wasm/utils"
)

const (
	ErrReadCount       = "Read count error"
	ErrReadSectionType = "Read section type error"
	ErrIncorrectOrder  = "Incorrect order for known section"

	ErrMagicNumber = "Magic number error"
	ErrVersion     = "Version number error"

	ErrInsufficientBytes = "Got insufficient bytes"
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
	Closed bool
}

type Section struct {
	Type            byte
	NumSectionBytes uint32
	Data            []byte
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
		Closed:false,
	}, nil
}

func (p *Parser) Parse() error{
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
			logrus.Info("EOF")
			close(p.ChDone)
			break
		} else if err != nil {
			p.ChErr <- err
			break
		}
		rawSectionType := bufType[0]
		logrus.Infof("section type: %d", rawSectionType)
		orderSection := sectionType2Order[rawSectionType]

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
		logrus.Infof("section num bytes: %d", sectionNumBytes)

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
		p.Wg.Add(1)
		p.ChSection <- section
	}

}

func (p *Parser) eventLoop() error{
	err := fmt.Errorf("quit")
	for {
		select {
		case err = <-p.ChErr:
			logrus.Errorf("loop(): error: %s", err.Error())
			p.Stop()

		case <-p.ChQuit:
			logrus.Infof("loop(): quit.")
			return err

		case section := <-p.ChSection:
			logrus.Infof("loop(): got section: %v", section)
			go p.parseSection(section)

		case <-p.ChDone:
			p.Wg.Wait()
			logrus.Infof("done.")
			return nil
		}
	}
}

func (p *Parser) parseSection(sec *Section) {
	defer p.Wg.Done()
	switch sec.Type {
	case OrderType:
		logrus.Info("parseSection(): type section")
	case OrderImport:
		logrus.Info("parseSection(): import section")
	case OrderFunctionDeclarations:
		logrus.Info("parseSection(): function delcarations section")
	case OrderTable:
		logrus.Info("parseSection(): table section")
	case OrderMemory:
		logrus.Info("parseSection(): memory section")
	case OrderGlobal:
		logrus.Info("parseSection(): global section")
	case OrderExceptionTypes:
		logrus.Info("parseSection(): exception types section")
	case OrderExport:
		logrus.Info("parseSection(): export section")
	case OrderStart:
		logrus.Info("parseSection(): start section")
	case OrderElem:
		logrus.Info("parseSection(): elem section")
	case OrderFunctionDefinitions:
		logrus.Info("parseSection(): function definition section")
	case OrderData:
		logrus.Info("parseSection(): data section")
	case OrderUser:
		logrus.Infof("begin parseSection(): user section %v", sec.NumSectionBytes)
		time.Sleep(time.Second * 1)
		logrus.Infof("end parseSection(): user section %v", sec.NumSectionBytes)
	}
	p.ChErr <- fmt.Errorf("parse section error")
}

func (p *Parser) Stop() {
	if !p.Closed{
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

func loadSection(rd io.Reader, decodebody decodeBody) error {

	//step 1
	sectionBytes, err := utils.DecodeUInt32(rd)
	if err != nil {
		return err
	}
	//step 2
	decodebody(rd, sectionBytes)

	return nil
}

//func LoadSections(rd io.Reader) error {
//	var (
//		buf             []byte
//		lastSectionType byte
//	)
//	buf = make([]byte, 1)
//
//	for {
//		_, err := rd.Read(buf)
//		if err != nil {
//			return fmt.Errorf(ErrReadSectionType)
//		}
//		lastSectionType = st_unknown
//		sectionType := buf[0]
//		if sectionType != st_user {
//			if sectionType > lastSectionType {
//				lastSectionType = sectionType
//			} else {
//				return fmt.Errorf(ErrIncorrectOrder)
//			}
//		}
//		switch sectionType {
//		case st_type:
//		case st_import:
//		case st_functionDeclarations:
//		case st_table:
//		case st_memory:
//		case st_global:
//		case st_exceptionTypes:
//		case st_export_:
//		case st_start:
//		case st_elem:
//		case st_functionDefinitions:
//		case st_data:
//		case st_user:
//		}
//
//	}
//}
