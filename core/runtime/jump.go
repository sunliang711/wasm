package runtime

import (
	"wasm/core/IR"
	"wasm/types"
	"wasm/utils"
)

func returnFunc(vm *WasmInterpreter, frame **Frame) (exit bool, err error) {
	defer utils.CatchError(&err)

	var retV IR.InterfaceValue
	vm.CurrentFrame -= 1
	hasResult := !(*frame).Stack.Empty()
	if hasResult {
		//has return value
		retV, _ = (*frame).Stack.Pop()
		//TODO check frame.FunctionSig with retV
	}
	if vm.CurrentFrame == -1 {
		if hasResult {
			vm.ReturnValue = retV
		}
		exit = true
		return
	} else {
		if hasResult {
			(*frame).Stack.Push(retV)
		}
		*frame = vm.GetCurrentFrame()
	}
	return
}

func call(vm *WasmInterpreter, frame **Frame, predefine bool) (err error) {
	defer utils.CatchError(&err)

	ins := (*frame).Instruction[(*frame).PC]
	(*frame).advance(1)
	funcImm, ok := ins.Imm.(*IR.FunctionImm)
	if !ok {
		vm.panic("opcode: call invalid imm")
	}
	funcIndex := funcImm.FunctionIndex
	var predefineLen uint64
	if predefine {
		predefineLen = uint64(len(PredefinedFuncs))
	}
	if funcIndex < predefineLen {
		PredefinedFuncs[funcIndex].Action(vm, *frame)
	} else {
		vm.CurrentFrame += 1
		if vm.CurrentFrame >= MAXFRAME {
			vm.panic(types.ErrBeyondMaxFrame)
		}

		fType := vm.Module.Types[int(vm.FunctionCodes[funcIndex].Type.Index)]
		paraCount := fType.Params.NumElems
		if (*frame).Stack.Len() < int(paraCount) {
			vm.panic(types.ErrStackSizeErr)
		}
		params := make([]IR.InterfaceValue, paraCount)
		for elemIndex := range fType.Params.Elems {
			v, _ := (*frame).Stack.Pop()
			params[int(paraCount)-1-elemIndex] = v
		}

		*frame = vm.GetCurrentFrame()
		err = (*frame).Init(int(funcIndex), vm, params)
		if err != nil {
			vm.panic(err)
		}
	}
	return
}

func call_indirect(vm *WasmInterpreter, frame **Frame) (err error) {
	defer utils.CatchError(&err)
	//TODO: check imm
	//imm, ok := ins.Imm.(*IR.CallIndirectImm)
	//if !ok {
	//	vm.panic("opcode: call_indirect invalid imm")
	//}
	(*frame).advance(1)
	calleeIndex, err := pop1(vm, *frame)
	if err != nil {
		vm.panic(err)
	}
	funcIndex, ok := calleeIndex.Value().(int)
	if !ok {
		vm.panic("call_indirect: funcIndex not int type")
	}
	vm.CurrentFrame += 1
	if vm.CurrentFrame >= MAXFRAME {
		vm.panic(types.ErrBeyondMaxFrame)
	}

	fType := vm.Module.Types[int(vm.FunctionCodes[funcIndex].Type.Index)]
	paraCount := fType.Params.NumElems
	if (*frame).Stack.Len() < int(paraCount) {
		vm.panic(types.ErrStackSizeErr)
	}
	params := make([]IR.InterfaceValue, paraCount)
	for elemIndex := range fType.Params.Elems {
		v, _ := (*frame).Stack.Pop()
		params[int(paraCount)-1-elemIndex] = v
	}

	*frame = vm.GetCurrentFrame()
	err = (*frame).Init(int(funcIndex), vm, params)
	if err != nil {
		vm.panic(err)
	}
	//get types[imm.Type]

	//get table Index

	//call func specified by table index
	return
}

func end(vm *WasmInterpreter, frame **Frame, ifStack *utils.Stack) (exit bool, err error) {
	defer utils.CatchError(&err)

	ins := (*frame).Instruction[(*frame).PC]
	switch ins.MatchedIndex {
	case types.LastOpcode:
		//if ins.Index == len(frame.Instruction)-1 {
		var retV IR.InterfaceValue
		vm.CurrentFrame -= 1
		hasResult := !(*frame).Stack.Empty()
		if (*frame).FunctionDef.FunctionType.Results.NumElems > 0 && !hasResult {
			panic("This function should have return value,but stack is empty")
		}
		if hasResult {
			//has return value
			retV, _ = (*frame).Stack.Pop()
			//TODO check frame.FunctionSig with retV
		}
		if vm.CurrentFrame == -1 {
			if hasResult {
				vm.ReturnValue = retV
			}
			exit = true
			return
		} else {
			*frame = vm.GetCurrentFrame()
			if hasResult {
				//TODO assert frame.Stack is empty
				(*frame).Stack.Push(retV)
			}
		}
	case -1:
		vm.panic("end ins matched index illegal")
	default:
		//TODO pop value stack since block or if
		mindex := ins.MatchedIndex
		switch (*frame).Instruction[mindex].Op.Code {
		case IR.OPCblock:
			(*frame).advance(1)
		case IR.OPCif_:
			ifStack.Pop()
			(*frame).advance(1)
		case IR.OPCloop:
			(*frame).advanceTo(mindex + 1)
		default:
			vm.panic("end ins not point to block/if/loop")
		}
	}
	return
}

func ifFunc(vm *WasmInterpreter, frame *Frame, ifStack *utils.Stack) (err error) {
	defer utils.CatchError(&err)

	if frame.Stack.Len() < 1 {
		vm.panic(types.ErrStackSizeErr)
	}
	ins := frame.Instruction[frame.PC]
	con, _ := frame.Stack.Pop()
	if !IsZero(con) {
		ifStack.Push(ins.Index)
		frame.advance(1)
	} else {
		//advanctTo else or end
		nextIndex := ins.MatchedIndex
		if nextIndex == -1 {
			vm.panic("if without else or end")
		}
		frame.advanceTo(nextIndex)
	}
	return
}

func elseFunc(vm *WasmInterpreter, frame *Frame, ifStack *utils.Stack) (err error) {
	defer utils.CatchError(&err)

	ins := frame.Instruction[frame.PC]
	if ifStack.Empty() {
		frame.advance(1)
	} else {
		frame.advanceTo(ins.MatchedIndex)
	}
	return
}

func selectFunc(vm *WasmInterpreter, frame *Frame) (err error) {
	defer utils.CatchError(&err)

	if frame.Stack.Len() < 3 {
		vm.panic(types.ErrStackSizeErr)
	}
	a, _ := frame.Stack.Pop()
	b, _ := frame.Stack.Pop()
	c, _ := frame.Stack.Pop()
	if IsZero(c) {
		frame.Stack.Push(b)
	} else {
		frame.Stack.Push(a)
	}
	frame.advance(1)
	return
}

func drop(vm *WasmInterpreter, frame *Frame) (err error) {
	defer utils.CatchError(&err)

	if frame.Stack.Len() < 1 {
		vm.panic(types.ErrStackSizeErr)
	}
	frame.Stack.Pop()
	frame.advance(1)
	return
}
