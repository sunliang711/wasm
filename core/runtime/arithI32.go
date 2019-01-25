package runtime

import (
	"math/bits"
	"wasm/core/IR"
	"wasm/types"
)

const (
	ARITH_ADD byte = iota
	ARITH_SUB
	ARITH_MUL
	ARITH_DIV
	ARITH_MIN
	ARITH_MAX
	ARITH_REM
	ARITH_AND
	ARITH_OR
	ARITH_XOR
	ARITH_SHL
	ARITH_SHR
	ARITH_ROTL
	ARITH_ROTR
)

func i32_arith(vm *VM, frame *Frame, arithType byte) {
	if frame.Stack.Len() < 2 {
		vm.panic(types.ErrStackSizeErr)
	}
	b, _ := frame.Stack.Pop()
	a, _ := frame.Stack.Pop()
	switch arithType {
	case ARITH_ADD:
		//int32 + int32  = int32
		//uint32 + uint32 = uint32
		switch a.Value().(type) {
		case int32:
			_, ok := b.Value().(int32)
			if !ok {
				vm.panic("i32.add int32 + non-int32")
			}
			result := a.Value().(int32) + b.Value().(int32)
			frame.Stack.Push(&Value{IR.TypeI32, result})

		case uint32:
			_, ok := b.Value().(uint32)
			if !ok {
				vm.panic("i32.add uint32 + non-uint32")
			}
			result := a.Value().(uint32) + b.Value().(uint32)
			frame.Stack.Push(&Value{IR.TypeI32, result})
		default:
			vm.panic("i32.add parameter not match")
		}

	case ARITH_SUB:
		//int32 - int32 = int32
		//uint32 - uint32 = uint32
		switch a.Value().(type) {
		case int32:
			_, ok := b.Value().(int32)
			if !ok {
				vm.panic("i32.sub int32 - non-int32")
			}
			result := a.Value().(int32) - b.Value().(int32)
			frame.Stack.Push(&Value{IR.TypeI32, result})

		case uint32:
			_, ok := b.Value().(uint32)
			if !ok {
				vm.panic("i32.sub uint32 - non-uint32")
			}
			result := a.Value().(uint32) - b.Value().(uint32)
			frame.Stack.Push(&Value{IR.TypeI32, result})
		default:
			vm.panic("i32.sub parameter not match")
		}
	case ARITH_MUL:
		//int32 * int32 = int32
		//uint32 * uint32 = uint32
		switch a.Value().(type) {
		case int32:
			_, ok := b.Value().(int32)
			if !ok {
				vm.panic("i32.mul int32 * non-int32")
			}
			result := a.Value().(int32) * b.Value().(int32)
			frame.Stack.Push(&Value{IR.TypeI32, result})

		case uint32:
			_, ok := b.Value().(uint32)
			if !ok {
				vm.panic("i32.mul uint32 * non-uint32")
			}
			result := a.Value().(uint32) * b.Value().(uint32)
			frame.Stack.Push(&Value{IR.TypeI32, result})
		default:
			vm.panic("i32.mul parameter not match")
		}
	case ARITH_DIV:
		//int32 / int32 = int32
		//uint32 / uint32 = uint32
		switch a.Value().(type) {
		case int32:
			_, ok := b.Value().(int32)
			if !ok {
				vm.panic("i32.div int32 / non-int32")
			}
			if b.Value().(int32) == 0 {
				vm.panic("Divided by zero")
			}
			result := a.Value().(int32) / b.Value().(int32)
			frame.Stack.Push(&Value{IR.TypeI32, result})

		case uint32:
			_, ok := b.Value().(uint32)
			if !ok {
				vm.panic("i32.div uint32 / non-uint32")
			}
			if b.Value().(uint32) == 0 {
				vm.panic("Divided by zero")
			}
			result := a.Value().(uint32) / b.Value().(uint32)
			frame.Stack.Push(&Value{IR.TypeI32, result})
		default:
			vm.panic("i32.div parameter not match")
		}

	case ARITH_MIN:
		//int32 min int32 = int32
		//uint32 min uint32 = uint32
		switch a.Value().(type) {
		case int32:
			_, ok := b.Value().(int32)
			if !ok {
				vm.panic("i32.min int32 min non-int32")
			}
			aVal := a.Value().(int32)
			bVal := b.Value().(int32)
			if aVal < bVal {
				frame.Stack.Push(&Value{IR.TypeI32, aVal})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, bVal})
			}

		case uint32:
			_, ok := b.Value().(uint32)
			if !ok {
				vm.panic("i32.min uint32 min non-uint32")
			}
			aVal := a.Value().(uint32)
			bVal := b.Value().(uint32)
			if aVal < bVal {
				frame.Stack.Push(&Value{IR.TypeI32, aVal})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, bVal})
			}

		default:
			vm.panic("i32.min parameter not match")
		}
	case ARITH_MAX:
		//int32 max int32 = int32
		//uint32 max uint32 = uint32
		switch a.Value().(type) {
		case int32:
			_, ok := b.Value().(int32)
			if !ok {
				vm.panic("i32.max int32 max non-int32")
			}
			aVal := a.Value().(int32)
			bVal := b.Value().(int32)
			if aVal < bVal {
				frame.Stack.Push(&Value{IR.TypeI32, bVal})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, aVal})
			}

		case uint32:
			_, ok := b.Value().(uint32)
			if !ok {
				vm.panic("i32.max uint32 max non-uint32")
			}
			aVal := a.Value().(uint32)
			bVal := b.Value().(uint32)
			if aVal < bVal {
				frame.Stack.Push(&Value{IR.TypeI32, bVal})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, aVal})
			}

		default:
			vm.panic("i32.max parameter not match")
		}
	case ARITH_REM:
		//int32 / int32 = int32
		//uint32 / uint32 = uint32
		switch a.Value().(type) {
		case int32:
			_, ok := b.Value().(int32)
			if !ok {
				vm.panic("i32.rem int32 rem non-int32")
			}
			if b.Value().(int32) == 0 {
				vm.panic("Divided by zero")
			}
			result := a.Value().(int32) % b.Value().(int32)
			frame.Stack.Push(&Value{IR.TypeI32, result})

		case uint32:
			_, ok := b.Value().(uint32)
			if !ok {
				vm.panic("i32.rem uint32 rem non-uint32")
			}
			if b.Value().(uint32) == 0 {
				vm.panic("Divided by zero")
			}
			result := a.Value().(uint32) % b.Value().(uint32)
			frame.Stack.Push(&Value{IR.TypeI32, result})
		default:
			vm.panic("i32.rem parameter not match")
		}

	case ARITH_AND:
		//int32 and int32 = int32
		//uint32 and uint32 = uint32
		switch a.Value().(type) {
		case int32:
			_, ok := b.Value().(int32)
			if !ok {
				vm.panic("i32.and int32 and non-int32")
			}
			result := a.Value().(int32) & b.Value().(int32)
			frame.Stack.Push(&Value{IR.TypeI32, result})

		case uint32:
			_, ok := b.Value().(uint32)
			if !ok {
				vm.panic("i32.and uint32 and non-uint32")
			}
			result := a.Value().(uint32) & b.Value().(uint32)
			frame.Stack.Push(&Value{IR.TypeI32, result})
		default:
			vm.panic("i32.and parameter not match")
		}
	case ARITH_OR:
		//int32 or int32 = int32
		//uint32 or uint32 = uint32
		switch a.Value().(type) {
		case int32:
			_, ok := b.Value().(int32)
			if !ok {
				vm.panic("i32.or int32 or non-int32")
			}
			result := a.Value().(int32) | b.Value().(int32)
			frame.Stack.Push(&Value{IR.TypeI32, result})

		case uint32:
			_, ok := b.Value().(uint32)
			if !ok {
				vm.panic("i32.or uint32 or non-uint32")
			}
			result := a.Value().(uint32) | b.Value().(uint32)
			frame.Stack.Push(&Value{IR.TypeI32, result})
		default:
			vm.panic("i32.or parameter not match")
		}
	case ARITH_XOR:
		//int32 xor int32 = int32
		//uint32 xor uint32 = uint32
		switch a.Value().(type) {
		case int32:
			_, ok := b.Value().(int32)
			if !ok {
				vm.panic("i32.xor int32 xor non-int32")
			}
			result := a.Value().(int32) ^ b.Value().(int32)
			frame.Stack.Push(&Value{IR.TypeI32, result})

		case uint32:
			_, ok := b.Value().(uint32)
			if !ok {
				vm.panic("i32.xor uint32 xor non-uint32")
			}
			result := a.Value().(uint32) ^ b.Value().(uint32)
			frame.Stack.Push(&Value{IR.TypeI32, result})
		default:
			vm.panic("i32.xor parameter not match")
		}
	case ARITH_SHL:
		length := uint(0)
		switch b.Value().(type) {
		case int32:
			length = uint(b.Value().(int32))
		case uint32:
			length = uint(b.Value().(uint32))
		default:
			vm.panic("i32.shl shift length not (u)i32")
		}

		switch a.Value().(type) {
		case int32:
			result := a.Value().(int32) << length
			frame.Stack.Push(&Value{IR.TypeI32, result})
		case uint32:
			result := a.Value().(uint32) << length
			frame.Stack.Push(&Value{IR.TypeI32, result})
		default:
			vm.panic("i32.shl oprand not (u)i32")
		}
	case ARITH_SHR: //sign
		length := uint(0)
		switch b.Value().(type) {
		case int32:
			length = uint(b.Value().(int32))
		case uint32:
			length = uint(b.Value().(uint32))
		default:
			vm.panic("i32.shl shift length not (u)i32")
		}

		switch a.Value().(type) {
		case int32:
			result := a.Value().(int32) >> length
			frame.Stack.Push(&Value{IR.TypeI32, result})
		case uint32:
			result := a.Value().(uint32) >> length
			frame.Stack.Push(&Value{IR.TypeI32, result})
		default:
			vm.panic("i32.shr oprand not (u)i32")
		}
	case ARITH_ROTL:
		var rotlLength int
		switch b.Value().(type) {
		case int32:
			rotlLength = int(b.Value().(int32))
		case uint32:
			rotlLength = int(b.Value().(uint32))
		default:
			vm.panic("i32.rotl length invalid")
		}

		switch a.Value().(type) {
		case int32:
			result := int32(bits.RotateLeft32(uint32(a.Value().(int32)), rotlLength))
			frame.Stack.Push(&Value{IR.TypeI32, result})
		case uint32:
			result := bits.RotateLeft32(a.Value().(uint32), rotlLength)
			frame.Stack.Push(&Value{IR.TypeI32, result})
		default:
			vm.panic("i32.rotl oprand not (u)i32")
		}
	case ARITH_ROTR:
		var rotlLength int
		switch b.Value().(type) {
		case int32:
			rotlLength = int(b.Value().(int32))
		case uint32:
			rotlLength = int(b.Value().(uint32))
		default:
			vm.panic("i32.rotr length invalid")
		}

		switch a.Value().(type) {
		case int32:
			result := int32(bits.RotateLeft32(uint32(a.Value().(int32)), -rotlLength))
			frame.Stack.Push(&Value{IR.TypeI32, result})
		case uint32:
			result := bits.RotateLeft32(a.Value().(uint32), -rotlLength)
			frame.Stack.Push(&Value{IR.TypeI32, result})
		default:
			vm.panic("i32.rotr oprand not (u)i32")
		}
	}
}
