package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"wasm/types"
	"wasm/types/IR"
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

	err := checkSection(sec, IR.OrderType)
	if err != nil {
		return err
	}

	var (
		functionType IR.FunctionType
	)

	rd := bytes.NewReader(sec.Data)
	// type section is an array
	// get array size
	var size uint32
	_, err = utils.DecodeVarInt(rd, 32, &size)
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
		paramTuple := &IR.TypeTuple{}
		err = DecodeTypeTuple(rd, paramTuple)
		if err != nil {
			return err
		}

		// 3. results (of function)
		resultTuple := &IR.TypeTuple{}
		err = DecodeTypeTuple(rd, resultTuple)
		if err != nil {
			return err
		}
		functionType.Params = paramTuple
		functionType.Results = resultTuple
		p.Module.Types = append(p.Module.Types, functionType)
		logrus.Infof("<type section>index: %d, FunctionType: %v", index, functionType)
	}
	err = p.ValidateTypes()
	return err
}

func DecodeTypeTuple(rd io.Reader, tuple *IR.TypeTuple) error {
	// 1. count of params (or results)
	var n uint32
	_, err := utils.DecodeVarInt(rd, 32, &n)
	if err != nil {
		return nil
	}
	tuple.NumElems = n
	// 2. params (or results) array
	for i := 0; i < int(n); i++ {
		valueType, err := DecodeValueTypeFromReader(rd)
		if err != nil {
			return err
		}
		tuple.Elems = append(tuple.Elems, valueType)
	}
	return nil
}

func (p *Parser) ValidateTypes() error {
	//TODO
	logrus.Info("TODO: ValidateTypes()")
	return nil
}

func DecodeValueTypeFromReader(rd io.Reader) (IR.ValueType, error) {
	var vType int8
	_, err := utils.DecodeVarInt(rd, 7, &vType)
	if err != nil {
		return IR.TypeNone, err
	}
	return DecodeValueType(vType)
}

func DecodeValueType(vt int8) (IR.ValueType,error)  {
	switch vt {
	case -1:
		return IR.TypeI32, nil
	case -2:
		return IR.TypeI64, nil
	case -3:
		return IR.TypeF32, nil
	case -4:
		return IR.TypeF64, nil
	case -5:
		return IR.TypeV128, nil
	case -16:
		return IR.TypeAnyFunc, nil
	case -17:
		return IR.TypeAnyRef, nil
	default:
		return IR.TypeNone, fmt.Errorf(IR.ErrInvalidValueType)
	}

}
