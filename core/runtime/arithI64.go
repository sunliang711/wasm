package runtime

import (
	"math/bits"
	"wasm/core/IR"
	"wasm/types"
)

func i64_arith(vm *VM, frame *Frame, arithType byte) {
	if frame.Stack.Len() < 2 {
		vm.panic(types.ErrStackSizeErr)
	}
	b, _ := frame.Stack.Pop()
	a, _ := frame.Stack.Pop()
	switch arithType {
	case ARITH_ADD:
		//int64 + int64  = int64
		//uint64 + uint64 = uint64
		switch a.Value().(type) {
		case int64:
			_, ok := b.Value().(int64)
			if !ok {
				vm.panic("i64.add int64 + non-int64")
			}
			result := a.Value().(int64) + b.Value().(int64)
			frame.Stack.Push(&Value{IR.TypeI64, result})

		case uint64:
			_, ok := b.Value().(uint64)
			if !ok {
				vm.panic("i64.add uint64 + non-uint64")
			}
			result := a.Value().(uint64) + b.Value().(uint64)
			frame.Stack.Push(&Value{IR.TypeI64, result})
		default:
			vm.panic("i64.add parameter not match")
		}

	case ARITH_SUB:
		//int64 - int64 = int64
		//uint64 - uint64 = uint64
		switch a.Value().(type) {
		case int64:
			_, ok := b.Value().(int64)
			if !ok {
				vm.panic("i64.sub int64 - non-int64")
			}
			result := a.Value().(int64) - b.Value().(int64)
			frame.Stack.Push(&Value{IR.TypeI64, result})

		case uint64:
			_, ok := b.Value().(uint64)
			if !ok {
				vm.panic("i64.sub uint64 - non-uint64")
			}
			result := a.Value().(uint64) - b.Value().(uint64)
			frame.Stack.Push(&Value{IR.TypeI64, result})
		default:
			vm.panic("i64.sub parameter not match")
		}
	case ARITH_MUL:
		//int64 * int64 = int64
		//uint64 * uint64 = uint64
		switch a.Value().(type) {
		case int64:
			_, ok := b.Value().(int64)
			if !ok {
				vm.panic("i64.mul int64 * non-int64")
			}
			result := a.Value().(int64) * b.Value().(int64)
			frame.Stack.Push(&Value{IR.TypeI64, result})

		case uint64:
			_, ok := b.Value().(uint64)
			if !ok {
				vm.panic("i64.mul uint64 * non-uint64")
			}
			result := a.Value().(uint64) * b.Value().(uint64)
			frame.Stack.Push(&Value{IR.TypeI64, result})
		default:
			vm.panic("i64.mul parameter not match")
		}
	case ARITH_DIV:
		//int64 / int64 = int64
		//uint64 / uint64 = uint64
		switch a.Value().(type) {
		case int64:
			_, ok := b.Value().(int64)
			if !ok {
				vm.panic("i64.div int64 / non-int64")
			}
			if b.Value().(int64) == 0 {
				vm.panic("Divided by zero")
			}
			result := a.Value().(int64) / b.Value().(int64)
			frame.Stack.Push(&Value{IR.TypeI64, result})

		case uint64:
			_, ok := b.Value().(uint64)
			if !ok {
				vm.panic("i64.div uint64 / non-uint64")
			}
			if b.Value().(uint64) == 0 {
				vm.panic("Divided by zero")
			}
			result := a.Value().(uint64) / b.Value().(uint64)
			frame.Stack.Push(&Value{IR.TypeI64, result})
		default:
			vm.panic("i64.div parameter not match")
		}

	case ARITH_MIN:
		//int64 min int64 = int64
		//uint64 min uint64 = uint64
		switch a.Value().(type) {
		case int64:
			_, ok := b.Value().(int64)
			if !ok {
				vm.panic("i64.min int64 min non-int64")
			}
			aVal := a.Value().(int64)
			bVal := b.Value().(int64)
			if aVal < bVal {
				frame.Stack.Push(&Value{IR.TypeI64, aVal})
			} else {
				frame.Stack.Push(&Value{IR.TypeI64, bVal})
			}

		case uint64:
			_, ok := b.Value().(uint64)
			if !ok {
				vm.panic("i64.min uint64 min non-uint64")
			}
			aVal := a.Value().(uint64)
			bVal := b.Value().(uint64)
			if aVal < bVal {
				frame.Stack.Push(&Value{IR.TypeI64, aVal})
			} else {
				frame.Stack.Push(&Value{IR.TypeI64, bVal})
			}

		default:
			vm.panic("i64.min parameter not match")
		}
	case ARITH_MAX:
		//int64 max int64 = int64
		//uint64 max uint64 = uint64
		switch a.Value().(type) {
		case int64:
			_, ok := b.Value().(int64)
			if !ok {
				vm.panic("i64.max int64 max non-int64")
			}
			aVal := a.Value().(int64)
			bVal := b.Value().(int64)
			if aVal < bVal {
				frame.Stack.Push(&Value{IR.TypeI64, bVal})
			} else {
				frame.Stack.Push(&Value{IR.TypeI64, aVal})
			}

		case uint64:
			_, ok := b.Value().(uint64)
			if !ok {
				vm.panic("i64.max uint64 max non-uint64")
			}
			aVal := a.Value().(uint64)
			bVal := b.Value().(uint64)
			if aVal < bVal {
				frame.Stack.Push(&Value{IR.TypeI64, bVal})
			} else {
				frame.Stack.Push(&Value{IR.TypeI64, aVal})
			}

		default:
			vm.panic("i64.max parameter not match")
		}
	case ARITH_REM:
		//int64 / int64 = int64
		//uint64 / uint64 = uint64
		switch a.Value().(type) {
		case int64:
			_, ok := b.Value().(int64)
			if !ok {
				vm.panic("i64.rem int64 rem non-int64")
			}
			if b.Value().(int64) == 0 {
				vm.panic("Divided by zero")
			}
			result := a.Value().(int64) % b.Value().(int64)
			frame.Stack.Push(&Value{IR.TypeI64, result})

		case uint64:
			_, ok := b.Value().(uint64)
			if !ok {
				vm.panic("i64.rem uint64 rem non-uint64")
			}
			if b.Value().(uint64) == 0 {
				vm.panic("Divided by zero")
			}
			result := a.Value().(uint64) % b.Value().(uint64)
			frame.Stack.Push(&Value{IR.TypeI64, result})
		default:
			vm.panic("i64.rem parameter not match")
		}

	case ARITH_AND:
		//int64 and int64 = int64
		//uint64 and uint64 = uint64
		switch a.Value().(type) {
		case int64:
			_, ok := b.Value().(int64)
			if !ok {
				vm.panic("i64.and int64 and non-int64")
			}
			result := a.Value().(int64) & b.Value().(int64)
			frame.Stack.Push(&Value{IR.TypeI64, result})

		case uint64:
			_, ok := b.Value().(uint64)
			if !ok {
				vm.panic("i64.and uint64 and non-uint64")
			}
			result := a.Value().(uint64) & b.Value().(uint64)
			frame.Stack.Push(&Value{IR.TypeI64, result})
		default:
			vm.panic("i64.and parameter not match")
		}
	case ARITH_OR:
		//int64 or int64 = int64
		//uint64 or uint64 = uint64
		switch a.Value().(type) {
		case int64:
			_, ok := b.Value().(int64)
			if !ok {
				vm.panic("i64.or int64 or non-int64")
			}
			result := a.Value().(int64) | b.Value().(int64)
			frame.Stack.Push(&Value{IR.TypeI64, result})

		case uint64:
			_, ok := b.Value().(uint64)
			if !ok {
				vm.panic("i64.or uint64 or non-uint64")
			}
			result := a.Value().(uint64) | b.Value().(uint64)
			frame.Stack.Push(&Value{IR.TypeI64, result})
		default:
			vm.panic("i64.or parameter not match")
		}
	case ARITH_XOR:
		//int64 xor int64 = int64
		//uint64 xor uint64 = uint64
		switch a.Value().(type) {
		case int64:
			_, ok := b.Value().(int64)
			if !ok {
				vm.panic("i64.xor int64 xor non-int64")
			}
			result := a.Value().(int64) ^ b.Value().(int64)
			frame.Stack.Push(&Value{IR.TypeI64, result})

		case uint64:
			_, ok := b.Value().(uint64)
			if !ok {
				vm.panic("i64.xor uint64 xor non-uint64")
			}
			result := a.Value().(uint64) ^ b.Value().(uint64)
			frame.Stack.Push(&Value{IR.TypeI64, result})
		default:
			vm.panic("i64.xor parameter not match")
		}
	case ARITH_SHL:
		length := uint(0)
		switch b.Value().(type) {
		case int64:
			length = uint(b.Value().(int64))
		case uint64:
			length = uint(b.Value().(uint64))
		default:
			vm.panic("i64.shl shift length not (u)i64")
		}

		switch a.Value().(type) {
		case int64:
			result := a.Value().(int64) << length
			frame.Stack.Push(&Value{IR.TypeI64, result})
		case uint64:
			result := a.Value().(uint64) << length
			frame.Stack.Push(&Value{IR.TypeI64, result})
		default:
			vm.panic("i64.shl oprand not (u)i64")
		}
	case ARITH_SHR: //sign
		length := uint(0)
		switch b.Value().(type) {
		case int64:
			length = uint(b.Value().(int64))
		case uint64:
			length = uint(b.Value().(uint64))
		default:
			vm.panic("i64.shl shift length not (u)i64")
		}

		switch a.Value().(type) {
		case int64:
			result := a.Value().(int64) >> length
			frame.Stack.Push(&Value{IR.TypeI64, result})
		case uint64:
			result := a.Value().(uint64) >> length
			frame.Stack.Push(&Value{IR.TypeI64, result})
		default:
			vm.panic("i64.shr oprand not (u)i64")
		}
	case ARITH_ROTL:
		var rotlLength int
		switch b.Value().(type) {
		case int64:
			rotlLength = int(b.Value().(int64))
		case uint64:
			rotlLength = int(b.Value().(uint64))
		default:
			vm.panic("i64.rotl length invalid")
		}

		switch a.Value().(type) {
		case int64:
			result := int64(bits.RotateLeft64(uint64(a.Value().(int64)), rotlLength))
			frame.Stack.Push(&Value{IR.TypeI64, result})
		case uint64:
			result := bits.RotateLeft64(a.Value().(uint64), rotlLength)
			frame.Stack.Push(&Value{IR.TypeI64, result})
		default:
			vm.panic("i64.rotl oprand not (u)i64")
		}
	case ARITH_ROTR:
		var rotlLength int
		switch b.Value().(type) {
		case int64:
			rotlLength = int(b.Value().(int64))
		case uint64:
			rotlLength = int(b.Value().(uint64))
		default:
			vm.panic("i64.rotr length invalid")
		}

		switch a.Value().(type) {
		case int64:
			result := int64(bits.RotateLeft64(uint64(a.Value().(int64)), -rotlLength))
			frame.Stack.Push(&Value{IR.TypeI64, result})
		case uint64:
			result := bits.RotateLeft64(a.Value().(uint64), -rotlLength)
			frame.Stack.Push(&Value{IR.TypeI64, result})
		default:
			vm.panic("i64.rotr oprand not (u)i64")
		}
	}
}