package runtime

import (
	"wasm/core/IR"
	"wasm/utils"
)

const (
	SF32TRUNC byte = iota
	SF64TRUNC
	UF32TRUNC
	UF64TRUNC
)

func i32Trunc(vm *WasmInterpreter, frame *Frame, truncType byte) (err error) {
	defer utils.CatchError(&err)

	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}

	switch truncType {
	case SF32TRUNC:
		v, ok := a.Value().(float32)
		if !ok {
			panic("i32.trunc_s/f32 parameter invalid")
		}
		frame.Stack.Push(&Value{IR.TypeI32, int32(v)})
	case SF64TRUNC:
		v, ok := a.Value().(float64)
		if !ok {
			panic("i32.trunc_s/f64 parameter invalid")
		}
		frame.Stack.Push(&Value{IR.TypeI32, int32(v)})
	case UF32TRUNC:
		v, ok := a.Value().(float32)
		if !ok {
			panic("i32.trunc_u/f32 parameter invalid")
		}
		frame.Stack.Push(&Value{IR.TypeI32, uint32(v)})
	case UF64TRUNC:
		v, ok := a.Value().(float64)
		if !ok {
			panic("i32.trunc_u/f64 parameter invalid")
		}
		frame.Stack.Push(&Value{IR.TypeI32, uint32(v)})
	default:
		vm.panic("i32Trunc trunc type invalid")
	}

	frame.advance(1)
	return
}

func i64Trunc(vm *WasmInterpreter, frame *Frame, truncType byte) (err error) {
	defer utils.CatchError(&err)

	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}

	switch truncType {
	case SF32TRUNC:
		v, ok := a.Value().(float32)
		if !ok {
			panic("i64.trunc_s/f32 parameter invalid")
		}
		frame.Stack.Push(&Value{IR.TypeI64, int64(v)})
	case SF64TRUNC:
		v, ok := a.Value().(float64)
		if !ok {
			panic("i64.trunc_s/f64 parameter invalid")
		}
		frame.Stack.Push(&Value{IR.TypeI64, int64(v)})
	case UF32TRUNC:
		v, ok := a.Value().(float32)
		if !ok {
			panic("i64.trunc_u/f32 parameter invalid")
		}
		frame.Stack.Push(&Value{IR.TypeI64, uint64(v)})
	case UF64TRUNC:
		v, ok := a.Value().(float64)
		if !ok {
			panic("i64.trunc_u/f64 parameter invalid")
		}
		frame.Stack.Push(&Value{IR.TypeI64, uint64(v)})
	default:
		vm.panic("i64Trunc trunc type invalid")
	}

	frame.advance(1)
	return
}

func wrapI64ToI32(vm *WasmInterpreter, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case int64:
		frame.Stack.Push(&Value{IR.TypeI32, int32(a.Value().(int64))})
	case uint64:
		frame.Stack.Push(&Value{IR.TypeI32, int32(a.Value().(uint64))})
	default:
		panic("i32.wrap/i64 parameter invalid")
	}

	frame.advance(1)
	return
}
