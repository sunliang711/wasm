package runtime

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"wasm/core/IR"
	"wasm/types"
	"wasm/utils"
)

func IsZero(v IR.InterfaceValue) bool {
	switch v.Type() {
	case IR.TypeI32:
		return v.Value().(int32) == 0
	case IR.TypeI64:
		return v.Value().(int64) == 0
	case IR.TypeF32:
		return v.Value().(float32) == 0
	case IR.TypeF64:
		return v.Value().(float64) == 0
	default:
		return false
	}
}

func (vm *VM) Run(functionNameOrID interface{}, params ...interface{}) (err error) {
	//TODO check all type assertion
	defer utils.CatchError(&err)
	frame := vm.GetCurrentFrame()
	if frame == nil {
		panic("Frame stack overflow.")
	}
	funcIndex := 0
	switch functionNameOrID.(type) {
	case string:
		funcIndex, err = vm.Module.GetFuncIndexWithName(functionNameOrID.(string))
		if err != nil {
			vm.panic(err)
		}
	case int:
		funcIndex = functionNameOrID.(int)
	default:
		panic("functionNameOrID must be a string or an int")
	}

	Params := make([]IR.InterfaceValue, 0)
	//int8,int16,int32 --> int32;  uint8,uint16,uint32 --> uint32
	//int,int64 --> int64; uint,uint64 --> uint64
	for _, p := range params {
		kind := reflect.TypeOf(p).Kind()
		switch kind {
		case reflect.Int8:
			Params = append(Params, &Value{IR.TypeI32, int32(p.(int8))})
		case reflect.Uint8:
			Params = append(Params, &Value{IR.TypeI32, uint32(p.(uint8))})
		case reflect.Int16:
			Params = append(Params, &Value{IR.TypeI32, int32(p.(int16))})
		case reflect.Uint16:
			Params = append(Params, &Value{IR.TypeI32, uint32(p.(uint16))})
		case reflect.Int32:
			Params = append(Params, &Value{IR.TypeI32, p.(int32)})
		case reflect.Uint32:
			Params = append(Params, &Value{IR.TypeI32, p.(uint32)})

		case reflect.Int:
			Params = append(Params, &Value{IR.TypeI64, int64(p.(int))})
		case reflect.Uint:
			Params = append(Params, &Value{IR.TypeI64, uint64(p.(uint))})
		case reflect.Int64:
			Params = append(Params, &Value{IR.TypeI64, p.(int64)})
		case reflect.Uint64:
			Params = append(Params, &Value{IR.TypeI64, p.(uint64)})

		case reflect.Float32:
			Params = append(Params, &Value{IR.TypeF32, p.(float32)})
		case reflect.Float64:
			Params = append(Params, &Value{IR.TypeF64, p.(float64)})

		default:
			panic("Parameter type not valid,only support (u)int8,u(int16),u(int32),u(int),u(int64),float32,float64")
		}
	}

	err = frame.Init(funcIndex, vm, Params)
	endIndice := vm.FunctionCodes[funcIndex].EndIndice
	if err != nil {
		vm.panic(err)
	}
	//Execute
	for {
		lastPC := frame.PC
		ins := frame.Instruction[frame.PC]

		switch ins.Op.Code {
		case IR.OPCunreachable:
			vm.panic("unreachable executed")
		case IR.OPCbr:
			td := ins.Imm.(*IR.BranchImm).TargetDepth
			tarInsIndex := endIndice[int(td)]
			frame.advanceTo(tarInsIndex)
		case IR.OPCbr_if:
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			con, _ := frame.Stack.Pop()
			v := con.Value().(int32)
			if v == 1 {
				td := ins.Imm.(*IR.BranchImm).TargetDepth
				tarInsIndex := endIndice[int(td)]
				frame.advanceTo(tarInsIndex)
			} else {
				frame.advance(1)
			}

		case IR.OPCbr_table:
		case IR.OPCreturn_:
			var retV IR.InterfaceValue
			vm.CurrentFrame -= 1
			hasResult := !frame.Stack.Empty()
			if hasResult {
				//has return value
				retV, _ = frame.Stack.Pop()
				//TODO check frame.FunctionSig with retV
			}
			if vm.CurrentFrame == -1 {
				if hasResult {
					vm.ReturnValue = retV
				}
				return nil
			} else {
				if hasResult {
					frame.Stack.Push(retV)
				}
				frame = vm.GetCurrentFrame()
			}
		case IR.OPCcall:
			frame.advance(1)
			funcIndex := ins.Imm.(*IR.FunctionImm).FunctionIndex
			vm.CurrentFrame += 1
			if vm.CurrentFrame >= MAXFRAME {
				vm.panic(types.ErrBeyondMaxFrame)
			}

			fType := vm.Module.Types[int(vm.FunctionCodes[funcIndex].Type.Index)]
			paraCount := fType.Params.NumElems
			if frame.Stack.Len() < int(paraCount) {
				vm.panic(types.ErrStackSizeErr)
			}
			params := make([]IR.InterfaceValue, paraCount)
			for elemIndex := range fType.Params.Elems {
				v, _ := frame.Stack.Pop()
				params[int(paraCount)-1-elemIndex] = v
			}

			frame = vm.GetCurrentFrame()
			err = frame.Init(int(funcIndex), vm, params)
			if err != nil {
				vm.panic(err)
			}

		case IR.OPCcall_indirect:
		case IR.OPCdrop:
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			frame.Stack.Pop()
			frame.advance(1)
		case IR.OPCselect:
			if frame.Stack.Len() < 3 {
				vm.panic(types.ErrStackSizeErr)
			}
			a, _ := frame.Stack.Pop()
			b, _ := frame.Stack.Pop()
			c, _ := frame.Stack.Pop()
			if IsZero(c) {
				frame.Stack.Push(b)
			} else {
				frame.Stack.Push(a)
			}
			frame.advance(1)

		case IR.OPCget_local:
			index := ins.Imm.(*IR.GetOrSetVariableImm).VariableIndex
			if index >= uint64(len(frame.Locals)) {
				vm.panic("get_local index out of range")
			}
			frame.Stack.Push(frame.Locals[index])
			frame.advance(1)
		case IR.OPCset_local:
			index := ins.Imm.(*IR.GetOrSetVariableImm).VariableIndex
			if index >= uint64(len(frame.Locals)) {
				vm.panic("set_local index out of range")
			}
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			a, _ := frame.Stack.Pop()
			frame.Locals[index] = a
			frame.advance(1)
		case IR.OPCtee_local:
			index := ins.Imm.(*IR.GetOrSetVariableImm).VariableIndex
			if index >= uint64(len(frame.Locals)) {
				vm.panic("tee_local index out of range")
			}
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			a, _ := frame.Stack.Top()
			frame.Locals[index] = a
			frame.advance(1)
		case IR.OPCget_global:
			index := ins.Imm.(*IR.GetOrSetVariableImm).VariableIndex
			if index >= uint64(len(vm.Global)) {
				vm.panic("get_local index out of range")
			}
			frame.Stack.Push(vm.Global[index])
			frame.advance(1)
		case IR.OPCset_global:
			index := ins.Imm.(*IR.GetOrSetVariableImm).VariableIndex
			if index >= uint64(len(vm.Global)) {
				vm.panic("set_local index out of range")
			}
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			a, _ := frame.Stack.Pop()
			vm.Global[index] = a
			frame.advance(1)
		case IR.OPCtable_get:
		case IR.OPCtable_set:
		case IR.OPCthrow_:
		case IR.OPCrethrow:
		case IR.OPCnop:
			frame.advance(1)
		case IR.OPCi32_load:
			if frame.Stack.Empty() {
				vm.panic(types.ErrStackSizeErr)
			}
			baseVal, _ := frame.Stack.Pop()
			base := baseVal.Value().(int32)
			offset := ins.Imm.(*IR.LoadOrStoreImm).Offset
			addr := base + int32(offset)
			if addr < 0 {
				addr += int32(len(vm.Memory))
			}
			val := binary.LittleEndian.Uint32(vm.Memory[addr : addr+4])
			frame.Stack.Push(&Value{IR.TypeI32, int32(val)})
			frame.advance(1)
		case IR.OPCi64_load:
		case IR.OPCf32_load:
		case IR.OPCf64_load:
		case IR.OPCi32_load8_s:
		case IR.OPCi32_load8_u:
		case IR.OPCi32_load16_s:
		case IR.OPCi32_load16_u:
		case IR.OPCi64_load8_s:
		case IR.OPCi64_load8_u:
		case IR.OPCi64_load16_s:
		case IR.OPCi64_load16_u:
		case IR.OPCi64_load32_s:
		case IR.OPCi64_load32_u:
		case IR.OPCi32_store:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			valVal, _ := frame.Stack.Pop()
			baseVal, _ := frame.Stack.Pop()
			base := baseVal.Value().(int32)
			val := valVal.Value().(int32)
			offset := ins.Imm.(*IR.LoadOrStoreImm).Offset
			addr := base + int32(offset)
			if addr < 0 {
				addr += int32(len(vm.Memory))
			}
			binary.LittleEndian.PutUint32(vm.Memory[addr:addr+4], uint32(val))
			frame.advance(1)
		case IR.OPCi64_store:
		case IR.OPCf32_store:
		case IR.OPCf64_store:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			valVal, _ := frame.Stack.Pop()
			baseVal, _ := frame.Stack.Pop()
			base := baseVal.Value().(int32)
			val := valVal.Value().(float64)
			offset := ins.Imm.(*IR.LoadOrStoreImm).Offset
			addr := base + int32(offset)
			if addr < 0 {
				addr += int32(len(vm.Memory))
			}
			binary.LittleEndian.PutUint64(vm.Memory[addr:addr+8], math.Float64bits((val)))
			frame.advance(1)

		case IR.OPCi32_store8:
		case IR.OPCi32_store16:
		case IR.OPCi64_store8:
		case IR.OPCi64_store16:
		case IR.OPCi64_store32:
		case IR.OPCmemory_size:
		case IR.OPCmemory_grow:
		case IR.OPCi32_const:
			val := ins.Imm.(*IR.LiteralImm_I32).Value
			frame.Stack.Push(&Value{Typ: IR.TypeI32, Val: val})
			frame.advance(1)
		case IR.OPCi64_const:
			val := ins.Imm.(*IR.LiteralImm_I64).Value
			frame.Stack.Push(&Value{IR.TypeI64, val})
			frame.advance(1)
		case IR.OPCf32_const:
			val := ins.Imm.(*IR.LiteralImm_F32).Value
			frame.Stack.Push(&Value{IR.TypeF32, val})
			frame.advance(1)
		case IR.OPCf64_const:
			val := ins.Imm.(*IR.LiteralImm_F64).Value
			frame.Stack.Push(&Value{IR.TypeF64, val})
			frame.advance(1)
		case IR.OPCi32_eqz:
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			a, _ := frame.Stack.Pop()
			if a.Value().(int32) == 0 {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
			frame.advance(1)
		case IR.OPCi32_eq:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			if a.Value().(int32) == b.Value().(int32) {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
			frame.advance(1)
		case IR.OPCi32_ne:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			if a.Value().(int32) != b.Value().(int32) {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
			frame.advance(1)
		case IR.OPCi32_lt_s:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			if a.Value().(int32) < b.Value().(int32) {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
			frame.advance(1)
		case IR.OPCi32_lt_u:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			if a.Value().(uint32) < b.Value().(uint32) {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
			frame.advance(1)
		case IR.OPCi32_gt_s:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			if a.Value().(int32) > b.Value().(int32) {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
			frame.advance(1)
		case IR.OPCi32_gt_u:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			if a.Value().(uint32) > b.Value().(uint32) {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
			frame.advance(1)
		case IR.OPCi32_le_s:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			if a.Value().(int32) <= b.Value().(int32) {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
			frame.advance(1)
		case IR.OPCi32_le_u:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			if a.Value().(uint32) <= b.Value().(uint32) {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
			frame.advance(1)
		case IR.OPCi32_ge_s:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			if a.Value().(int32) >= b.Value().(int32) {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
			frame.advance(1)
		case IR.OPCi32_ge_u:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			if a.Value().(uint32) >= b.Value().(uint32) {
				frame.Stack.Push(&Value{IR.TypeI32, int32(1)})
			} else {
				frame.Stack.Push(&Value{IR.TypeI32, int32(0)})
			}
			frame.advance(1)

		case IR.OPCi64_eqz:
		case IR.OPCi64_eq:
		case IR.OPCi64_ne:
		case IR.OPCi64_lt_s:
		case IR.OPCi64_lt_u:
		case IR.OPCi64_gt_s:
		case IR.OPCi64_gt_u:
		case IR.OPCi64_le_s:
		case IR.OPCi64_le_u:
		case IR.OPCi64_ge_s:
		case IR.OPCi64_ge_u:

		case IR.OPCf32_eq:
		case IR.OPCf32_ne:
		case IR.OPCf32_lt:
		case IR.OPCf32_gt:
		case IR.OPCf32_le:
		case IR.OPCf32_ge:
		case IR.OPCf64_eq:
		case IR.OPCf64_ne:
		case IR.OPCf64_lt:
		case IR.OPCf64_gt:
		case IR.OPCf64_le:
		case IR.OPCf64_ge:

		case IR.OPCi32_clz:
		case IR.OPCi32_ctz:
		case IR.OPCi32_popcnt:
		case IR.OPCi32_add:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			c := a.Value().(int32) + b.Value().(int32)
			frame.Stack.Push(&Value{Typ: IR.TypeI32, Val: c})
			frame.advance(1)
		case IR.OPCi32_sub:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			c := a.Value().(int32) - b.Value().(int32)
			frame.Stack.Push(&Value{Typ: IR.TypeI32, Val: c})
			frame.advance(1)
		case IR.OPCi32_mul:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			c := a.Value().(int32) * b.Value().(int32)
			frame.Stack.Push(&Value{Typ: IR.TypeI32, Val: c})
			frame.advance(1)
		case IR.OPCi32_div_s:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			if b.Value().(int32) == 0 {
				vm.panic(types.ErrDivideByZero)
			}
			a, _ := frame.Stack.Pop()
			c := a.Value().(int32) / b.Value().(int32)
			frame.Stack.Push(&Value{Typ: IR.TypeI32, Val: int32(c)})
			frame.advance(1)
		case IR.OPCi32_div_u:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			if b.Value().(uint32) == 0 {
				vm.panic(types.ErrDivideByZero)
			}
			a, _ := frame.Stack.Pop()
			c := a.Value().(uint32) / b.Value().(uint32)
			frame.Stack.Push(&Value{Typ: IR.TypeI32, Val: uint32(c)})
			frame.advance(1)
		case IR.OPCi32_rem_s:
		case IR.OPCi32_rem_u:
		case IR.OPCi32_and_:
		case IR.OPCi32_or_:
		case IR.OPCi32_xor_:
		case IR.OPCi32_shl:
		case IR.OPCi32_shr_s:
		case IR.OPCi32_shr_u:
		case IR.OPCi32_rotl:
		case IR.OPCi32_rotr:
		case IR.OPCi64_clz:
		case IR.OPCi64_ctz:
		case IR.OPCi64_popcnt:
		case IR.OPCi64_add:
		case IR.OPCi64_sub:
		case IR.OPCi64_mul:
		case IR.OPCi64_div_s:
		case IR.OPCi64_div_u:
		case IR.OPCi64_rem_s:
		case IR.OPCi64_rem_u:
		case IR.OPCi64_and_:
		case IR.OPCi64_or_:
		case IR.OPCi64_xor_:
		case IR.OPCi64_shl:
		case IR.OPCi64_shr_s:
		case IR.OPCi64_shr_u:
		case IR.OPCi64_rotl:
		case IR.OPCi64_rotr:
		case IR.OPCf32_abs:
		case IR.OPCf32_neg:
		case IR.OPCf32_ceil:
		case IR.OPCf32_floor:
		case IR.OPCf32_trunc:
		case IR.OPCf32_nearest:
		case IR.OPCf32_sqrt:
		case IR.OPCf32_add:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			c := a.Value().(float32) + b.Value().(float32)
			frame.Stack.Push(&Value{Typ: IR.TypeF32, Val: c})
			frame.advance(1)
		case IR.OPCf32_sub:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			c := a.Value().(float32) - b.Value().(float32)
			frame.Stack.Push(&Value{Typ: IR.TypeF32, Val: c})
			frame.advance(1)
		case IR.OPCf32_mul:
		case IR.OPCf32_div:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			v := a.Value().(float32) / b.Value().(float32)
			frame.Stack.Push(&Value{IR.TypeF32, v})
			frame.advance(1)
		case IR.OPCf32_min:
		case IR.OPCf32_max:
		case IR.OPCf32_copysign:
		case IR.OPCf64_abs:
		case IR.OPCf64_neg:
		case IR.OPCf64_ceil:
		case IR.OPCf64_floor:
		case IR.OPCf64_trunc:
		case IR.OPCf64_nearest:
		case IR.OPCf64_sqrt:
		case IR.OPCf64_add:
		case IR.OPCf64_sub:
		case IR.OPCf64_mul:
			if frame.Stack.Len() < 2 {
				vm.panic(types.ErrStackSizeErr)
			}
			b, _ := frame.Stack.Pop()
			a, _ := frame.Stack.Pop()
			c := a.Value().(float64) * b.Value().(float64)
			frame.Stack.Push(&Value{Typ: IR.TypeF64, Val: c})
			frame.advance(1)
		case IR.OPCf64_div:
		case IR.OPCf64_min:
		case IR.OPCf64_max:
		case IR.OPCf64_copysign:
		case IR.OPCi32_wrap_i64:
		case IR.OPCi32_trunc_s_f32:
		case IR.OPCi32_trunc_u_f32:
		case IR.OPCi32_trunc_s_f64:
		case IR.OPCi32_trunc_u_f64:
		case IR.OPCi64_extend_s_i32:
		case IR.OPCi64_extend_u_i32:
		case IR.OPCi64_trunc_s_f32:
		case IR.OPCi64_trunc_u_f32:
		case IR.OPCi64_trunc_s_f64:
		case IR.OPCi64_trunc_u_f64:
		case IR.OPCf32_convert_s_i32:
		case IR.OPCf32_convert_u_i32:
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			a, _ := frame.Stack.Pop()
			m, ok := a.Value().(uint32)
			if !ok {
				vm.panic(fmt.Sprintf(types.ErrTypeAssertion, "uint32"))
			}
			v := float32(m)
			frame.Stack.Push(&Value{IR.TypeF32, v})
			frame.advance(1)
		case IR.OPCf32_convert_s_i64:
		case IR.OPCf32_convert_u_i64:
		case IR.OPCf32_demote_f64:
		case IR.OPCf64_convert_s_i32:
		case IR.OPCf64_convert_u_i32:
		case IR.OPCf64_convert_s_i64:
		case IR.OPCf64_convert_u_i64:
		case IR.OPCf64_promote_f32:
		case IR.OPCi32_reinterpret_f32:
		case IR.OPCi64_reinterpret_f64:
		case IR.OPCf32_reinterpret_i32:
		case IR.OPCf64_reinterpret_i64:
		case IR.OPCi32_extend8_s:
		case IR.OPCi32_extend16_s:
		case IR.OPCi64_extend8_s:
		case IR.OPCi64_extend16_s:
		case IR.OPCi64_extend32_s:
		case IR.OPCref_null:
		case IR.OPCref_isnull:
		case IR.OPCref_func:

		case IR.OPCblock:
			frame.advance(1)
		case IR.OPCloop:
			frame.advance(1)
		case IR.OPCif_:
			if frame.Stack.Len() < 1 {
				vm.panic(types.ErrStackSizeErr)
			}
			con, _ := frame.Stack.Pop()
			v := con.Value().(int32)
			if v == 1 {
				frame.advanceTo(1)
			} else {
				//TODO advanctTo else or end
			}
		case IR.OPCelse_:
		case IR.OPCend:
			switch ins.MatchedIndex {
			case types.LastOpcode:
				//if ins.Index == len(frame.Instruction)-1 {
				var retV IR.InterfaceValue
				vm.CurrentFrame -= 1
				hasResult := !frame.Stack.Empty()
				if hasResult {
					//has return value
					retV, _ = frame.Stack.Pop()
					//TODO check frame.FunctionSig with retV
				}
				if vm.CurrentFrame == -1 {
					if hasResult {
						vm.ReturnValue = retV
					}
					return nil
				} else {
					frame = vm.GetCurrentFrame()
					if hasResult {
						//TODO assert frame.Stack is empty
						frame.Stack.Push(retV)
					}
				}
			case -1:
				vm.panic("end ins matched index illegal")
			default:
				//TODO pop value stack since block or if
				mindex := ins.MatchedIndex
				switch frame.Instruction[mindex].Op.Code {
				case IR.OPCblock:
					frame.advance(1)
				case IR.OPCif_:
					frame.advance(1)
				case IR.OPCloop:
					frame.advanceTo(mindex + 1)
				default:
					vm.panic("end ins not point to block/if/loop")
				}
			}

			//TODO :end for 'if','loop','block'
		case IR.OPCtry_:
		case IR.OPCcatch_:
		case IR.OPCcatch_all:
		default:
			frame.runBinaryOp(vm, &ins)
		}
		if lastPC == frame.PC {
			panicStr := fmt.Sprintf("PC not changed. PC: %d instruction: %v", frame.PC, frame.Instruction[frame.PC].Op.Name)
			vm.panic(panicStr)
		}
	}
}
