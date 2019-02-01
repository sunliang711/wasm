package runtime

import (
	"wasm/core/IR"
	"wasm/utils"
)

const (
	SI32Convert byte = iota
	UI32Convert
	SI64Convert
	UI64Convert
)

func f32Convert(vm *VM, frame *Frame, convertType byte) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}

	switch convertType {
	case SI32Convert:
		m, ok := a.Value().(int32)
		if !ok {
			panic("f32.convert parameter not valid")
		}
		frame.Stack.Push(&Value{IR.TypeF32, float32(m)})
	case UI32Convert:
		m, ok := a.Value().(uint32)
		if !ok {
			panic("f32.convert parameter not valid")
		}
		frame.Stack.Push(&Value{IR.TypeF32, float32(m)})
	case SI64Convert:
		m, ok := a.Value().(int64)
		if !ok {
			panic("f32.convert parameter not valid")
		}
		frame.Stack.Push(&Value{IR.TypeF32, float32(m)})
	case UI64Convert:
		m, ok := a.Value().(uint64)
		if !ok {
			panic("f32.convert parameter not valid")
		}
		frame.Stack.Push(&Value{IR.TypeF32, float32(m)})
	default:
		panic("f32.convert type error")
	}
	frame.advance(1)

	return
}

func f64Convert(vm *VM, frame *Frame, convertType byte) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}

	switch convertType {
	case SI32Convert:
		m, ok := a.Value().(int32)
		if !ok {
			panic("f64.convert parameter not valid")
		}
		frame.Stack.Push(&Value{IR.TypeF64, float64(m)})
	case UI32Convert:
		m, ok := a.Value().(uint32)
		if !ok {
			panic("f64.convert parameter not valid")
		}
		frame.Stack.Push(&Value{IR.TypeF64, float64(m)})
	case SI64Convert:
		m, ok := a.Value().(int64)
		if !ok {
			panic("f64.convert parameter not valid")
		}
		frame.Stack.Push(&Value{IR.TypeF64, float64(m)})
	case UI64Convert:
		m, ok := a.Value().(uint64)
		if !ok {
			panic("f64.convert parameter not valid")
		}
		frame.Stack.Push(&Value{IR.TypeF64, float64(m)})
	default:
		panic("f64.convert type error")
	}
	frame.advance(1)

	return
}
