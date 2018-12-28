package types
#include "OperatorTable.h"

//Note this file is created by makeTypes.sh,Don't modify this file directly
type Opcode uint16

const (
#define VISIT_OPCODE(opcode,name,...) name Opcode = opcode;
    ENUM_OPERATORS(VISIT_OPCODE)
#undef VISIT_OPCODE
    OPCMaxSingleByteOpcode Opcode = 0xdf
)

//1. cd to tools/
//2. run 'macro2go.sh -d -o ../types/opcode.go opcode.tmpl'
