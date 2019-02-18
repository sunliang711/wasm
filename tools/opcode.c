package IR
#include "OperatorTable.h"
type (
	Opcode uint16
	Op struct {
		Code Opcode
		Gas  uint64
		Name string
	}
)

var Ops [0xffff]Op

func NewOp(code Opcode, gas uint64, name string) Opcode {
	o := Op{code, gas, name}
	Ops[code] = o
	return code
}

var(
#define VISIT_OPCODE(opcode,name,nameStr,imm,gas,...) \
    name = NewOp(opcode,gas,nameStr);
    ENUM_OPERATORS(VISIT_OPCODE)
#undef VISIT_OPCODE
     OPCMaxSingleByteOpcode Opcode = 0xdf

)
