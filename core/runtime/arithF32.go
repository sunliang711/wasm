package runtime

import (
	"wasm/core/IR"
	"wasm/utils"
)

func f32_arith(vm *VM, frame *Frame, arithType byte) (err error) {
	defer utils.CatchError(&err)
	b, a, err := pop2(vm, frame)
	if err != nil {
		panic(err)
	}
	switch arithType {
	case ARITH_ADD:
		switch a.Value().(type) {
		case float32:
			//f32 + f32 = f32
			_, ok := b.Value().(float32)
			if !ok {
				vm.panic("f32 + non-f32")
			}
			result := a.Value().(float32) + b.Value().(float32)
			frame.Stack.Push(&Value{IR.TypeF32, result})
		default:
			vm.panic("f32.add parameter not match")
		}
	case ARITH_SUB:
		switch a.Value().(type) {
		case float32:
			// f32 - f32 = f32
			_, ok := b.Value().(float32)
			if !ok {
				vm.panic("f32 + non-f32")
			}
			result := a.Value().(float32) - b.Value().(float32)
			frame.Stack.Push(&Value{IR.TypeF32, result})
		default:
			vm.panic("f32.sub parameter not match")
		}
	case ARITH_MUL:
		switch a.Value().(type) {
		case float32:
			// f32 - f32 = f32
			_, ok := b.Value().(float32)
			if !ok {
				vm.panic("f32 + non-f32")
			}
			result := a.Value().(float32) * b.Value().(float32)
			frame.Stack.Push(&Value{IR.TypeF32, result})
		default:
			vm.panic("f32.mul parameter not match")
		}
	case ARITH_DIV:
		switch a.Value().(type) {
		case float32:
			// f32 - f32 = f32
			_, ok := b.Value().(float32)
			if !ok {
				vm.panic("f32 + non-f32")
			}
			result := a.Value().(float32) / b.Value().(float32)
			frame.Stack.Push(&Value{IR.TypeF32, result})
		default:
			vm.panic("f32.div parameter not match")
		}
	case ARITH_MIN:
		switch a.Value().(type) {
		case float32:
			// f32 min f32 = f32
			_, ok := b.Value().(float32)
			if !ok {
				vm.panic("f32 min non-f32")
			}
			aVal := a.Value().(float32)
			bVal := b.Value().(float32)
			var result float32
			if aVal < bVal {
				result = aVal
			} else {
				result = bVal
			}
			frame.Stack.Push(&Value{IR.TypeF32, result})
		default:
			vm.panic("f32.min parameter not match")
		}
	case ARITH_MAX:
		switch a.Value().(type) {
		case float32:
			// f32 max f32 = f32
			_, ok := b.Value().(float32)
			if !ok {
				vm.panic("f32 max non-f32")
			}
			aVal := a.Value().(float32)
			bVal := b.Value().(float32)
			var result float32
			if aVal > bVal {
				result = aVal
			} else {
				result = bVal
			}
			frame.Stack.Push(&Value{IR.TypeF32, result})
		default:
			vm.panic("f32.max parameter not match")
		}
	}
	frame.advance(1)
	return
}
