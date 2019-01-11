package IR
#include "OperatorTable.h"
type (
	Opcode uint16
	Op struct {
		Code Opcode
		Name string
	}
)

var Ops [0xffff]Op

func NewOp(code Opcode, name string) Opcode {
	o := Op{code, name}
	Ops[code] = o
	return code
}

var(
#define VISIT_OPCODE(opcode,name,nameStr,...) \
    name = NewOp(opcode,nameStr);
    ENUM_OPERATORS(VISIT_OPCODE)
#undef VISIT_OPCODE
     OPCMaxSingleByteOpcode Opcode = 0xdf

)
