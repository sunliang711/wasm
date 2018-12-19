package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"wasm/types"
	"wasm/utils"
)

var (
	functionTag = []byte{0x60}
)

func (p *Parser) typeSection(sec *Section) error {
	defer func() {
		//notify import and functionDeclaration sections
		p.typeParsed <- struct{}{}
		p.typeParsed <- struct{}{}
	}()

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
	var size uint32
	err = utils.DecodeVarInt(rd, 32, &size)
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
		logrus.Infof("<type section> FunctionType: %v,index: %d", functionType, index)
	}
	err = p.ValidateTypes()
	return err
}

func DecodeTypeTuple(rd io.Reader, tuple *types.TypeTuple) error {
	// 1. count of params (or results)
	var n uint32
	err := utils.DecodeVarInt(rd, 32, &n)
	if err != nil {
		return nil
	}
	tuple.NumElems = n
	// 2. params (or results) array
	for i := 0; i < int(n); i++ {
		valueType, err := DecodeValueType(rd)
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

func DecodeValueType(rd io.Reader) (types.ValueType, error) {
	var vType int8
	err := utils.DecodeVarInt(rd, 7, &vType)
	if err != nil {
		return types.TypeNone, err
	}

	switch vType {
	case -1:
		return types.TypeI32, nil;
	case -2:
		return types.TypeI64, nil;
	case -3:
		return types.TypeF32, nil;
	case -4:
		return types.TypeF64, nil;
	case -5:
		return types.TypeV128, nil;
	case -16:
		return types.TypeAnyFunc, nil;
	case -17:
		return types.TypeAnyRef, nil;
	default:
		return types.TypeNone, fmt.Errorf(types.ErrInvalidValueType)
	}
}
