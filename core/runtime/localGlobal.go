package runtime

import (
	"wasm/core/IR"
	"wasm/types"
	"wasm/utils"
)

func getLocal(vm *WasmInterpreter, frame *Frame) (err error) {
	defer utils.CatchError(&err)

	ins := frame.Instruction[frame.PC]
	imm, ok := ins.Imm.(*IR.GetOrSetVariableImm)
	if !ok {
		vm.panic("opcode: get_local invalid imm")
	}
	index := imm.VariableIndex
	if index >= uint64(len(frame.Locals)) {
		vm.panic("get_local index out of range")
	}
	frame.Stack.Push(frame.Locals[index])
	frame.advance(1)
	return
}

func setLocal(vm *WasmInterpreter, frame *Frame) (err error) {
	defer utils.CatchError(&err)

	ins := frame.Instruction[frame.PC]
	imm, ok := ins.Imm.(*IR.GetOrSetVariableImm)
	if !ok {
		vm.panic("opcode: set_local invalid imm")
	}
	index := imm.VariableIndex
	if index >= uint64(len(frame.Locals)) {
		vm.panic("set_local index out of range")
	}
	if frame.Stack.Len() < 1 {
		vm.panic(types.ErrStackSizeErr)
	}
	a, _ := frame.Stack.Pop()
	frame.Locals[index] = a
	frame.advance(1)
	return
}

func teeLocal(vm *WasmInterpreter, frame *Frame) (err error) {
	defer utils.CatchError(&err)

	ins := frame.Instruction[frame.PC]
	imm, ok := ins.Imm.(*IR.GetOrSetVariableImm)
	if !ok {
		vm.panic("opcode: tee_local invalid imm")
	}
	index := imm.VariableIndex
	if index >= uint64(len(frame.Locals)) {
		vm.panic("tee_local index out of range")
	}
	if frame.Stack.Len() < 1 {
		vm.panic(types.ErrStackSizeErr)
	}
	a, _ := frame.Stack.Top()
	frame.Locals[index] = a
	frame.advance(1)
	return
}

func getGlobal(vm *WasmInterpreter, frame *Frame) (err error) {
	defer utils.CatchError(&err)

	ins := frame.Instruction[frame.PC]
	imm, ok := ins.Imm.(*IR.GetOrSetVariableImm)
	if !ok {
		vm.panic("opcode: get_local invalid imm")
	}
	index := imm.VariableIndex
	if index >= uint64(len(vm.Global)) {
		vm.panic("get_local index out of range")
	}
	frame.Stack.Push(vm.Global[index])
	frame.advance(1)
	return
}

func setGlobal(vm *WasmInterpreter, frame *Frame) (err error) {
	defer utils.CatchError(&err)

	ins := frame.Instruction[frame.PC]
	imm, ok := ins.Imm.(*IR.GetOrSetVariableImm)
	if !ok {
		vm.panic("opcode: get_local invalid imm")
	}
	index := imm.VariableIndex
	if index >= uint64(len(vm.Global)) {
		vm.panic("set_local index out of range")
	}
	if frame.Stack.Len() < 1 {
		vm.panic(types.ErrStackSizeErr)
	}
	a, _ := frame.Stack.Pop()
	vm.Global[index] = a
	frame.advance(1)
	return
}
