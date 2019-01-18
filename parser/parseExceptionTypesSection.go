package parser

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"wasm/core/IR"
	"wasm/utils"
)

func (p *Parser) exceptionTypesSection(sec *Section) error {
	err := checkSection(sec, IR.OrderExceptionTypes)
	if err != nil {
		return err
	}
	rd := bytes.NewReader(sec.Data)

	//1. num exceptionType
	var numExcep uint32
	_, err = utils.DecodeVarInt(rd, 32, &numExcep)
	if err != nil {
		return err
	}

	//2. exceptionTypes
	for i := 0; i < int(numExcep); i++ {
		params := IR.TypeTuple{}
		err = DecodeTypeTuple(rd, &params)
		if err != nil {
			return err
		}

		excType := IR.ExceptionType{Params: params}
		excDef := IR.ExceptionTypeDef{Type: excType}
		p.Module.ExceptionTypes.Defs = append(p.Module.ExceptionTypes.Defs, excDef)
		logrus.Infof("<exceptiontype section> exection def: %v", excDef)
	}

	err = p.validateExceptionTypes()
	return err
}

func (p *Parser) validateExceptionTypes() error {
	//TODO
	logrus.Info("TODO: validateExceptionTypes()")
	return nil
}
