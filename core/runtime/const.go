package runtime

import (
	"wasm/core/IR"
	"wasm/utils"
)

const (
	I32_CONST byte = iota
	I64_CONST
	F32_CONST
	F64_CONST
)

func defConst(vm *VM, frame *Frame, constType byte) (err error) {
	defer utils.CatchError(&err)
	ins := frame.Instruction[frame.PC]
	switch constType {
	case I32_CONST:
		literalImm, ok := ins.Imm.(*IR.LiteralImm_I32)
		if !ok {
			panic("i32.const imm invalid")
		}
		val := literalImm.Value
		frame.Stack.Push(&Value{Typ: IR.TypeI32, Val: val})
	case I64_CONST:
		literalImm, ok := ins.Imm.(*IR.LiteralImm_I64)
		if !ok {
			panic("i64.const imm invalid")
		}
		val := literalImm.Value
		frame.Stack.Push(&Value{IR.TypeI64, val})
	case F32_CONST:
		literalImm, ok := ins.Imm.(*IR.LiteralImm_F32)
		if !ok {
			panic("f32.const imm invalid")
		}
		val := literalImm.Value
		frame.Stack.Push(&Value{IR.TypeF32, val})
	case F64_CONST:
		literalImm, ok := ins.Imm.(*IR.LiteralImm_F64)
		if !ok {
			panic("f64.const imm invalid")
		}
		val := literalImm.Value
		frame.Stack.Push(&Value{IR.TypeF64, val})
	default:
		panic("defConst const type invalid")
	}

	frame.advance(1)
	return
}
