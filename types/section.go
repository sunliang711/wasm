package types

import "fmt"

type RawSecType byte

const (
	//section type (0x00-0x0b,0x7f)
	SectionUser RawSecType = iota
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

	SectionExceptionTypes RawSecType = 0x7f
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
	sectionType2Order = map[RawSecType]byte{
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

	orderSectionStr = map[byte]string{
		OrderUser:                 "user section",
		OrderType:                 "type section",
		OrderImport:               "import section",
		OrderFunctionDeclarations: "function declarations section",
		OrderTable:                "table section",
		OrderMemory:               "memory section",
		OrderGlobal:               "global section",
		OrderExport:               "export section",
		OrderStart:                "start section",
		OrderElem:                 "elem section",
		OrderFunctionDefinitions:  "function definitions section",
		OrderData:                 "data section",
		OrderExceptionTypes:       "exception types section",
	}
)

func SectionType2Order(rawSec RawSecType) (byte, error) {
	orderSec, ok := sectionType2Order[rawSec]
	if !ok {
		return OrderUnknown, fmt.Errorf(ErrUnknownSection, rawSec)
	}
	return orderSec, nil
}

func OrderSectionString(orderSec byte) string {
	return orderSectionStr[orderSec]
}
