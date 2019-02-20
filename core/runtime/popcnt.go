package runtime

import (
	"wasm/core/IR"
	"wasm/utils"
)

const (
	I32_POPCNT byte = iota
	I64_POPCNT
)

func popcnt(vm *WasmInterpreter, frame *Frame, popCntType byte) (err error) {
	defer utils.CatchError(&err)

	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch popCntType {
	case I32_POPCNT:
		switch a.Value().(type) {
		case int32:
			frame.Stack.Push(&Value{IR.TypeI32, doPopcnt32(a.Value().(int32))})
		case uint32:
			frame.Stack.Push(&Value{IR.TypeI32, doPopcnt32(int32(a.Value().(uint32)))})
		default:
			vm.panic("i32.popcnt parameter not valid")
		}
	case I64_POPCNT:
		switch a.Value().(type) {
		case int64:
			frame.Stack.Push(&Value{IR.TypeI64, doPopcnt64(a.Value().(int64))})
		case uint64:
			frame.Stack.Push(&Value{IR.TypeI64, doPopcnt64(int64(a.Value().(uint64)))})
		default:
			vm.panic("i32.popcnt parameter not valid")
		}
	default:
		vm.panic("popcnt() parameter invalid")
	}
	frame.advance(1)
	return
}

func doPopcnt32(v int32) (ret int32) {
	for i := 0; i < 32; i++ {
		shiftV := int32(1) << uint32(i)
		if shiftV&v != 0 {
			ret += 1
		}
	}
	return
}

func doPopcnt64(v int64) (ret int64) {
	for i := 0; i < 64; i++ {
		shiftV := int64(1) << uint64(i)
		if shiftV&v != 0 {
			ret += 1
		}
	}
	return
}
