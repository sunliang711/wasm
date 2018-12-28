package parser
#include "OperatorTable.h"
//Note this file is created by makeTypes.sh,Don't modify this file directly
import(
	"bytes"
	"encoding/binary"
	"io"
	"wasm/types/IR"
)

#define VISIT_OPCODE(OPCODE,NAME,NAMESTRING,IMM,...) \
    		case IR.NAME: \
    			imm := IR.IMM{}; \
    			err = Decode##IMM(rd, &imm, funcDef); \
    			if err != nil { \
    				return nil, err; \
    			}; \
    			opimm := IR.OpcodeAndImm_##IMM{}; \
    			opimm.Imm = imm; \
    			opimm.Opcode = IR.NAME; \
    			err = binary.Write(&buf, binary.LittleEndian, &opimm); \
    			if err != nil { \
    				return nil, err; \
    			}; \
    			ret = append(ret, buf.Bytes()...);

func DecodeOpcodeAndImm(opcodeBytes []byte, funcDef *IR.FunctionDef) ([]byte, error) {
	rd := bytes.NewReader(opcodeBytes)
	var ret []byte

	for {
		opc, err := DecodeOpcode(rd)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		var buf bytes.Buffer
		switch IR.Opcode(opc) {
ENUM_OPERATORS(VISIT_OPCODE)
		}
	}
	return ret, nil
}

#undef VISIT_OPCODE