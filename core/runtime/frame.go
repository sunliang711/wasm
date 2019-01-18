package runtime

import (
	"fmt"
	"wasm/core/IR"
)

type Frame struct {
	*IR.FunctionDef
	FuncitonID int
	PC         int
	*IR.Stack
	Locals []IR.InterfaceValue //parameters、locals :used by get_local set_local ...
}

func (f *Frame) advance(step int) {
	f.PC += step
}

func (f *Frame) advanceTo(dest int) {
	f.PC = dest
}

func (f *Frame) Init(fID int, vm *VM, params []IR.InterfaceValue) error {
	f.FuncitonID = fID
	f.FunctionDef = vm.FunctionCodes[fID]
	f.PC = 0

	fType := vm.Module.Types[int(f.FunctionDef.Type.Index)]
	//check parameters count
	if int(fType.Params.NumElems) != len(params) {
		return fmt.Errorf("parameter counter not match")
	}
	//check parameters type
	for index := range fType.Params.Elems {
		if fType.Params.Elems[index] != params[index].Type() {
			return fmt.Errorf("parameter type not match")
		}
	}
	localLen := int(fType.Params.NumElems) + len(f.FunctionDef.NonParameterLocalTypes)
	f.Locals = make([]IR.InterfaceValue, localLen)
	for i := range params {
		f.Locals[i] = params[i]
	}

	return nil
}
