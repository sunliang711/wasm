package runtime

import "wasm/core/IR"

const (
	I32_CONST byte = iota
	I64_CONST
	F32_CONST
	F64_CONST
)

func defConst(vm *VM, frame *Frame, constType byte) {
	ins := frame.Instruction[frame.PC]
	switch constType {
	case I32_CONST:
		val := ins.Imm.(*IR.LiteralImm_I32).Value
		frame.Stack.Push(&Value{Typ: IR.TypeI32, Val: val})
	case I64_CONST:
		val := ins.Imm.(*IR.LiteralImm_I64).Value
		frame.Stack.Push(&Value{IR.TypeI64, val})
	case F32_CONST:
		val := ins.Imm.(*IR.LiteralImm_F32).Value
		frame.Stack.Push(&Value{IR.TypeF32, val})
	case F64_CONST:
		val := ins.Imm.(*IR.LiteralImm_F64).Value
		frame.Stack.Push(&Value{IR.TypeF64, val})
	}

	frame.advance(1)
}
