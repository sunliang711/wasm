package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"wasm/types"
	"wasm/utils"
)

const (
	ErrNotTypeSection = "Not a type section"
	ErrTypeSectionNum = "Type section number valid"

	ErrFunctionTag = "Not a function tag"
)

var (
	functionTag = []byte{0x60}
)

func (p *Parser) typeSection(sec *Section) error {
	if sec.Type != OrderType {
		return fmt.Errorf(ErrNotTypeSection)
	}

	if len(sec.Data) != int(sec.NumSectionBytes) {
		return fmt.Errorf(ErrTypeSectionNum)
	}

	var (
		functionType types.FunctionType
	)

	rd := bytes.NewReader(sec.Data)
	// type section is an array
	// get array size
	size, err := utils.DecodeUInt32(rd)
	if err != nil {
		return err
	}
	for index := 0; index < int(size); index++ {
		// 1. functionTag
		err = CheckConstant(rd, functionTag, ErrFunctionTag)
		if err != nil {
			return err
		}

		// 2. params (of function)
		paramTuple := &types.TypeTuple{}
		err = DecodeTypeTuple(rd, paramTuple)
		if err != nil {
			return err
		}

		// 3. results (of function)
		resultTuple := &types.TypeTuple{}
		err = DecodeTypeTuple(rd, resultTuple)
		if err != nil {
			return err
		}
		functionType.Params = paramTuple
		functionType.Results = resultTuple
		p.Module.Types = append(p.Module.Types, functionType)
		logrus.Infof("FunctionType: %v", functionType)
	}
	err = p.ValidateTypes()
	return err
}

func DecodeTypeTuple(rd io.Reader, tuple *types.TypeTuple) error {
	// 1. count of params (or results)
	n, err := utils.DecodeUInt32(rd)
	if err != nil {
		return nil
	}
	tuple.NumElems = n
	// 2. params (or results) array
	for i := 0; i < int(n); i++ {
		vType, err := utils.DecodeInt32(rd)
		if err != nil {
			return err
		}
		valueType, err := types.DecodeValueType(vType)
		if err != nil {
			return err
		}
		tuple.Elems = append(tuple.Elems, valueType)
	}
	return nil
}

func (p *Parser) ValidateTypes() error {
	logrus.Info("TODO: ValidateTypes()")
	return nil
}
