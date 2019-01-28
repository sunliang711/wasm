package runtime

import (
	"wasm/core/IR"
	"wasm/utils"
)

const (
	I32_CTZ byte = iota
	I64_CTZ byte = iota
)

func ctz(vm *VM, frame *Frame, ctzType byte) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch ctzType {
	case I32_CTZ:
		switch a.Value().(type) {
		case int32:
			v := a.Value().(int32)
			frame.Stack.Push(&Value{IR.TypeI32, doCtz32(v)})
		case uint32:
			v := a.Value().(uint32)
			frame.Stack.Push(&Value{IR.TypeI32, doCtz32(int32(v))})
		default:
			panic("i32.ctz parameter not (u)int32")
		}
	case I64_CTZ:
		switch a.Value().(type) {
		case int64:
			v := a.Value().(int64)
			frame.Stack.Push(&Value{IR.TypeI64, doCtz64(v)})
		case uint64:
			v := a.Value().(uint64)
			frame.Stack.Push(&Value{IR.TypeI64, doCtz64(int64(v))})
		default:
			panic("i64.ctz parameter not (u)int64")
		}
	default:
		panic("ctz() parameter 'ctzType' invalid")
	}
	frame.advance(1)
	return
}

func doCtz32(v int32) (ret int32) {
	for i := 0; i < 32; i++ {
		shiftV := int32(1) << uint32(i)
		if v&shiftV != 0 {
			break
		}
		ret += 1
	}
	return
}

func doCtz64(v int64) (ret int64) {
	for i := 0; i < 64; i++ {
		shiftV := int64(1) << uint64(i)
		if v&shiftV != 0 {
			break
		}
		ret += 1
	}
	return
}
