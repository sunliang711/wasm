package runtime

import (
	"encoding/binary"
	"math"
	"wasm/core/IR"
	"wasm/utils"
)

func i32Reinterpret(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float32:
		v, _ := a.Value().(float32)
		frame.Stack.Push(&Value{IR.TypeI32, math.Float32bits(v)})
	default:
		panic("i32.reinterpret/f32 parameter invalid")
	}
	frame.advance(1)
	return
}

func i64Reinterpret(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch a.Value().(type) {
	case float64:
		v, _ := a.Value().(float64)
		frame.Stack.Push(&Value{IR.TypeI64, math.Float64bits(v)})
	default:
		panic("i64.reinterpret/f64 parameter invalid")
	}
	frame.advance(1)
	return
}

func f32Reinterpret(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 4)
	switch a.Value().(type) {
	case int32:
		v, _ := a.Value().(int32)
		binary.LittleEndian.PutUint32(buf, uint32(v))
	case uint32:
		v, _ := a.Value().(uint32)
		binary.LittleEndian.PutUint32(buf, v)
	default:
		panic("f32.reinterpret/i32 parameter invalid")
	}
	ret, err := utils.Bytes2float32(buf, true)
	if err != nil {
		panic(err)
	}
	frame.Stack.Push(&Value{IR.TypeF32, ret})
	frame.advance(1)
	return
}

func f64Reinterpret(vm *VM, frame *Frame) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 8)
	switch a.Value().(type) {
	case int64:
		v, _ := a.Value().(int64)
		binary.LittleEndian.PutUint64(buf, uint64(v))
	case uint64:
		v, _ := a.Value().(uint64)
		binary.LittleEndian.PutUint64(buf, v)
	default:
		panic("f64.reinterpret/i64 parameter invalid")
	}
	ret, err := utils.Bytes2float64(buf, true)
	if err != nil {
		panic(err)
	}
	frame.Stack.Push(&Value{IR.TypeF64, ret})
	frame.advance(1)
	return
}
