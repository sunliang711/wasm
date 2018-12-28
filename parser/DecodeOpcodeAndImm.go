package parser

import (
	"bytes"
	"encoding/binary"
	"io"
	"wasm/types"
)

func DecodeOpcodeAndImm(opcodeBytes []byte, funcDef *types.FunctionDef) ([]byte, error) {
	rd := bytes.NewReader(opcodeBytes)
	var ret []byte

	for {
		opc, err := DecodeOpcode(rd)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		var buf bytes.Buffer
		switch types.Opcode(opc) {
		case types.OPCunreachable:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCunreachable
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCbr:
			imm := types.BranchImm{}
			err = DecodeBranchImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_BranchImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCbr
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCbr_if:
			imm := types.BranchImm{}
			err = DecodeBranchImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_BranchImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCbr_if
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCbr_table:
			imm := types.BranchTableImm{}
			err = DecodeBranchTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_BranchTableImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCbr_table
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCreturn_:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCreturn_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCcall:
			imm := types.FunctionImm{}
			err = DecodeFunctionImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_FunctionImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCcall
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCcall_indirect:
			imm := types.CallIndirectImm{}
			err = DecodeCallIndirectImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_CallIndirectImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCcall_indirect
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCdrop:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCdrop
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCselect:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCselect
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCget_local:
			imm := types.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_GetOrSetVariableImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCget_local
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCset_local:
			imm := types.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_GetOrSetVariableImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCset_local
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCtee_local:
			imm := types.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_GetOrSetVariableImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCtee_local
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCget_global:
			imm := types.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_GetOrSetVariableImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCget_global
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCset_global:
			imm := types.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_GetOrSetVariableImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCset_global
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCtable_get:
			imm := types.TableImm{}
			err = DecodeTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_TableImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCtable_get
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCtable_set:
			imm := types.TableImm{}
			err = DecodeTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_TableImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCtable_set
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCthrow_:
			imm := types.ExceptionTypeImm{}
			err = DecodeExceptionTypeImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_ExceptionTypeImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCthrow_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCrethrow:
			imm := types.RethrowImm{}
			err = DecodeRethrowImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_RethrowImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCrethrow
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCnop:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCnop
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_load:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_load:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_load:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_load:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_load8_s:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_load8_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_load8_u:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_load8_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_load16_s:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_load16_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_load16_u:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_load16_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_load8_s:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_load8_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_load8_u:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_load8_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_load16_s:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_load16_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_load16_u:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_load16_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_load32_s:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_load32_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_load32_u:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_load32_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_store:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_store:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_store:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_store:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_store8:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_store8
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_store16:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_store16
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_store8:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_store8
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_store16:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_store16
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_store32:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_store32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCmemory_size:
			imm := types.MemoryImm{}
			err = DecodeMemoryImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_MemoryImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCmemory_size
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCmemory_grow:
			imm := types.MemoryImm{}
			err = DecodeMemoryImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_MemoryImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCmemory_grow
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_const:
			imm := types.LiteralImm_I32{}
			err = DecodeLiteralImm_I32(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LiteralImm_I32{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_const
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_const:
			imm := types.LiteralImm_I64{}
			err = DecodeLiteralImm_I64(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LiteralImm_I64{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_const
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_const:
			imm := types.LiteralImm_F32{}
			err = DecodeLiteralImm_F32(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LiteralImm_F32{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_const
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_const:
			imm := types.LiteralImm_F64{}
			err = DecodeLiteralImm_F64(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LiteralImm_F64{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_const
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_eqz:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_eqz
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_eq:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_ne:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_lt_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_lt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_lt_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_lt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_gt_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_gt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_gt_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_gt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_le_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_le_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_le_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_le_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_ge_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_ge_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_ge_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_ge_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_eqz:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_eqz
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_eq:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_ne:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_lt_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_lt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_lt_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_lt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_gt_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_gt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_gt_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_gt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_le_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_le_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_le_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_le_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_ge_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_ge_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_ge_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_ge_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_eq:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_ne:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_lt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_lt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_gt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_gt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_le:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_le
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_ge:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_ge
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_eq:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_ne:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_lt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_lt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_gt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_gt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_le:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_le
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_ge:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_ge
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_clz:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_clz
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_ctz:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_ctz
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_popcnt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_popcnt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_add:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_sub:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_mul:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_div_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_div_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_div_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_div_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_rem_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_rem_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_rem_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_rem_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_and_:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_and_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_or_:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_or_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_xor_:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_xor_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_shl:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_shl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_shr_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_shr_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_shr_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_shr_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_rotl:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_rotl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_rotr:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_rotr
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_clz:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_clz
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_ctz:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_ctz
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_popcnt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_popcnt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_add:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_sub:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_mul:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_div_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_div_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_div_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_div_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_rem_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_rem_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_rem_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_rem_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_and_:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_and_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_or_:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_or_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_xor_:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_xor_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_shl:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_shl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_shr_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_shr_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_shr_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_shr_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_rotl:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_rotl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_rotr:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_rotr
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_abs:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_abs
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_neg:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_ceil:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_ceil
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_floor:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_floor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_trunc:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_trunc
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_nearest:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_nearest
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_sqrt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_sqrt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_add:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_sub:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_mul:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_div:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_div
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_min:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_min
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_max:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_max
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_copysign:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_copysign
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_abs:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_abs
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_neg:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_ceil:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_ceil
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_floor:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_floor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_trunc:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_trunc
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_nearest:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_nearest
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_sqrt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_sqrt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_add:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_sub:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_mul:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_div:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_div
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_min:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_min
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_max:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_max
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_copysign:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_copysign
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_wrap_i64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_wrap_i64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_trunc_s_f32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_trunc_s_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_trunc_u_f32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_trunc_u_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_trunc_s_f64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_trunc_s_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_trunc_u_f64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_trunc_u_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_extend_s_i32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_extend_s_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_extend_u_i32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_extend_u_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_trunc_s_f32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_trunc_s_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_trunc_u_f32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_trunc_u_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_trunc_s_f64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_trunc_s_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_trunc_u_f64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_trunc_u_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_convert_s_i32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_convert_s_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_convert_u_i32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_convert_u_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_convert_s_i64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_convert_s_i64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_convert_u_i64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_convert_u_i64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_demote_f64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_demote_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_convert_s_i32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_convert_s_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_convert_u_i32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_convert_u_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_convert_s_i64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_convert_s_i64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_convert_u_i64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_convert_u_i64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_promote_f32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_promote_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_reinterpret_f32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_reinterpret_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_reinterpret_f64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_reinterpret_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32_reinterpret_i32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32_reinterpret_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64_reinterpret_i64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64_reinterpret_i64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_extend8_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_extend8_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_extend16_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_extend16_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_extend8_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_extend8_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_extend16_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_extend16_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_extend32_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_extend32_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCref_null:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCref_null
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCref_isnull:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCref_isnull
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCref_func:
			imm := types.FunctionImm{}
			err = DecodeFunctionImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_FunctionImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCref_func
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_trunc_s_sat_f32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_trunc_s_sat_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_trunc_u_sat_f32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_trunc_u_sat_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_trunc_s_sat_f64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_trunc_s_sat_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_trunc_u_sat_f64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_trunc_u_sat_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_trunc_s_sat_f32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_trunc_s_sat_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_trunc_u_sat_f32:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_trunc_u_sat_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_trunc_s_sat_f64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_trunc_s_sat_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_trunc_u_sat_f64:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_trunc_u_sat_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCmemory_init:
			imm := types.DataSegmentAndMemImm{}
			err = DecodeDataSegmentAndMemImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_DataSegmentAndMemImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCmemory_init
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCmemory_drop:
			imm := types.DataSegmentImm{}
			err = DecodeDataSegmentImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_DataSegmentImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCmemory_drop
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCmemory_copy:
			imm := types.MemoryImm{}
			err = DecodeMemoryImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_MemoryImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCmemory_copy
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCmemory_fill:
			imm := types.MemoryImm{}
			err = DecodeMemoryImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_MemoryImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCmemory_fill
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCtable_init:
			imm := types.ElemSegmentAndTableImm{}
			err = DecodeElemSegmentAndTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_ElemSegmentAndTableImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCtable_init
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCtable_drop:
			imm := types.ElemSegmentImm{}
			err = DecodeElemSegmentImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_ElemSegmentImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCtable_drop
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCtable_copy:
			imm := types.TableImm{}
			err = DecodeTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_TableImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCtable_copy
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCv128_const:
			imm := types.LiteralImm_V128{}
			err = DecodeLiteralImm_V128(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LiteralImm_V128{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCv128_const
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCv128_load:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCv128_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCv128_store:
			imm := types.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCv128_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_splat:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_splat
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_splat:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_splat
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_splat:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_splat
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64x2_splat:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64x2_splat
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_splat:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_splat
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_splat:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_splat
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_extract_lane_s:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_extract_lane_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_extract_lane_u:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_extract_lane_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_extract_lane_s:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_extract_lane_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_extract_lane_u:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_extract_lane_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_extract_lane:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_extract_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64x2_extract_lane:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64x2_extract_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_extract_lane:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_extract_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_extract_lane:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_extract_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_replace_lane:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_replace_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_replace_lane:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_replace_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_replace_lane:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_replace_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64x2_replace_lane:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64x2_replace_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_replace_lane:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_replace_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_replace_lane:
			imm := types.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_replace_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCv8x16_shuffle:
			imm := types.ShuffleImm_16{}
			err = DecodeShuffleImm_16(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_ShuffleImm_16{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCv8x16_shuffle
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_add:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_add:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_add:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64x2_add:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64x2_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_sub:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_sub:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_sub:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64x2_sub:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64x2_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_mul:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_mul:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_mul:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_neg:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_neg:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_neg:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64x2_neg:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64x2_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_add_saturate_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_add_saturate_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_add_saturate_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_add_saturate_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_add_saturate_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_add_saturate_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_add_saturate_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_add_saturate_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_sub_saturate_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_sub_saturate_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_sub_saturate_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_sub_saturate_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_sub_saturate_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_sub_saturate_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_sub_saturate_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_sub_saturate_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_shl:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_shl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_shl:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_shl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_shl:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_shl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64x2_shl:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64x2_shl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_shr_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_shr_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_shr_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_shr_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_shr_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_shr_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_shr_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_shr_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_shr_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_shr_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_shr_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_shr_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64x2_shr_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64x2_shr_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64x2_shr_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64x2_shr_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCv128_and:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCv128_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCv128_or:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCv128_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCv128_xor:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCv128_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCv128_not:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCv128_not
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCv128_bitselect:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCv128_bitselect
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_any_true:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_any_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_any_true:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_any_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_any_true:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_any_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64x2_any_true:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64x2_any_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_all_true:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_all_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_all_true:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_all_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_all_true:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_all_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64x2_all_true:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64x2_all_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_eq:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_eq:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_eq:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_eq:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_eq:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_ne:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_ne:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_ne:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_ne:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_ne:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_lt_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_lt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_lt_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_lt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_lt_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_lt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_lt_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_lt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_lt_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_lt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_lt_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_lt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_lt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_lt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_lt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_lt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_le_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_le_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_le_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_le_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_le_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_le_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_le_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_le_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_le_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_le_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_le_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_le_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_le:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_le
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_le:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_le
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_gt_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_gt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_gt_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_gt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_gt_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_gt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_gt_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_gt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_gt_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_gt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_gt_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_gt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_gt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_gt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_gt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_gt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_ge_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_ge_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi8x16_ge_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi8x16_ge_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_ge_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_ge_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi16x8_ge_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi16x8_ge_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_ge_s:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_ge_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_ge_u:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_ge_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_ge:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_ge
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_ge:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_ge
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_neg:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_neg:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_abs:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_abs
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_abs:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_abs
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_min:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_min
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_min:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_min
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_max:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_max
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_max:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_max
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_add:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_add:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_sub:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_sub:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_div:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_div
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_div:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_div
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_mul:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_mul:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_sqrt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_sqrt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_sqrt:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_sqrt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_convert_s_i32x4:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_convert_s_i32x4
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf32x4_convert_u_i32x4:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf32x4_convert_u_i32x4
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_convert_s_i64x2:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_convert_s_i64x2
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCf64x2_convert_u_i64x2:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCf64x2_convert_u_i64x2
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_trunc_s_sat_f32x4:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_trunc_s_sat_f32x4
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32x4_trunc_u_sat_f32x4:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32x4_trunc_u_sat_f32x4
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64x2_trunc_s_sat_f64x2:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64x2_trunc_s_sat_f64x2
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64x2_trunc_u_sat_f64x2:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64x2_trunc_u_sat_f64x2
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCatomic_wake:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCatomic_wake
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_wait:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_wait
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_wait:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_wait
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_load:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_load:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_load8_u:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_load8_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_load16_u:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_load16_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_load8_u:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_load8_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_load16_u:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_load16_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_load32_u:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_load32_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_store:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_store:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_store8:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_store8
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_store16:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_store16
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_store8:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_store8
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_store16:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_store16
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_store32:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_store32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw_add:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw_add:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw8_u_add:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw8_u_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw16_u_add:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw16_u_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw8_u_add:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw8_u_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw16_u_add:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw16_u_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw32_u_add:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw32_u_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw_sub:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw_sub:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw8_u_sub:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw8_u_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw16_u_sub:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw16_u_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw8_u_sub:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw8_u_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw16_u_sub:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw16_u_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw32_u_sub:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw32_u_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw_and:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw_and:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw8_u_and:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw8_u_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw16_u_and:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw16_u_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw8_u_and:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw8_u_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw16_u_and:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw16_u_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw32_u_and:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw32_u_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw_or:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw_or:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw8_u_or:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw8_u_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw16_u_or:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw16_u_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw8_u_or:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw8_u_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw16_u_or:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw16_u_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw32_u_or:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw32_u_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw_xor:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw_xor:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw8_u_xor:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw8_u_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw16_u_xor:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw16_u_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw8_u_xor:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw8_u_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw16_u_xor:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw16_u_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw32_u_xor:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw32_u_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw_xchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw_xchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw8_u_xchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw8_u_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw16_u_xchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw16_u_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw8_u_xchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw8_u_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw16_u_xchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw16_u_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw32_u_xchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw32_u_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw_cmpxchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw_cmpxchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw8_u_cmpxchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw8_u_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi32_atomic_rmw16_u_cmpxchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi32_atomic_rmw16_u_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw8_u_cmpxchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw8_u_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw16_u_cmpxchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw16_u_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCi64_atomic_rmw32_u_cmpxchg:
			imm := types.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCi64_atomic_rmw32_u_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCblock:
			imm := types.ControlStructureImm{}
			err = DecodeControlStructureImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_ControlStructureImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCblock
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCloop:
			imm := types.ControlStructureImm{}
			err = DecodeControlStructureImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_ControlStructureImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCloop
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCif_:
			imm := types.ControlStructureImm{}
			err = DecodeControlStructureImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_ControlStructureImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCif_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCelse_:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCelse_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCend:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCend
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCtry_:
			imm := types.ControlStructureImm{}
			err = DecodeControlStructureImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_ControlStructureImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCtry_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCcatch_:
			imm := types.ExceptionTypeImm{}
			err = DecodeExceptionTypeImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_ExceptionTypeImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCcatch_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case types.OPCcatch_all:
			imm := types.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := types.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = types.OPCcatch_all
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		}
	}
	return ret, nil
}
