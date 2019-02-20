package runtime

import (
	"wasm/core/IR"
	"wasm/utils"
)

const (
	SI32Extend byte = iota
	UI32Extend
)

func i64Extend(vm *WasmInterpreter, frame *Frame, extendType byte) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch extendType {
	case SI32Extend:
		switch a.Value().(type) {
		case int32:
			frame.Stack.Push(&Value{IR.TypeI64, int64(a.Value().(int32))})
		default:
			panic("i64.extend_s/i32 parameter invalid")
		}

	case UI32Extend:
		switch a.Value().(type) {
		case uint32:
			frame.Stack.Push(&Value{IR.TypeI64, int64(a.Value().(uint32))})
		default:
			panic("i64.extend_u/i32 parameter invalid")
		}
	default:
		panic("i64Extend extend type invalid")
	}

	frame.advance(1)
	return
}
