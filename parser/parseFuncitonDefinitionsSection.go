package parser

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"wasm/types"
	"wasm/types/IR"
	"wasm/utils"
)

func (p *Parser) functionDefinitionsSection(sec *Section) error {
	err := checkSection(sec, IR.OrderFunctionDefinitions)
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

		var ls IR.LocalSet
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
		////TODO code validation
		//extendedCodes, err := DecodeOpcodeAndImm(opcodeBytes, &funcDef)
		//if err != nil {
		//	return err
		//}
		//funcDef.Code = extendedCodes
		funcDef.Code = opcodeBytes
		ins, err := DecodeOpcodeAndImm(opcodeBytes, &funcDef)
		if err != nil {
			return err
		}
		funcDef.Instruction = ins
		endIndice, err := buildInsRelationship(ins)
		funcDef.EndIndice = endIndice

		p.Module.Functions.Defs[i] = funcDef
		logrus.Infof("<function definition> function definition: %v", funcDef)
	}

	return nil
}

func buildInsRelationship(ins []IR.Instruction) ([]int, error) {
	if len(ins) == 0 {
		return nil, fmt.Errorf("Empty instruction")
	}
	if ins[len(ins)-1].Op.Code != IR.OPCend {
		return nil, fmt.Errorf("instruction not end with 'end'")
	}
	ins[len(ins)-1].MatchedIndex = -2
	var endIndice []int
	fullStack := utils.Stack{}
	ifStack := utils.Stack{}
	for index := range ins[:len(ins)-1] {
		fullStack.Push(&ins[index])
		switch ins[index].Op.Code {
		case IR.OPCblock, IR.OPCloop:
		case IR.OPCif_:
			ifStack.Push(&ins[index])
		case IR.OPCelse_:
			t, err := ifStack.Pop()
			if err != nil {
				return nil, fmt.Errorf("buildInsRelationship(): ifStack empty")
			}
			t.(*IR.Instruction).MatchedIndex = ins[index].Index
		case IR.OPCend:
			endIndice = append(endIndice, ins[index].Index)
		INNER_LOOP:
			for {
				popedIns, err := fullStack.Pop()
				if err != nil {
					return nil, fmt.Errorf("buildInsRelationship(): fullStack empty")
				}
				switch popedIns.(*IR.Instruction).Op.Code {
				case IR.OPCif_, IR.OPCloop, IR.OPCblock:
					ins[index].MatchedIndex = popedIns.(*IR.Instruction).Index
					break INNER_LOOP
				}
			}
		}
	}
	endIndice = append(endIndice, len(ins)-1)

	return endIndice, nil
}
