package runtime

import (
	"wasm/core/IR"
	"wasm/utils"
)

func f64_arith(vm *VM, frame *Frame, arithType byte) (err error) {
	defer utils.CatchError(&err)
	b, a, err := pop2(vm, frame)
	if err != nil {
		panic(err)
	}
	switch arithType {
	case ARITH_ADD:
		switch a.Value().(type) {
		case float64:
			//f64 + f64 = f64
			_, ok := b.Value().(float64)
			if !ok {
				vm.panic("f64 + non-f64")
			}
			result := a.Value().(float64) + b.Value().(float64)
			frame.Stack.Push(&Value{IR.TypeF64, result})
		default:
			vm.panic("f64.add parameter not match")
		}
	case ARITH_SUB:
		switch a.Value().(type) {
		case float64:
			// f64 - f64 = f64
			_, ok := b.Value().(float64)
			if !ok {
				vm.panic("f64 + non-f64")
			}
			result := a.Value().(float64) - b.Value().(float64)
			frame.Stack.Push(&Value{IR.TypeF64, result})
		default:
			vm.panic("f64.sub parameter not match")
		}
	case ARITH_MUL:
		switch a.Value().(type) {
		case float64:
			// f64 - f64 = f64
			_, ok := b.Value().(float64)
			if !ok {
				vm.panic("f64 + non-f64")
			}
			result := a.Value().(float64) * b.Value().(float64)
			frame.Stack.Push(&Value{IR.TypeF64, result})
		default:
			vm.panic("f64.mul parameter not match")
		}
	case ARITH_DIV:
		switch a.Value().(type) {
		case float64:
			// f64 - f64 = f64
			_, ok := b.Value().(float64)
			if !ok {
				vm.panic("f64 + non-f64")
			}
			result := a.Value().(float64) / b.Value().(float64)
			frame.Stack.Push(&Value{IR.TypeF64, result})
		default:
			vm.panic("f64.div parameter not match")
		}
	case ARITH_MIN:
		switch a.Value().(type) {
		case float64:
			// f64 min f64 = f64
			_, ok := b.Value().(float64)
			if !ok {
				vm.panic("f64 min non-f64")
			}
			aVal := a.Value().(float64)
			bVal := b.Value().(float64)
			var result float64
			if aVal < bVal {
				result = aVal
			} else {
				result = bVal
			}
			frame.Stack.Push(&Value{IR.TypeF64, result})
		default:
			vm.panic("f64.min parameter not match")
		}
	case ARITH_MAX:
		switch a.Value().(type) {
		case float64:
			// f64 max f64 = f64
			_, ok := b.Value().(float64)
			if !ok {
				vm.panic("f64 max non-f64")
			}
			aVal := a.Value().(float64)
			bVal := b.Value().(float64)
			var result float64
			if aVal > bVal {
				result = aVal
			} else {
				result = bVal
			}
			frame.Stack.Push(&Value{IR.TypeF64, result})
		default:
			vm.panic("f64.max parameter not match")
		}
	}
	frame.advance(1)
	return
}
