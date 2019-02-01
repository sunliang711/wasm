package runtime

import (
	"math/bits"
	"wasm/core/IR"
	"wasm/types"
	"wasm/utils"
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

func i32_arith(vm *VM, frame *Frame, arithType byte) (err error) {
	defer utils.CatchError(&err)
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
	frame.advance(1)
	return
}

const (
	I32_EQZ byte = iota
	I64_EQZ
)

func eqz(vm *VM, frame *Frame, eqType byte) (err error) {
	defer utils.CatchError(&err)
	a, err := pop1(vm, frame)
	if err != nil {
		panic(err)
	}
	switch eqType {
	case I32_EQZ:
		if IsZero(a) {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		}
	case I64_EQZ:
		if IsZero(a) {
			frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
		}
	}
	frame.advance(1)
	return
}

func i64_arith(vm *VM, frame *Frame, arithType byte) (err error) {
	defer utils.CatchError(&err)
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
	frame.advance(1)
	return
}

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
