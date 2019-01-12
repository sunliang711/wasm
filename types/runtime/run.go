package runtime

import (
	"fmt"
	"wasm/types"
	"wasm/types/IR"
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

func (vm *VM) Run(funcName string, funcIndex int, params []IR.InterfaceValue) (err error) {
	//TODO check all type assertion
	defer utils.CatchError(&err)
	frame := vm.GetCurrentFrame()
	if frame == nil {
		vm.panic("Frame stack overflow.")
	}
	// enable funcName,if funcIndex < 0
	if funcIndex < 0 {
		funcIndex, err = vm.Module.GetFuncIndexWithName(funcName)
		if err != nil {
			vm.panic(err)
		}
	}
	err = frame.Init(funcIndex, vm, params)
	if err != nil {
		vm.panic(err)
	}
	//Execute
	for {
		ins := frame.Instruction[frame.PC]

		switch ins.Op.Code {
		case IR.OPCunreachable:
			vm.panic("unreachable executed")
		case IR.OPCbr:
		case IR.OPCbr_if:
		case IR.OPCbr_table:
		case IR.OPCreturn_:
		case IR.OPCcall:
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
		case IR.OPCset_global:
		case IR.OPCtable_get:
		case IR.OPCtable_set:
		case IR.OPCthrow_:
		case IR.OPCrethrow:
		case IR.OPCnop:
			frame.advance(1)
		case IR.OPCi32_load:
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
		case IR.OPCi64_store:
		case IR.OPCf32_store:
		case IR.OPCf64_store:
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
		case IR.OPCi32_trunc_s_sat_f32:
		case IR.OPCi32_trunc_u_sat_f32:
		case IR.OPCi32_trunc_s_sat_f64:
		case IR.OPCi32_trunc_u_sat_f64:
		case IR.OPCi64_trunc_s_sat_f32:
		case IR.OPCi64_trunc_u_sat_f32:
		case IR.OPCi64_trunc_s_sat_f64:
		case IR.OPCi64_trunc_u_sat_f64:
		case IR.OPCmemory_init:
		case IR.OPCmemory_drop:
		case IR.OPCmemory_copy:
		case IR.OPCmemory_fill:
		case IR.OPCtable_init:
		case IR.OPCtable_drop:
		case IR.OPCtable_copy:
		case IR.OPCv128_const:
		case IR.OPCv128_load:
		case IR.OPCv128_store:
		case IR.OPCi8x16_splat:
		case IR.OPCi16x8_splat:
		case IR.OPCi32x4_splat:
		case IR.OPCi64x2_splat:
		case IR.OPCf32x4_splat:
		case IR.OPCf64x2_splat:
		case IR.OPCi8x16_extract_lane_s:
		case IR.OPCi8x16_extract_lane_u:
		case IR.OPCi16x8_extract_lane_s:
		case IR.OPCi16x8_extract_lane_u:
		case IR.OPCi32x4_extract_lane:
		case IR.OPCi64x2_extract_lane:
		case IR.OPCf32x4_extract_lane:
		case IR.OPCf64x2_extract_lane:
		case IR.OPCi8x16_replace_lane:
		case IR.OPCi16x8_replace_lane:
		case IR.OPCi32x4_replace_lane:
		case IR.OPCi64x2_replace_lane:
		case IR.OPCf32x4_replace_lane:
		case IR.OPCf64x2_replace_lane:
		case IR.OPCv8x16_shuffle:
		case IR.OPCi8x16_add:
		case IR.OPCi16x8_add:
		case IR.OPCi32x4_add:
		case IR.OPCi64x2_add:
		case IR.OPCi8x16_sub:
		case IR.OPCi16x8_sub:
		case IR.OPCi32x4_sub:
		case IR.OPCi64x2_sub:
		case IR.OPCi8x16_mul:
		case IR.OPCi16x8_mul:
		case IR.OPCi32x4_mul:
		case IR.OPCi8x16_neg:
		case IR.OPCi16x8_neg:
		case IR.OPCi32x4_neg:
		case IR.OPCi64x2_neg:
		case IR.OPCi8x16_add_saturate_s:
		case IR.OPCi8x16_add_saturate_u:
		case IR.OPCi16x8_add_saturate_s:
		case IR.OPCi16x8_add_saturate_u:
		case IR.OPCi8x16_sub_saturate_s:
		case IR.OPCi8x16_sub_saturate_u:
		case IR.OPCi16x8_sub_saturate_s:
		case IR.OPCi16x8_sub_saturate_u:
		case IR.OPCi8x16_shl:
		case IR.OPCi16x8_shl:
		case IR.OPCi32x4_shl:
		case IR.OPCi64x2_shl:
		case IR.OPCi8x16_shr_s:
		case IR.OPCi8x16_shr_u:
		case IR.OPCi16x8_shr_s:
		case IR.OPCi16x8_shr_u:
		case IR.OPCi32x4_shr_s:
		case IR.OPCi32x4_shr_u:
		case IR.OPCi64x2_shr_s:
		case IR.OPCi64x2_shr_u:
		case IR.OPCv128_and:
		case IR.OPCv128_or:
		case IR.OPCv128_xor:
		case IR.OPCv128_not:
		case IR.OPCv128_bitselect:
		case IR.OPCi8x16_any_true:
		case IR.OPCi16x8_any_true:
		case IR.OPCi32x4_any_true:
		case IR.OPCi64x2_any_true:
		case IR.OPCi8x16_all_true:
		case IR.OPCi16x8_all_true:
		case IR.OPCi32x4_all_true:
		case IR.OPCi64x2_all_true:
		case IR.OPCi8x16_eq:
		case IR.OPCi16x8_eq:
		case IR.OPCi32x4_eq:
		case IR.OPCf32x4_eq:
		case IR.OPCf64x2_eq:
		case IR.OPCi8x16_ne:
		case IR.OPCi16x8_ne:
		case IR.OPCi32x4_ne:
		case IR.OPCf32x4_ne:
		case IR.OPCf64x2_ne:
		case IR.OPCi8x16_lt_s:
		case IR.OPCi8x16_lt_u:
		case IR.OPCi16x8_lt_s:
		case IR.OPCi16x8_lt_u:
		case IR.OPCi32x4_lt_s:
		case IR.OPCi32x4_lt_u:
		case IR.OPCf32x4_lt:
		case IR.OPCf64x2_lt:
		case IR.OPCi8x16_le_s:
		case IR.OPCi8x16_le_u:
		case IR.OPCi16x8_le_s:
		case IR.OPCi16x8_le_u:
		case IR.OPCi32x4_le_s:
		case IR.OPCi32x4_le_u:
		case IR.OPCf32x4_le:
		case IR.OPCf64x2_le:
		case IR.OPCi8x16_gt_s:
		case IR.OPCi8x16_gt_u:
		case IR.OPCi16x8_gt_s:
		case IR.OPCi16x8_gt_u:
		case IR.OPCi32x4_gt_s:
		case IR.OPCi32x4_gt_u:
		case IR.OPCf32x4_gt:
		case IR.OPCf64x2_gt:
		case IR.OPCi8x16_ge_s:
		case IR.OPCi8x16_ge_u:
		case IR.OPCi16x8_ge_s:
		case IR.OPCi16x8_ge_u:
		case IR.OPCi32x4_ge_s:
		case IR.OPCi32x4_ge_u:
		case IR.OPCf32x4_ge:
		case IR.OPCf64x2_ge:
		case IR.OPCf32x4_neg:
		case IR.OPCf64x2_neg:
		case IR.OPCf32x4_abs:
		case IR.OPCf64x2_abs:
		case IR.OPCf32x4_min:
		case IR.OPCf64x2_min:
		case IR.OPCf32x4_max:
		case IR.OPCf64x2_max:
		case IR.OPCf32x4_add:
		case IR.OPCf64x2_add:
		case IR.OPCf32x4_sub:
		case IR.OPCf64x2_sub:
		case IR.OPCf32x4_div:
		case IR.OPCf64x2_div:
		case IR.OPCf32x4_mul:
		case IR.OPCf64x2_mul:
		case IR.OPCf32x4_sqrt:
		case IR.OPCf64x2_sqrt:
		case IR.OPCf32x4_convert_s_i32x4:
		case IR.OPCf32x4_convert_u_i32x4:
		case IR.OPCf64x2_convert_s_i64x2:
		case IR.OPCf64x2_convert_u_i64x2:
		case IR.OPCi32x4_trunc_s_sat_f32x4:
		case IR.OPCi32x4_trunc_u_sat_f32x4:
		case IR.OPCi64x2_trunc_s_sat_f64x2:
		case IR.OPCi64x2_trunc_u_sat_f64x2:
		case IR.OPCatomic_wake:
		case IR.OPCi32_atomic_wait:
		case IR.OPCi64_atomic_wait:
		case IR.OPCi32_atomic_load:
		case IR.OPCi64_atomic_load:
		case IR.OPCi32_atomic_load8_u:
		case IR.OPCi32_atomic_load16_u:
		case IR.OPCi64_atomic_load8_u:
		case IR.OPCi64_atomic_load16_u:
		case IR.OPCi64_atomic_load32_u:
		case IR.OPCi32_atomic_store:
		case IR.OPCi64_atomic_store:
		case IR.OPCi32_atomic_store8:
		case IR.OPCi32_atomic_store16:
		case IR.OPCi64_atomic_store8:
		case IR.OPCi64_atomic_store16:
		case IR.OPCi64_atomic_store32:
		case IR.OPCi32_atomic_rmw_add:
		case IR.OPCi64_atomic_rmw_add:
		case IR.OPCi32_atomic_rmw8_u_add:
		case IR.OPCi32_atomic_rmw16_u_add:
		case IR.OPCi64_atomic_rmw8_u_add:
		case IR.OPCi64_atomic_rmw16_u_add:
		case IR.OPCi64_atomic_rmw32_u_add:
		case IR.OPCi32_atomic_rmw_sub:
		case IR.OPCi64_atomic_rmw_sub:
		case IR.OPCi32_atomic_rmw8_u_sub:
		case IR.OPCi32_atomic_rmw16_u_sub:
		case IR.OPCi64_atomic_rmw8_u_sub:
		case IR.OPCi64_atomic_rmw16_u_sub:
		case IR.OPCi64_atomic_rmw32_u_sub:
		case IR.OPCi32_atomic_rmw_and:
		case IR.OPCi64_atomic_rmw_and:
		case IR.OPCi32_atomic_rmw8_u_and:
		case IR.OPCi32_atomic_rmw16_u_and:
		case IR.OPCi64_atomic_rmw8_u_and:
		case IR.OPCi64_atomic_rmw16_u_and:
		case IR.OPCi64_atomic_rmw32_u_and:
		case IR.OPCi32_atomic_rmw_or:
		case IR.OPCi64_atomic_rmw_or:
		case IR.OPCi32_atomic_rmw8_u_or:
		case IR.OPCi32_atomic_rmw16_u_or:
		case IR.OPCi64_atomic_rmw8_u_or:
		case IR.OPCi64_atomic_rmw16_u_or:
		case IR.OPCi64_atomic_rmw32_u_or:
		case IR.OPCi32_atomic_rmw_xor:
		case IR.OPCi64_atomic_rmw_xor:
		case IR.OPCi32_atomic_rmw8_u_xor:
		case IR.OPCi32_atomic_rmw16_u_xor:
		case IR.OPCi64_atomic_rmw8_u_xor:
		case IR.OPCi64_atomic_rmw16_u_xor:
		case IR.OPCi64_atomic_rmw32_u_xor:
		case IR.OPCi32_atomic_rmw_xchg:
		case IR.OPCi64_atomic_rmw_xchg:
		case IR.OPCi32_atomic_rmw8_u_xchg:
		case IR.OPCi32_atomic_rmw16_u_xchg:
		case IR.OPCi64_atomic_rmw8_u_xchg:
		case IR.OPCi64_atomic_rmw16_u_xchg:
		case IR.OPCi64_atomic_rmw32_u_xchg:
		case IR.OPCi32_atomic_rmw_cmpxchg:
		case IR.OPCi64_atomic_rmw_cmpxchg:
		case IR.OPCi32_atomic_rmw8_u_cmpxchg:
		case IR.OPCi32_atomic_rmw16_u_cmpxchg:
		case IR.OPCi64_atomic_rmw8_u_cmpxchg:
		case IR.OPCi64_atomic_rmw16_u_cmpxchg:
		case IR.OPCi64_atomic_rmw32_u_cmpxchg:
		case IR.OPCblock:
		case IR.OPCloop:
		case IR.OPCif_:
		case IR.OPCelse_:
		case IR.OPCend:
			//the last opcode
			if ins.Index == len(frame.Instruction)-1 {
				if !frame.Stack.Empty() {
					//has return value
					retV, _ := frame.Stack.Pop()
					//TODO check frame.FunctionSig with retV
					vm.CurrentFrame -= 1
					//empty frame stack
					if vm.CurrentFrame == -1 {
						vm.ReturnValue = retV
						return nil
					} else {
						frame = vm.GetCurrentFrame()
					}
				}
			}
			//TODO :end for 'if','loop','block'
		case IR.OPCtry_:
		case IR.OPCcatch_:
		case IR.OPCcatch_all:
		}
	}
}
