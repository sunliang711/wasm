package runtime

import (
	"fmt"
	"wasm/types/IR"
)

const (
	PAGESIZE = 65536
	MAXFRAME = 128
)

type Value struct {
	Typ IR.ValueType
	Val interface{}
}

func (v *Value) Type() IR.ValueType {
	return v.Typ
}
func (v *Value) Value() interface{} {
	return v.Val
}

//
//type FunctionCode struct {
//	NumParams      int              // to init frame.Locals
//	NonParamLocals []IR.InterfaceValue //locals only,not including parameters,to init frame.Locals
//	Codes          []IR.Instruction
//	Result         IR.InterfaceValue
//}

var (
	//TODO: for resolving import
	modules = make(map[string]*IR.Module)
)

type VM struct {
	Module       *IR.Module //decodeOpcodeAndImm to []FunctionCode //TODO TODO
	Frames       []*Frame
	CurrentFrame int
	//FunctionCodes []FunctionCode // frame.FunctionCode:= functionCode[i]  i:=frame.functionID
	FunctionCodes []*IR.FunctionDef // frame.FunctionCode:= functionCode[i]  i:=frame.functionID
	ReturnValue   IR.InterfaceValue
	Global        []IR.InterfaceValue
	Table         []uint32
	Memory        []byte //pagesize * inital
}

func NewVM(module *IR.Module) (*VM, error) {
	vm := &VM{Module: module, CurrentFrame: 0}
	//TODO resolve import

	//1. init functionCodes
	vm.FunctionCodes = make([]*IR.FunctionDef, len(module.Functions.Defs)+len(module.Functions.Imports))
	for i := range module.Functions.Defs {
		vm.FunctionCodes[i] = &module.Functions.Defs[i]
	}

	//2. init global
	vm.Global = make([]IR.InterfaceValue, len(module.Globals.Defs))
	for i, def := range module.Globals.Defs {
		valType := def.Type.ValType
		initType := def.Initializer.Type
		switch initType {
		case IR.I32_const:
			vm.Global[i] = &Value{valType, def.Initializer.I32}
		case IR.I64_const:
			vm.Global[i] = &Value{valType, def.Initializer.I64}
		case IR.F32_const:
			vm.Global[i] = &Value{valType, def.Initializer.F32}
		case IR.F64_const:
			vm.Global[i] = &Value{valType, def.Initializer.F64}
		case IR.Get_global:
			vm.Global[i] = &Value{valType, vm.Global[def.Initializer.GlobalRef]}
			//TODO :other type
		}
	}

	//3. init memory
	if len(module.Memories.Defs) > 0 {
		vm.Memory = make([]byte, module.Memories.Defs[0].Type.Size.Min*PAGESIZE)
		for i := 0; i < len(vm.Memory); i++ {
			vm.Memory[i] = 0
		}

		offset := 0
		for _, dataSeg := range module.DataSegments {
			switch dataSeg.BaseOffset.Type {
			case IR.I32_const:
				offset = int(dataSeg.BaseOffset.I32)
			case IR.I64_const:
				offset = int(dataSeg.BaseOffset.I64)
			case IR.F32_const:
				offset = int(dataSeg.BaseOffset.F32)
			case IR.F64_const:
				offset = int(dataSeg.BaseOffset.F64)
			case IR.Get_global:
				offset = vm.Global[int(dataSeg.BaseOffset.GlobalRef)].Value().(int)
			}
			copy(vm.Memory[offset:], dataSeg.Data)
		}
	}
	//4. init table TODO

	//5. init frames
	vm.Frames = make([]*Frame, MAXFRAME)
	for i := range vm.Frames {
		vm.Frames[i] = &Frame{
			PC:    0,
			Stack: &IR.Stack{},
		}
	}
	return vm, nil
}

func (vm *VM) GetCurrentFrame() *Frame {
	if vm.CurrentFrame >= MAXFRAME {
		return nil
	}
	return vm.Frames[vm.CurrentFrame]
}

func (vm *VM) Traceback() {
	ret := ""
	for i := vm.CurrentFrame; i >= 0; i-- {
		ret += "func " + vm.Frames[i].Name + " "
		ret += vm.Frames[i].FunctionType.String() + "\n"
	}
	fmt.Println("Traceback:")
	fmt.Printf(ret)
}

func (vm *VM) panic(v interface{}) {
	vm.Traceback()
	panic(v)
}
