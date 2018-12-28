package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"wasm/types"
	"wasm/utils"
)

func (p *Parser) functionDefinitionsSection(sec *Section) error {
	err := checkSection(sec, types.OrderFunctionDefinitions)
	if err != nil {
		return err
	}
	rd := bytes.NewReader(sec.Data)

	<-p.funcDeclarationParsed
	//1. num function bodys
	var numFunc uint32
	_, err = utils.DecodeVarInt(rd, 32, &numFunc)
	if err != nil {
		return err
	}
	if int(numFunc) != len(p.Module.Functions.Defs) {
		return fmt.Errorf(types.ErrFuncDeclarationDefinitionMismatch)
	}

	//2. function bodys
	for i := 0; i < int(numFunc); i++ {
		funcDef := p.Module.Functions.Defs[i]
		//a. num body size
		var numBodyBytes uint32
		_, err := utils.DecodeVarInt(rd, 32, &numBodyBytes)
		if err != nil {
			return err
		}
		//b. localSet
		var numLocalSets uint32
		numUsedBytes, err := utils.DecodeVarInt(rd, 32, &numLocalSets)
		if err != nil {
			return err
		}

		var ls types.LocalSet
		localSetBodyUsedBytes := 0
		for j := 0; j < int(numLocalSets); j++ {
			usedBytes, err := DecodeLocalSet(rd, &ls)
			if err != nil {
				return err
			}
			if uint64(len(funcDef.NonParameterLocalTypes))+ls.Num > p.Module.FeatureSpec.MaxLocals {
				return fmt.Errorf(types.ErrTooManyLocals)
			}

			for k := uint64(0); k < ls.Num; k++ {
				funcDef.NonParameterLocalTypes = append(funcDef.NonParameterLocalTypes, ls.Type)
			}
			localSetBodyUsedBytes += usedBytes
		}
		//c. op code
		opcodeSize := int(numBodyBytes) - numUsedBytes - localSetBodyUsedBytes
		opcodeBytes, err := utils.ReadNByte(rd, opcodeSize)
		if err != nil {
			return err
		}
		extendedCodes,err :=DecodeOpcodeAndImm(opcodeBytes,&funcDef)
		if err != nil {
			return err
		}
		funcDef.Code = extendedCodes
		p.Module.Functions.Defs[i] = funcDef
		logrus.Infof("<function definition> function definition: %v", funcDef)
	}

	return nil
}
