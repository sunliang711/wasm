package parser
#include "OperatorTable.h"
//Note this file is created by makeTypes.sh,Don't modify this file directly
import(
	"bytes"
	"io"
	"wasm/core/IR"
	"fmt"
)

#define VISIT_OPCODE(OPCODE,NAME,NAMESTRING,IMM,...) \
    		case IR.NAME: \
    			imm := IR.IMM{}; \
    			err = Decode##IMM(rd, &imm, funcDef); \
    			if err != nil { \
    				return nil, err; \
    			}; \
                ins = append(ins, IR.Instruction{&IR.Ops[IR.NAME], &imm, codeIndex,-1});

func DecodeOpcodeAndImm(opcodeBytes []byte, funcDef *IR.FunctionDef) ([]IR.Instruction, error) {
	rd := bytes.NewReader(opcodeBytes)
    var (
        ins       []IR.Instruction
        codeIndex int
    	)

	for {
		opc, err := DecodeOpcode(rd)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch IR.Opcode(opc) {
ENUM_OPERATORS(VISIT_OPCODE)
		}
		codeIndex += 1
	}
		if ins[len(ins)-1].Op.Code != IR.OPCend {
    		return  nil,fmt.Errorf("code not end with \"end\"")
    	}

/*
        ins[len(ins)-1].MatchedIndex = -2
    	stack := IR.Stack{}
    	endIndice := make([]int, 0)
    	beginPush := false
    	for instrIndex, instr := range ins[:len(ins)-1] {
    		switch instr.Op.Code {
    		case IR.OPCloop, IR.OPCif_, IR.OPCblock:
    			beginPush = true
    			stack.Push(&ins[instrIndex])
    		case IR.OPCend:
    			endIndice = append(endIndice, instr.Index)
			INNER_LOOP:
    			for {
    				i, err := stack.Pop()
    				if err != nil {
    					return  nil, fmt.Errorf("Stack pop failed")
    				}
    				switch i.Value().(*IR.Instruction).Op.Code {
    				case IR.OPCloop, IR.OPCif_, IR.OPCblock:
                        ins[instrIndex].MatchedIndex = i.Value().(*IR.Instruction).Index
                        break INNER_LOOP
    				}
    			}

    		default:
    			if beginPush {
                    stack.Push(&ins[instrIndex])
    			}
    		}
    	}
    	if stack.Len()>1 {
    		return nil, nil, fmt.Errorf("instructions end count not match")
    	}
*/
    	return ins, nil
}

#undef VISIT_OPCODE