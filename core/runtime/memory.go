package runtime

import (
	"encoding/binary"
	"math"
	"wasm/core/IR"
	"wasm/types"
)

func i32_load(vm *VM, frame *Frame, offset uint32, numBytes int, isSignExtend bool) {
	if frame.Stack.Empty() {
		vm.panic(types.ErrStackSizeErr)
	}
	baseVal, _ := frame.Stack.Pop()
	base := baseVal.Value().(int32)
	addr := base + int32(offset)
	if addr < 0 {
		addr += int32(len(vm.Memory))
	}

	switch numBytes {
	case 1:
		val := uint8(vm.Memory[addr])
		if isSignExtend {
			frame.Stack.Push(&Value{IR.TypeI32, int32(val)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, uint32(val)})
		}
	case 2:
		val := binary.LittleEndian.Uint16(vm.Memory[addr : addr+2])
		if isSignExtend {
			frame.Stack.Push(&Value{IR.TypeI32, int32(val)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI32, uint32(val)})
		}
	case 4:
		val := binary.LittleEndian.Uint32(vm.Memory[addr : addr+4])
		frame.Stack.Push(&Value{IR.TypeI32, int32(val)})
	default:
		vm.panic("i32_load numBytes error")
	}
}

func i64_load(vm *VM, frame *Frame, offset uint32, numBytes int, isSignExtend bool) {
	if frame.Stack.Empty() {
		vm.panic(types.ErrStackSizeErr)
	}
	baseVal, _ := frame.Stack.Pop()
	base := baseVal.Value().(int32)
	addr := base + int32(offset)
	if addr < 0 {
		addr += int32(len(vm.Memory))
	}

	switch numBytes {
	case 1:
		val := uint8(vm.Memory[addr])
		if isSignExtend {
			frame.Stack.Push(&Value{IR.TypeI64, int64(val)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI64, uint64(val)})
		}
	case 2:
		val := binary.LittleEndian.Uint16(vm.Memory[addr : addr+2])
		if isSignExtend {
			frame.Stack.Push(&Value{IR.TypeI64, int64(val)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI64, uint64(val)})
		}
	case 4:
		val := binary.LittleEndian.Uint32(vm.Memory[addr : addr+4])
		if isSignExtend {
			frame.Stack.Push(&Value{IR.TypeI64, int64(val)})
		} else {
			frame.Stack.Push(&Value{IR.TypeI64, uint64(val)})
		}
	case 8:
		val := binary.LittleEndian.Uint64(vm.Memory[addr : addr+8])
		frame.Stack.Push(&Value{IR.TypeI64, int64(val)})
	default:
		vm.panic("i64_load numBytes error")
	}
}

func i32_store(vm *VM, frame *Frame, offset uint32, numBytes int) {
	if frame.Stack.Len() < 2 {
		vm.panic(types.ErrStackSizeErr)
	}
	valVal, _ := frame.Stack.Pop()
	baseVal, _ := frame.Stack.Pop()
	base := baseVal.Value().(int32)
	val := valVal.Value().(int32)
	addr := base + int32(offset)
	if addr < 0 {
		addr += int32(len(vm.Memory))
	}

	switch numBytes {
	case 1:
		vm.Memory[addr] = uint8(int8(val))
	case 2:
		binary.LittleEndian.PutUint16(vm.Memory[addr:addr+2], uint16(int16(val)))
	case 4:
		binary.LittleEndian.PutUint32(vm.Memory[addr:addr+4], uint32(val))
	default:
		vm.panic("i32_store numBytes error")

	}
}

func i64_store(vm *VM, frame *Frame, offset uint32, numBytes int) {
	if frame.Stack.Len() < 2 {
		vm.panic(types.ErrStackSizeErr)
	}
	valVal, _ := frame.Stack.Pop()
	baseVal, _ := frame.Stack.Pop()
	base := baseVal.Value().(int32)
	val := valVal.Value().(int64)
	addr := base + int32(offset)
	if addr < 0 {
		addr += int32(len(vm.Memory))
	}

	switch numBytes {
	case 1:
		vm.Memory[addr] = uint8(int8(val))
	case 2:
		binary.LittleEndian.PutUint16(vm.Memory[addr:addr+2], uint16(int16(val)))
	case 4:
		binary.LittleEndian.PutUint32(vm.Memory[addr:addr+4], uint32(int32(val)))
	case 8:
		binary.LittleEndian.PutUint64(vm.Memory[addr:addr+8], uint64(val))
	default:
		vm.panic("i64_store numBytes error")

	}
}

func float_store(vm *VM, frame *Frame, offset uint32, numBytes int) {

	if frame.Stack.Len() < 2 {
		vm.panic(types.ErrStackSizeErr)
	}
	valVal, _ := frame.Stack.Pop()
	baseVal, _ := frame.Stack.Pop()
	base := baseVal.Value().(int32)
	addr := base + int32(offset)
	if addr < 0 {
		addr += int32(len(vm.Memory))
	}
	switch numBytes {
	case 4:
		val := valVal.Value().(float32)
		binary.LittleEndian.PutUint32(vm.Memory[addr:addr+4], math.Float32bits(val))
	case 8:
		val := valVal.Value().(float64)
		binary.LittleEndian.PutUint64(vm.Memory[addr:addr+8], math.Float64bits(val))
	default:
		vm.panic("float_store numBytes error")
	}
}

func float_load(vm *VM, frame *Frame, offset uint32, numBytes int) {

	if frame.Stack.Empty() {
		vm.panic(types.ErrStackSizeErr)
	}
	baseVal, _ := frame.Stack.Pop()
	base := baseVal.Value().(int32)
	addr := base + int32(offset)
	if addr < 0 {
		addr += int32(len(vm.Memory))
	}

	switch numBytes {
	case 4:
		val := math.Float32frombits(binary.LittleEndian.Uint32(vm.Memory[addr : addr+4]))
		frame.Stack.Push(&Value{IR.TypeF32, val})
	case 8:
		val := math.Float64frombits(binary.LittleEndian.Uint64(vm.Memory[addr : addr+8]))
		frame.Stack.Push(&Value{IR.TypeF64, val})
	default:
		vm.panic("float_load numBytes error")
	}
}
