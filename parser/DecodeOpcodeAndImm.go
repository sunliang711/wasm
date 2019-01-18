// Note this file is created by makeType.sh
// Don't modify it directly
package parser

import (
	"bytes"
	"fmt"
	"io"
	"wasm/core/IR"
)

func DecodeOpcodeAndImm(opcodeBytes []byte, funcDef *IR.FunctionDef) ([]IR.Instruction, error) {
	rd := bytes.NewReader(opcodeBytes)
	var (
		ins       []IR.Instruction
		codeIndex int
	)

	for {
		opc, err := DecodeOpcode(rd)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch IR.Opcode(opc) {
		case IR.OPCunreachable:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCunreachable], &imm, codeIndex, -1})
		case IR.OPCbr:
			imm := IR.BranchImm{}
			err = DecodeBranchImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCbr], &imm, codeIndex, -1})
		case IR.OPCbr_if:
			imm := IR.BranchImm{}
			err = DecodeBranchImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCbr_if], &imm, codeIndex, -1})
		case IR.OPCbr_table:
			imm := IR.BranchTableImm{}
			err = DecodeBranchTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCbr_table], &imm, codeIndex, -1})
		case IR.OPCreturn_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCreturn_], &imm, codeIndex, -1})
		case IR.OPCcall:
			imm := IR.FunctionImm{}
			err = DecodeFunctionImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCcall], &imm, codeIndex, -1})
		case IR.OPCcall_indirect:
			imm := IR.CallIndirectImm{}
			err = DecodeCallIndirectImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCcall_indirect], &imm, codeIndex, -1})
		case IR.OPCdrop:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCdrop], &imm, codeIndex, -1})
		case IR.OPCselect:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCselect], &imm, codeIndex, -1})
		case IR.OPCget_local:
			imm := IR.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCget_local], &imm, codeIndex, -1})
		case IR.OPCset_local:
			imm := IR.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCset_local], &imm, codeIndex, -1})
		case IR.OPCtee_local:
			imm := IR.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCtee_local], &imm, codeIndex, -1})
		case IR.OPCget_global:
			imm := IR.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCget_global], &imm, codeIndex, -1})
		case IR.OPCset_global:
			imm := IR.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCset_global], &imm, codeIndex, -1})
		case IR.OPCtable_get:
			imm := IR.TableImm{}
			err = DecodeTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCtable_get], &imm, codeIndex, -1})
		case IR.OPCtable_set:
			imm := IR.TableImm{}
			err = DecodeTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCtable_set], &imm, codeIndex, -1})
		case IR.OPCthrow_:
			imm := IR.ExceptionTypeImm{}
			err = DecodeExceptionTypeImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCthrow_], &imm, codeIndex, -1})
		case IR.OPCrethrow:
			imm := IR.RethrowImm{}
			err = DecodeRethrowImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCrethrow], &imm, codeIndex, -1})
		case IR.OPCnop:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCnop], &imm, codeIndex, -1})
		case IR.OPCi32_load:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_load], &imm, codeIndex, -1})
		case IR.OPCi64_load:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_load], &imm, codeIndex, -1})
		case IR.OPCf32_load:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_load], &imm, codeIndex, -1})
		case IR.OPCf64_load:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_load], &imm, codeIndex, -1})
		case IR.OPCi32_load8_s:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_load8_s], &imm, codeIndex, -1})
		case IR.OPCi32_load8_u:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_load8_u], &imm, codeIndex, -1})
		case IR.OPCi32_load16_s:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_load16_s], &imm, codeIndex, -1})
		case IR.OPCi32_load16_u:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_load16_u], &imm, codeIndex, -1})
		case IR.OPCi64_load8_s:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_load8_s], &imm, codeIndex, -1})
		case IR.OPCi64_load8_u:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_load8_u], &imm, codeIndex, -1})
		case IR.OPCi64_load16_s:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_load16_s], &imm, codeIndex, -1})
		case IR.OPCi64_load16_u:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_load16_u], &imm, codeIndex, -1})
		case IR.OPCi64_load32_s:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_load32_s], &imm, codeIndex, -1})
		case IR.OPCi64_load32_u:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_load32_u], &imm, codeIndex, -1})
		case IR.OPCi32_store:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_store], &imm, codeIndex, -1})
		case IR.OPCi64_store:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_store], &imm, codeIndex, -1})
		case IR.OPCf32_store:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_store], &imm, codeIndex, -1})
		case IR.OPCf64_store:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_store], &imm, codeIndex, -1})
		case IR.OPCi32_store8:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_store8], &imm, codeIndex, -1})
		case IR.OPCi32_store16:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_store16], &imm, codeIndex, -1})
		case IR.OPCi64_store8:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_store8], &imm, codeIndex, -1})
		case IR.OPCi64_store16:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_store16], &imm, codeIndex, -1})
		case IR.OPCi64_store32:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_store32], &imm, codeIndex, -1})
		case IR.OPCmemory_size:
			imm := IR.MemoryImm{}
			err = DecodeMemoryImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCmemory_size], &imm, codeIndex, -1})
		case IR.OPCmemory_grow:
			imm := IR.MemoryImm{}
			err = DecodeMemoryImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCmemory_grow], &imm, codeIndex, -1})
		case IR.OPCi32_const:
			imm := IR.LiteralImm_I32{}
			err = DecodeLiteralImm_I32(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_const], &imm, codeIndex, -1})
		case IR.OPCi64_const:
			imm := IR.LiteralImm_I64{}
			err = DecodeLiteralImm_I64(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_const], &imm, codeIndex, -1})
		case IR.OPCf32_const:
			imm := IR.LiteralImm_F32{}
			err = DecodeLiteralImm_F32(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_const], &imm, codeIndex, -1})
		case IR.OPCf64_const:
			imm := IR.LiteralImm_F64{}
			err = DecodeLiteralImm_F64(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_const], &imm, codeIndex, -1})
		case IR.OPCi32_eqz:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_eqz], &imm, codeIndex, -1})
		case IR.OPCi32_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_eq], &imm, codeIndex, -1})
		case IR.OPCi32_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_ne], &imm, codeIndex, -1})
		case IR.OPCi32_lt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_lt_s], &imm, codeIndex, -1})
		case IR.OPCi32_lt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_lt_u], &imm, codeIndex, -1})
		case IR.OPCi32_gt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_gt_s], &imm, codeIndex, -1})
		case IR.OPCi32_gt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_gt_u], &imm, codeIndex, -1})
		case IR.OPCi32_le_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_le_s], &imm, codeIndex, -1})
		case IR.OPCi32_le_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_le_u], &imm, codeIndex, -1})
		case IR.OPCi32_ge_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_ge_s], &imm, codeIndex, -1})
		case IR.OPCi32_ge_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_ge_u], &imm, codeIndex, -1})
		case IR.OPCi64_eqz:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_eqz], &imm, codeIndex, -1})
		case IR.OPCi64_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_eq], &imm, codeIndex, -1})
		case IR.OPCi64_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_ne], &imm, codeIndex, -1})
		case IR.OPCi64_lt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_lt_s], &imm, codeIndex, -1})
		case IR.OPCi64_lt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_lt_u], &imm, codeIndex, -1})
		case IR.OPCi64_gt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_gt_s], &imm, codeIndex, -1})
		case IR.OPCi64_gt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_gt_u], &imm, codeIndex, -1})
		case IR.OPCi64_le_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_le_s], &imm, codeIndex, -1})
		case IR.OPCi64_le_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_le_u], &imm, codeIndex, -1})
		case IR.OPCi64_ge_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_ge_s], &imm, codeIndex, -1})
		case IR.OPCi64_ge_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_ge_u], &imm, codeIndex, -1})
		case IR.OPCf32_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_eq], &imm, codeIndex, -1})
		case IR.OPCf32_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_ne], &imm, codeIndex, -1})
		case IR.OPCf32_lt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_lt], &imm, codeIndex, -1})
		case IR.OPCf32_gt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_gt], &imm, codeIndex, -1})
		case IR.OPCf32_le:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_le], &imm, codeIndex, -1})
		case IR.OPCf32_ge:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_ge], &imm, codeIndex, -1})
		case IR.OPCf64_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_eq], &imm, codeIndex, -1})
		case IR.OPCf64_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_ne], &imm, codeIndex, -1})
		case IR.OPCf64_lt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_lt], &imm, codeIndex, -1})
		case IR.OPCf64_gt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_gt], &imm, codeIndex, -1})
		case IR.OPCf64_le:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_le], &imm, codeIndex, -1})
		case IR.OPCf64_ge:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_ge], &imm, codeIndex, -1})
		case IR.OPCi32_clz:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_clz], &imm, codeIndex, -1})
		case IR.OPCi32_ctz:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_ctz], &imm, codeIndex, -1})
		case IR.OPCi32_popcnt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_popcnt], &imm, codeIndex, -1})
		case IR.OPCi32_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_add], &imm, codeIndex, -1})
		case IR.OPCi32_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_sub], &imm, codeIndex, -1})
		case IR.OPCi32_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_mul], &imm, codeIndex, -1})
		case IR.OPCi32_div_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_div_s], &imm, codeIndex, -1})
		case IR.OPCi32_div_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_div_u], &imm, codeIndex, -1})
		case IR.OPCi32_rem_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_rem_s], &imm, codeIndex, -1})
		case IR.OPCi32_rem_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_rem_u], &imm, codeIndex, -1})
		case IR.OPCi32_and_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_and_], &imm, codeIndex, -1})
		case IR.OPCi32_or_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_or_], &imm, codeIndex, -1})
		case IR.OPCi32_xor_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_xor_], &imm, codeIndex, -1})
		case IR.OPCi32_shl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_shl], &imm, codeIndex, -1})
		case IR.OPCi32_shr_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_shr_s], &imm, codeIndex, -1})
		case IR.OPCi32_shr_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_shr_u], &imm, codeIndex, -1})
		case IR.OPCi32_rotl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_rotl], &imm, codeIndex, -1})
		case IR.OPCi32_rotr:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_rotr], &imm, codeIndex, -1})
		case IR.OPCi64_clz:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_clz], &imm, codeIndex, -1})
		case IR.OPCi64_ctz:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_ctz], &imm, codeIndex, -1})
		case IR.OPCi64_popcnt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_popcnt], &imm, codeIndex, -1})
		case IR.OPCi64_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_add], &imm, codeIndex, -1})
		case IR.OPCi64_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_sub], &imm, codeIndex, -1})
		case IR.OPCi64_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_mul], &imm, codeIndex, -1})
		case IR.OPCi64_div_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_div_s], &imm, codeIndex, -1})
		case IR.OPCi64_div_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_div_u], &imm, codeIndex, -1})
		case IR.OPCi64_rem_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_rem_s], &imm, codeIndex, -1})
		case IR.OPCi64_rem_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_rem_u], &imm, codeIndex, -1})
		case IR.OPCi64_and_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_and_], &imm, codeIndex, -1})
		case IR.OPCi64_or_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_or_], &imm, codeIndex, -1})
		case IR.OPCi64_xor_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_xor_], &imm, codeIndex, -1})
		case IR.OPCi64_shl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_shl], &imm, codeIndex, -1})
		case IR.OPCi64_shr_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_shr_s], &imm, codeIndex, -1})
		case IR.OPCi64_shr_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_shr_u], &imm, codeIndex, -1})
		case IR.OPCi64_rotl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_rotl], &imm, codeIndex, -1})
		case IR.OPCi64_rotr:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_rotr], &imm, codeIndex, -1})
		case IR.OPCf32_abs:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_abs], &imm, codeIndex, -1})
		case IR.OPCf32_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_neg], &imm, codeIndex, -1})
		case IR.OPCf32_ceil:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_ceil], &imm, codeIndex, -1})
		case IR.OPCf32_floor:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_floor], &imm, codeIndex, -1})
		case IR.OPCf32_trunc:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_trunc], &imm, codeIndex, -1})
		case IR.OPCf32_nearest:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_nearest], &imm, codeIndex, -1})
		case IR.OPCf32_sqrt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_sqrt], &imm, codeIndex, -1})
		case IR.OPCf32_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_add], &imm, codeIndex, -1})
		case IR.OPCf32_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_sub], &imm, codeIndex, -1})
		case IR.OPCf32_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_mul], &imm, codeIndex, -1})
		case IR.OPCf32_div:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_div], &imm, codeIndex, -1})
		case IR.OPCf32_min:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_min], &imm, codeIndex, -1})
		case IR.OPCf32_max:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_max], &imm, codeIndex, -1})
		case IR.OPCf32_copysign:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_copysign], &imm, codeIndex, -1})
		case IR.OPCf64_abs:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_abs], &imm, codeIndex, -1})
		case IR.OPCf64_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_neg], &imm, codeIndex, -1})
		case IR.OPCf64_ceil:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_ceil], &imm, codeIndex, -1})
		case IR.OPCf64_floor:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_floor], &imm, codeIndex, -1})
		case IR.OPCf64_trunc:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_trunc], &imm, codeIndex, -1})
		case IR.OPCf64_nearest:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_nearest], &imm, codeIndex, -1})
		case IR.OPCf64_sqrt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_sqrt], &imm, codeIndex, -1})
		case IR.OPCf64_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_add], &imm, codeIndex, -1})
		case IR.OPCf64_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_sub], &imm, codeIndex, -1})
		case IR.OPCf64_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_mul], &imm, codeIndex, -1})
		case IR.OPCf64_div:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_div], &imm, codeIndex, -1})
		case IR.OPCf64_min:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_min], &imm, codeIndex, -1})
		case IR.OPCf64_max:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_max], &imm, codeIndex, -1})
		case IR.OPCf64_copysign:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_copysign], &imm, codeIndex, -1})
		case IR.OPCi32_wrap_i64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_wrap_i64], &imm, codeIndex, -1})
		case IR.OPCi32_trunc_s_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_trunc_s_f32], &imm, codeIndex, -1})
		case IR.OPCi32_trunc_u_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_trunc_u_f32], &imm, codeIndex, -1})
		case IR.OPCi32_trunc_s_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_trunc_s_f64], &imm, codeIndex, -1})
		case IR.OPCi32_trunc_u_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_trunc_u_f64], &imm, codeIndex, -1})
		case IR.OPCi64_extend_s_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_extend_s_i32], &imm, codeIndex, -1})
		case IR.OPCi64_extend_u_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_extend_u_i32], &imm, codeIndex, -1})
		case IR.OPCi64_trunc_s_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_trunc_s_f32], &imm, codeIndex, -1})
		case IR.OPCi64_trunc_u_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_trunc_u_f32], &imm, codeIndex, -1})
		case IR.OPCi64_trunc_s_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_trunc_s_f64], &imm, codeIndex, -1})
		case IR.OPCi64_trunc_u_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_trunc_u_f64], &imm, codeIndex, -1})
		case IR.OPCf32_convert_s_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_convert_s_i32], &imm, codeIndex, -1})
		case IR.OPCf32_convert_u_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_convert_u_i32], &imm, codeIndex, -1})
		case IR.OPCf32_convert_s_i64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_convert_s_i64], &imm, codeIndex, -1})
		case IR.OPCf32_convert_u_i64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_convert_u_i64], &imm, codeIndex, -1})
		case IR.OPCf32_demote_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_demote_f64], &imm, codeIndex, -1})
		case IR.OPCf64_convert_s_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_convert_s_i32], &imm, codeIndex, -1})
		case IR.OPCf64_convert_u_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_convert_u_i32], &imm, codeIndex, -1})
		case IR.OPCf64_convert_s_i64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_convert_s_i64], &imm, codeIndex, -1})
		case IR.OPCf64_convert_u_i64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_convert_u_i64], &imm, codeIndex, -1})
		case IR.OPCf64_promote_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_promote_f32], &imm, codeIndex, -1})
		case IR.OPCi32_reinterpret_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_reinterpret_f32], &imm, codeIndex, -1})
		case IR.OPCi64_reinterpret_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_reinterpret_f64], &imm, codeIndex, -1})
		case IR.OPCf32_reinterpret_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32_reinterpret_i32], &imm, codeIndex, -1})
		case IR.OPCf64_reinterpret_i64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64_reinterpret_i64], &imm, codeIndex, -1})
		case IR.OPCi32_extend8_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_extend8_s], &imm, codeIndex, -1})
		case IR.OPCi32_extend16_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_extend16_s], &imm, codeIndex, -1})
		case IR.OPCi64_extend8_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_extend8_s], &imm, codeIndex, -1})
		case IR.OPCi64_extend16_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_extend16_s], &imm, codeIndex, -1})
		case IR.OPCi64_extend32_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_extend32_s], &imm, codeIndex, -1})
		case IR.OPCref_null:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCref_null], &imm, codeIndex, -1})
		case IR.OPCref_isnull:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCref_isnull], &imm, codeIndex, -1})
		case IR.OPCref_func:
			imm := IR.FunctionImm{}
			err = DecodeFunctionImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCref_func], &imm, codeIndex, -1})
		case IR.OPCi32_trunc_s_sat_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_trunc_s_sat_f32], &imm, codeIndex, -1})
		case IR.OPCi32_trunc_u_sat_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_trunc_u_sat_f32], &imm, codeIndex, -1})
		case IR.OPCi32_trunc_s_sat_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_trunc_s_sat_f64], &imm, codeIndex, -1})
		case IR.OPCi32_trunc_u_sat_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_trunc_u_sat_f64], &imm, codeIndex, -1})
		case IR.OPCi64_trunc_s_sat_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_trunc_s_sat_f32], &imm, codeIndex, -1})
		case IR.OPCi64_trunc_u_sat_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_trunc_u_sat_f32], &imm, codeIndex, -1})
		case IR.OPCi64_trunc_s_sat_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_trunc_s_sat_f64], &imm, codeIndex, -1})
		case IR.OPCi64_trunc_u_sat_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_trunc_u_sat_f64], &imm, codeIndex, -1})
		case IR.OPCmemory_init:
			imm := IR.DataSegmentAndMemImm{}
			err = DecodeDataSegmentAndMemImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCmemory_init], &imm, codeIndex, -1})
		case IR.OPCmemory_drop:
			imm := IR.DataSegmentImm{}
			err = DecodeDataSegmentImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCmemory_drop], &imm, codeIndex, -1})
		case IR.OPCmemory_copy:
			imm := IR.MemoryImm{}
			err = DecodeMemoryImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCmemory_copy], &imm, codeIndex, -1})
		case IR.OPCmemory_fill:
			imm := IR.MemoryImm{}
			err = DecodeMemoryImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCmemory_fill], &imm, codeIndex, -1})
		case IR.OPCtable_init:
			imm := IR.ElemSegmentAndTableImm{}
			err = DecodeElemSegmentAndTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCtable_init], &imm, codeIndex, -1})
		case IR.OPCtable_drop:
			imm := IR.ElemSegmentImm{}
			err = DecodeElemSegmentImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCtable_drop], &imm, codeIndex, -1})
		case IR.OPCtable_copy:
			imm := IR.TableImm{}
			err = DecodeTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCtable_copy], &imm, codeIndex, -1})
		case IR.OPCv128_const:
			imm := IR.LiteralImm_V128{}
			err = DecodeLiteralImm_V128(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCv128_const], &imm, codeIndex, -1})
		case IR.OPCv128_load:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCv128_load], &imm, codeIndex, -1})
		case IR.OPCv128_store:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCv128_store], &imm, codeIndex, -1})
		case IR.OPCi8x16_splat:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_splat], &imm, codeIndex, -1})
		case IR.OPCi16x8_splat:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_splat], &imm, codeIndex, -1})
		case IR.OPCi32x4_splat:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_splat], &imm, codeIndex, -1})
		case IR.OPCi64x2_splat:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64x2_splat], &imm, codeIndex, -1})
		case IR.OPCf32x4_splat:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_splat], &imm, codeIndex, -1})
		case IR.OPCf64x2_splat:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_splat], &imm, codeIndex, -1})
		case IR.OPCi8x16_extract_lane_s:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_extract_lane_s], &imm, codeIndex, -1})
		case IR.OPCi8x16_extract_lane_u:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_extract_lane_u], &imm, codeIndex, -1})
		case IR.OPCi16x8_extract_lane_s:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_extract_lane_s], &imm, codeIndex, -1})
		case IR.OPCi16x8_extract_lane_u:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_extract_lane_u], &imm, codeIndex, -1})
		case IR.OPCi32x4_extract_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_extract_lane], &imm, codeIndex, -1})
		case IR.OPCi64x2_extract_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64x2_extract_lane], &imm, codeIndex, -1})
		case IR.OPCf32x4_extract_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_extract_lane], &imm, codeIndex, -1})
		case IR.OPCf64x2_extract_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_extract_lane], &imm, codeIndex, -1})
		case IR.OPCi8x16_replace_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_replace_lane], &imm, codeIndex, -1})
		case IR.OPCi16x8_replace_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_replace_lane], &imm, codeIndex, -1})
		case IR.OPCi32x4_replace_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_replace_lane], &imm, codeIndex, -1})
		case IR.OPCi64x2_replace_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64x2_replace_lane], &imm, codeIndex, -1})
		case IR.OPCf32x4_replace_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_replace_lane], &imm, codeIndex, -1})
		case IR.OPCf64x2_replace_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_replace_lane], &imm, codeIndex, -1})
		case IR.OPCv8x16_shuffle:
			imm := IR.ShuffleImm_16{}
			err = DecodeShuffleImm_16(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCv8x16_shuffle], &imm, codeIndex, -1})
		case IR.OPCi8x16_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_add], &imm, codeIndex, -1})
		case IR.OPCi16x8_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_add], &imm, codeIndex, -1})
		case IR.OPCi32x4_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_add], &imm, codeIndex, -1})
		case IR.OPCi64x2_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64x2_add], &imm, codeIndex, -1})
		case IR.OPCi8x16_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_sub], &imm, codeIndex, -1})
		case IR.OPCi16x8_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_sub], &imm, codeIndex, -1})
		case IR.OPCi32x4_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_sub], &imm, codeIndex, -1})
		case IR.OPCi64x2_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64x2_sub], &imm, codeIndex, -1})
		case IR.OPCi8x16_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_mul], &imm, codeIndex, -1})
		case IR.OPCi16x8_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_mul], &imm, codeIndex, -1})
		case IR.OPCi32x4_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_mul], &imm, codeIndex, -1})
		case IR.OPCi8x16_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_neg], &imm, codeIndex, -1})
		case IR.OPCi16x8_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_neg], &imm, codeIndex, -1})
		case IR.OPCi32x4_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_neg], &imm, codeIndex, -1})
		case IR.OPCi64x2_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64x2_neg], &imm, codeIndex, -1})
		case IR.OPCi8x16_add_saturate_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_add_saturate_s], &imm, codeIndex, -1})
		case IR.OPCi8x16_add_saturate_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_add_saturate_u], &imm, codeIndex, -1})
		case IR.OPCi16x8_add_saturate_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_add_saturate_s], &imm, codeIndex, -1})
		case IR.OPCi16x8_add_saturate_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_add_saturate_u], &imm, codeIndex, -1})
		case IR.OPCi8x16_sub_saturate_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_sub_saturate_s], &imm, codeIndex, -1})
		case IR.OPCi8x16_sub_saturate_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_sub_saturate_u], &imm, codeIndex, -1})
		case IR.OPCi16x8_sub_saturate_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_sub_saturate_s], &imm, codeIndex, -1})
		case IR.OPCi16x8_sub_saturate_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_sub_saturate_u], &imm, codeIndex, -1})
		case IR.OPCi8x16_shl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_shl], &imm, codeIndex, -1})
		case IR.OPCi16x8_shl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_shl], &imm, codeIndex, -1})
		case IR.OPCi32x4_shl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_shl], &imm, codeIndex, -1})
		case IR.OPCi64x2_shl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64x2_shl], &imm, codeIndex, -1})
		case IR.OPCi8x16_shr_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_shr_s], &imm, codeIndex, -1})
		case IR.OPCi8x16_shr_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_shr_u], &imm, codeIndex, -1})
		case IR.OPCi16x8_shr_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_shr_s], &imm, codeIndex, -1})
		case IR.OPCi16x8_shr_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_shr_u], &imm, codeIndex, -1})
		case IR.OPCi32x4_shr_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_shr_s], &imm, codeIndex, -1})
		case IR.OPCi32x4_shr_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_shr_u], &imm, codeIndex, -1})
		case IR.OPCi64x2_shr_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64x2_shr_s], &imm, codeIndex, -1})
		case IR.OPCi64x2_shr_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64x2_shr_u], &imm, codeIndex, -1})
		case IR.OPCv128_and:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCv128_and], &imm, codeIndex, -1})
		case IR.OPCv128_or:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCv128_or], &imm, codeIndex, -1})
		case IR.OPCv128_xor:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCv128_xor], &imm, codeIndex, -1})
		case IR.OPCv128_not:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCv128_not], &imm, codeIndex, -1})
		case IR.OPCv128_bitselect:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCv128_bitselect], &imm, codeIndex, -1})
		case IR.OPCi8x16_any_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_any_true], &imm, codeIndex, -1})
		case IR.OPCi16x8_any_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_any_true], &imm, codeIndex, -1})
		case IR.OPCi32x4_any_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_any_true], &imm, codeIndex, -1})
		case IR.OPCi64x2_any_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64x2_any_true], &imm, codeIndex, -1})
		case IR.OPCi8x16_all_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_all_true], &imm, codeIndex, -1})
		case IR.OPCi16x8_all_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_all_true], &imm, codeIndex, -1})
		case IR.OPCi32x4_all_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_all_true], &imm, codeIndex, -1})
		case IR.OPCi64x2_all_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64x2_all_true], &imm, codeIndex, -1})
		case IR.OPCi8x16_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_eq], &imm, codeIndex, -1})
		case IR.OPCi16x8_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_eq], &imm, codeIndex, -1})
		case IR.OPCi32x4_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_eq], &imm, codeIndex, -1})
		case IR.OPCf32x4_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_eq], &imm, codeIndex, -1})
		case IR.OPCf64x2_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_eq], &imm, codeIndex, -1})
		case IR.OPCi8x16_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_ne], &imm, codeIndex, -1})
		case IR.OPCi16x8_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_ne], &imm, codeIndex, -1})
		case IR.OPCi32x4_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_ne], &imm, codeIndex, -1})
		case IR.OPCf32x4_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_ne], &imm, codeIndex, -1})
		case IR.OPCf64x2_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_ne], &imm, codeIndex, -1})
		case IR.OPCi8x16_lt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_lt_s], &imm, codeIndex, -1})
		case IR.OPCi8x16_lt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_lt_u], &imm, codeIndex, -1})
		case IR.OPCi16x8_lt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_lt_s], &imm, codeIndex, -1})
		case IR.OPCi16x8_lt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_lt_u], &imm, codeIndex, -1})
		case IR.OPCi32x4_lt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_lt_s], &imm, codeIndex, -1})
		case IR.OPCi32x4_lt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_lt_u], &imm, codeIndex, -1})
		case IR.OPCf32x4_lt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_lt], &imm, codeIndex, -1})
		case IR.OPCf64x2_lt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_lt], &imm, codeIndex, -1})
		case IR.OPCi8x16_le_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_le_s], &imm, codeIndex, -1})
		case IR.OPCi8x16_le_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_le_u], &imm, codeIndex, -1})
		case IR.OPCi16x8_le_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_le_s], &imm, codeIndex, -1})
		case IR.OPCi16x8_le_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_le_u], &imm, codeIndex, -1})
		case IR.OPCi32x4_le_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_le_s], &imm, codeIndex, -1})
		case IR.OPCi32x4_le_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_le_u], &imm, codeIndex, -1})
		case IR.OPCf32x4_le:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_le], &imm, codeIndex, -1})
		case IR.OPCf64x2_le:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_le], &imm, codeIndex, -1})
		case IR.OPCi8x16_gt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_gt_s], &imm, codeIndex, -1})
		case IR.OPCi8x16_gt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_gt_u], &imm, codeIndex, -1})
		case IR.OPCi16x8_gt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_gt_s], &imm, codeIndex, -1})
		case IR.OPCi16x8_gt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_gt_u], &imm, codeIndex, -1})
		case IR.OPCi32x4_gt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_gt_s], &imm, codeIndex, -1})
		case IR.OPCi32x4_gt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_gt_u], &imm, codeIndex, -1})
		case IR.OPCf32x4_gt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_gt], &imm, codeIndex, -1})
		case IR.OPCf64x2_gt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_gt], &imm, codeIndex, -1})
		case IR.OPCi8x16_ge_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_ge_s], &imm, codeIndex, -1})
		case IR.OPCi8x16_ge_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi8x16_ge_u], &imm, codeIndex, -1})
		case IR.OPCi16x8_ge_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_ge_s], &imm, codeIndex, -1})
		case IR.OPCi16x8_ge_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi16x8_ge_u], &imm, codeIndex, -1})
		case IR.OPCi32x4_ge_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_ge_s], &imm, codeIndex, -1})
		case IR.OPCi32x4_ge_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_ge_u], &imm, codeIndex, -1})
		case IR.OPCf32x4_ge:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_ge], &imm, codeIndex, -1})
		case IR.OPCf64x2_ge:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_ge], &imm, codeIndex, -1})
		case IR.OPCf32x4_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_neg], &imm, codeIndex, -1})
		case IR.OPCf64x2_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_neg], &imm, codeIndex, -1})
		case IR.OPCf32x4_abs:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_abs], &imm, codeIndex, -1})
		case IR.OPCf64x2_abs:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_abs], &imm, codeIndex, -1})
		case IR.OPCf32x4_min:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_min], &imm, codeIndex, -1})
		case IR.OPCf64x2_min:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_min], &imm, codeIndex, -1})
		case IR.OPCf32x4_max:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_max], &imm, codeIndex, -1})
		case IR.OPCf64x2_max:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_max], &imm, codeIndex, -1})
		case IR.OPCf32x4_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_add], &imm, codeIndex, -1})
		case IR.OPCf64x2_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_add], &imm, codeIndex, -1})
		case IR.OPCf32x4_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_sub], &imm, codeIndex, -1})
		case IR.OPCf64x2_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_sub], &imm, codeIndex, -1})
		case IR.OPCf32x4_div:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_div], &imm, codeIndex, -1})
		case IR.OPCf64x2_div:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_div], &imm, codeIndex, -1})
		case IR.OPCf32x4_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_mul], &imm, codeIndex, -1})
		case IR.OPCf64x2_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_mul], &imm, codeIndex, -1})
		case IR.OPCf32x4_sqrt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_sqrt], &imm, codeIndex, -1})
		case IR.OPCf64x2_sqrt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_sqrt], &imm, codeIndex, -1})
		case IR.OPCf32x4_convert_s_i32x4:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_convert_s_i32x4], &imm, codeIndex, -1})
		case IR.OPCf32x4_convert_u_i32x4:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf32x4_convert_u_i32x4], &imm, codeIndex, -1})
		case IR.OPCf64x2_convert_s_i64x2:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_convert_s_i64x2], &imm, codeIndex, -1})
		case IR.OPCf64x2_convert_u_i64x2:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCf64x2_convert_u_i64x2], &imm, codeIndex, -1})
		case IR.OPCi32x4_trunc_s_sat_f32x4:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_trunc_s_sat_f32x4], &imm, codeIndex, -1})
		case IR.OPCi32x4_trunc_u_sat_f32x4:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32x4_trunc_u_sat_f32x4], &imm, codeIndex, -1})
		case IR.OPCi64x2_trunc_s_sat_f64x2:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64x2_trunc_s_sat_f64x2], &imm, codeIndex, -1})
		case IR.OPCi64x2_trunc_u_sat_f64x2:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64x2_trunc_u_sat_f64x2], &imm, codeIndex, -1})
		case IR.OPCatomic_wake:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCatomic_wake], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_wait:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_wait], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_wait:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_wait], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_load:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_load], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_load:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_load], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_load8_u:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_load8_u], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_load16_u:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_load16_u], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_load8_u:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_load8_u], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_load16_u:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_load16_u], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_load32_u:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_load32_u], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_store:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_store], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_store:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_store], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_store8:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_store8], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_store16:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_store16], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_store8:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_store8], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_store16:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_store16], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_store32:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_store32], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw_add], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw_add], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw8_u_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw8_u_add], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw16_u_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw16_u_add], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw8_u_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw8_u_add], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw16_u_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw16_u_add], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw32_u_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw32_u_add], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw_sub], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw_sub], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw8_u_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw8_u_sub], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw16_u_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw16_u_sub], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw8_u_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw8_u_sub], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw16_u_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw16_u_sub], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw32_u_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw32_u_sub], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw_and], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw_and], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw8_u_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw8_u_and], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw16_u_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw16_u_and], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw8_u_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw8_u_and], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw16_u_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw16_u_and], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw32_u_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw32_u_and], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw_or], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw_or], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw8_u_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw8_u_or], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw16_u_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw16_u_or], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw8_u_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw8_u_or], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw16_u_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw16_u_or], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw32_u_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw32_u_or], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw_xor], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw_xor], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw8_u_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw8_u_xor], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw16_u_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw16_u_xor], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw8_u_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw8_u_xor], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw16_u_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw16_u_xor], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw32_u_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw32_u_xor], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw_xchg], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw_xchg], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw8_u_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw8_u_xchg], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw16_u_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw16_u_xchg], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw8_u_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw8_u_xchg], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw16_u_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw16_u_xchg], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw32_u_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw32_u_xchg], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw_cmpxchg], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw_cmpxchg], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw8_u_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw8_u_cmpxchg], &imm, codeIndex, -1})
		case IR.OPCi32_atomic_rmw16_u_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi32_atomic_rmw16_u_cmpxchg], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw8_u_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw8_u_cmpxchg], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw16_u_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw16_u_cmpxchg], &imm, codeIndex, -1})
		case IR.OPCi64_atomic_rmw32_u_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCi64_atomic_rmw32_u_cmpxchg], &imm, codeIndex, -1})
		case IR.OPCblock:
			imm := IR.ControlStructureImm{}
			err = DecodeControlStructureImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCblock], &imm, codeIndex, -1})
		case IR.OPCloop:
			imm := IR.ControlStructureImm{}
			err = DecodeControlStructureImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCloop], &imm, codeIndex, -1})
		case IR.OPCif_:
			imm := IR.ControlStructureImm{}
			err = DecodeControlStructureImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCif_], &imm, codeIndex, -1})
		case IR.OPCelse_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCelse_], &imm, codeIndex, -1})
		case IR.OPCend:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCend], &imm, codeIndex, -1})
		case IR.OPCtry_:
			imm := IR.ControlStructureImm{}
			err = DecodeControlStructureImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCtry_], &imm, codeIndex, -1})
		case IR.OPCcatch_:
			imm := IR.ExceptionTypeImm{}
			err = DecodeExceptionTypeImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCcatch_], &imm, codeIndex, -1})
		case IR.OPCcatch_all:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			ins = append(ins, IR.Instruction{&IR.Ops[IR.OPCcatch_all], &imm, codeIndex, -1})
		}
		codeIndex += 1
	}
	if ins[len(ins)-1].Op.Code != IR.OPCend {
		return nil, fmt.Errorf("code not end with \"end\"")
	}
	return ins, nil
}
