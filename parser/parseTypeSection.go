package parser

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"wasm/types"
	"wasm/utils"
)

var (
	functionTag = []byte{0x60}
)

func (p *Parser) typeSection(sec *Section) error {

	err := checkSection(sec, types.OrderType)
	if err != nil {
		return err
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
		err = checkConstant(rd, functionTag, types.ErrFunctionTag)
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
		valueType, err := types.DecodeValueType(rd)
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
