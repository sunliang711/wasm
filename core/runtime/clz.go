package runtime

import (
	"wasm/core/IR"
	"wasm/utils"
)

const (
	I32_CLZ byte = iota
	I64_CLZ byte = iota
)

func clz(vm *VM, frame *Frame, clzType byte) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch clzType {
	case I32_CLZ:
		switch a.Value().(type) {
		case int32:
			v := a.Value().(int32)
			frame.Stack.Push(&Value{IR.TypeI32, doClz32(v)})
		case uint32:
			v := a.Value().(uint32)
			frame.Stack.Push(&Value{IR.TypeI32, doClz32(int32(v))})
		default:
			vm.panic("i32.clz parameter not (u)int32")
		}
	case I64_CLZ:
		switch a.Value().(type) {
		case int64:
			v := a.Value().(int64)
			frame.Stack.Push(&Value{IR.TypeI64, doClz64(v)})
		case uint64:
			v := a.Value().(uint64)
			frame.Stack.Push(&Value{IR.TypeI64, doClz64(int64(v))})
		default:
			vm.panic("i64.clz parameter not (u)int64")
		}
	default:
		vm.panic("clz() parameter 'clzType' invalid")
	}
	frame.advance(1)
	return
}

func doClz32(v int32) int32 {
	var ret int32
	for i := 31; i >= 0; i-- {
		shiftV := int32(1) << uint32(i)
		if v&shiftV != 0 {
			break
		}
		ret += 1
	}
	return ret
}

func doClz64(v int64) int64 {
	var ret int64
	for i := 63; i >= 0; i-- {
		shiftV := int64(1) << uint64(i)
		if v&shiftV != 0 {
			break
		}
		ret += 1
	}
	return ret
}
