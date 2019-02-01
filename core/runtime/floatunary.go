package runtime

import (
	"math"
	"wasm/core/IR"
	"wasm/utils"
)

func f32Abs(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float32:
		ret := a.Value().(float32)
		if ret < 0 {
			ret = -ret
		}
		frame.Stack.Push(&Value{IR.TypeF32, ret})
	default:
		panic("f32.abs parameter invalid")
	}
	frame.advance(1)
	return
}

func f64Abs(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float64:
		ret := a.Value().(float64)
		if ret < 0 {
			ret = -ret
		}
		frame.Stack.Push(&Value{IR.TypeF64, ret})
	default:
		panic("f64.abs parameter invalid")
	}
	frame.advance(1)
	return
}

func f32Neg(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float32:
		frame.Stack.Push(&Value{IR.TypeF32, -a.Value().(float32)})
	default:
		panic("f32.neg parameter invalid")
	}
	frame.advance(1)
	return
}

func f64Neg(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float64:
		frame.Stack.Push(&Value{IR.TypeF64, -a.Value().(float64)})
	default:
		panic("f64.neg parameter invalid")
	}
	frame.advance(1)
	return
}

func f32Ceil(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float32:
		frame.Stack.Push(&Value{IR.TypeF32, float32(math.Ceil(float64(a.Value().(float32))))})
	default:
		panic("f32.ceil parameter invalid")
	}
	frame.advance(1)
	return
}

func f64Ceil(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float64:
		frame.Stack.Push(&Value{IR.TypeF64, math.Ceil(a.Value().(float64))})
	default:
		panic("f64.ceil parameter invalid")
	}
	frame.advance(1)
	return
}

func f32Floor(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float32:
		frame.Stack.Push(&Value{IR.TypeF32, float32(math.Floor(float64(a.Value().(float32))))})
	default:
		panic("f32.ceil parameter invalid")
	}
	frame.advance(1)
	return
}

func f64Floor(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float64:
		frame.Stack.Push(&Value{IR.TypeF64, math.Floor(a.Value().(float64))})
	default:
		panic("f64.floor parameter invalid")
	}
	frame.advance(1)
	return
}

func f32Trunc(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float32:
		frame.Stack.Push(&Value{IR.TypeF32, float32(math.Trunc(float64(a.Value().(float32))))})
	default:
		panic("f32.trunc parameter invalid")
	}
	frame.advance(1)
	return
}

func f64Trunc(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float64:
		frame.Stack.Push(&Value{IR.TypeF64, math.Trunc(a.Value().(float64))})
	default:
		panic("f64.trunc parameter invalid")
	}
	frame.advance(1)
	return
}

func f32Nearest(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float32:
		frame.Stack.Push(&Value{IR.TypeF32, float32(math.RoundToEven(float64(a.Value().(float32))))})
	default:
		panic("f32.nearest parameter invalid")
	}
	frame.advance(1)
	return
}

func f64Nearest(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float64:
		frame.Stack.Push(&Value{IR.TypeF64, math.RoundToEven(a.Value().(float64))})
	default:
		panic("f64.nearest parameter invalid")
	}
	frame.advance(1)
	return
}

func f32Sqrt(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float32:
		frame.Stack.Push(&Value{IR.TypeF32, float32(math.Sqrt(float64(a.Value().(float32))))})
	default:
		panic("f32.sqrt parameter invalid")
	}
	frame.advance(1)
	return
}

func f64Sqrt(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float64:
		frame.Stack.Push(&Value{IR.TypeF64, math.Sqrt(a.Value().(float64))})
	default:
		panic("f64.sqrt parameter invalid")
	}
	frame.advance(1)
	return
}

func f32copySign(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	b, a, err := pop2(vm, frame)
	if err != nil {
		panic(err)
	}
	_, ok := a.Value().(float32)
	if !ok {
		panic("f32.copysign parameter invalid")
	}
	_, ok = b.Value().(float32)
	if !ok {
		panic("f32.copysign parameter invalid")
	}

	ret := math.Copysign(float64(a.Value().(float32)), float64(b.Value().(float32)))
	frame.Stack.Push(&Value{IR.TypeF32, float32(ret)})
	frame.advance(1)

	return
}

func f64copySign(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	b, a, err := pop2(vm, frame)
	if err != nil {
		panic(err)
	}
	_, ok := a.Value().(float64)
	if !ok {
		panic("f64.copysign parameter invalid")
	}
	_, ok = b.Value().(float64)
	if !ok {
		panic("f64.copysign parameter invalid")
	}

	ret := math.Copysign(float64(a.Value().(float64)), float64(b.Value().(float64)))
	frame.Stack.Push(&Value{IR.TypeF64, float64(ret)})
	frame.advance(1)

	return
}

func promoteF32ToF64(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float32:
		frame.Stack.Push(&Value{IR.TypeF64, float64(a.Value().(float32))})
	default:
		panic("f64.promote/f32 parameter invalid")
	}
	frame.advance(1)
	return
}
