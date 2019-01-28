package runtime

import (
	"fmt"
	"wasm/core/IR"
	"wasm/types"
)

func pop1(vm *VM, frame *Frame) (IR.InterfaceValue, error) {
	if frame.Stack.Len() < 1 {
		return nil, fmt.Errorf(types.ErrStackSizeErr)
	}
	ret, _ := frame.Stack.Pop()
	return ret, nil
}

func pop2(vm *VM, frame *Frame) (IR.InterfaceValue, IR.InterfaceValue, error) {
	if frame.Stack.Len() < 2 {
		return nil, nil, fmt.Errorf(types.ErrStackSizeErr)
	}
	a, _ := frame.Stack.Pop()
	b, _ := frame.Stack.Pop()
	return a, b, nil
}

func IsZero(v IR.InterfaceValue) bool {
	switch v.Type() {
	case IR.TypeI32:
		switch v.Value().(type) {
		case int32:
			return v.Value().(int32) == 0
		case uint32:
			return v.Value().(uint32) == 0
		}
	case IR.TypeI64:
		switch v.Value().(type) {
		case int64:
			return v.Value().(int64) == 0
		case uint64:
			return v.Value().(uint64) == 0
		}
	case IR.TypeF32:
		return v.Value().(float32) == 0
	case IR.TypeF64:
		return v.Value().(float64) == 0
	default:
		return false
	}
	return false
}
