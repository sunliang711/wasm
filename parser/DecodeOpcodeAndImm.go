// Note this file is created by makeType.sh
// Don't modify it directly
package parser

import (
	"bytes"
	"encoding/binary"
	"io"
	"wasm/types/IR"
)

func DecodeOpcodeAndImm(opcodeBytes []byte, funcDef *IR.FunctionDef) ([]byte, error) {
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
		switch IR.Opcode(opc) {
		case IR.OPCunreachable:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCunreachable
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCbr:
			imm := IR.BranchImm{}
			err = DecodeBranchImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_BranchImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCbr
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCbr_if:
			imm := IR.BranchImm{}
			err = DecodeBranchImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_BranchImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCbr_if
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCbr_table:
			imm := IR.BranchTableImm{}
			err = DecodeBranchTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_BranchTableImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCbr_table
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCreturn_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCreturn_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCcall:
			imm := IR.FunctionImm{}
			err = DecodeFunctionImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_FunctionImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCcall
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCcall_indirect:
			imm := IR.CallIndirectImm{}
			err = DecodeCallIndirectImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_CallIndirectImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCcall_indirect
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCdrop:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCdrop
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCselect:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCselect
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCget_local:
			imm := IR.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_GetOrSetVariableImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCget_local
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCset_local:
			imm := IR.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_GetOrSetVariableImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCset_local
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCtee_local:
			imm := IR.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_GetOrSetVariableImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCtee_local
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCget_global:
			imm := IR.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_GetOrSetVariableImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCget_global
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCset_global:
			imm := IR.GetOrSetVariableImm{}
			err = DecodeGetOrSetVariableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_GetOrSetVariableImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCset_global
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCtable_get:
			imm := IR.TableImm{}
			err = DecodeTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_TableImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCtable_get
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCtable_set:
			imm := IR.TableImm{}
			err = DecodeTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_TableImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCtable_set
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCthrow_:
			imm := IR.ExceptionTypeImm{}
			err = DecodeExceptionTypeImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_ExceptionTypeImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCthrow_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCrethrow:
			imm := IR.RethrowImm{}
			err = DecodeRethrowImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_RethrowImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCrethrow
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCnop:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCnop
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_load:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_load:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_load:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_load:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_load8_s:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_load8_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_load8_u:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_load8_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_load16_s:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_load16_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_load16_u:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_load16_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_load8_s:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_load8_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_load8_u:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_load8_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_load16_s:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_load16_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_load16_u:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_load16_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_load32_s:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_load32_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_load32_u:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_load32_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_store:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_store:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_store:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_store:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_store8:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_store8
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_store16:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_store16
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_store8:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_store8
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_store16:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_store16
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_store32:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_store32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCmemory_size:
			imm := IR.MemoryImm{}
			err = DecodeMemoryImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_MemoryImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCmemory_size
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCmemory_grow:
			imm := IR.MemoryImm{}
			err = DecodeMemoryImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_MemoryImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCmemory_grow
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_const:
			imm := IR.LiteralImm_I32{}
			err = DecodeLiteralImm_I32(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LiteralImm_I32{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_const
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_const:
			imm := IR.LiteralImm_I64{}
			err = DecodeLiteralImm_I64(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LiteralImm_I64{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_const
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_const:
			imm := IR.LiteralImm_F32{}
			err = DecodeLiteralImm_F32(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LiteralImm_F32{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_const
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_const:
			imm := IR.LiteralImm_F64{}
			err = DecodeLiteralImm_F64(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LiteralImm_F64{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_const
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_eqz:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_eqz
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_lt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_lt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_lt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_lt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_gt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_gt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_gt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_gt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_le_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_le_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_le_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_le_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_ge_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_ge_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_ge_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_ge_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_eqz:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_eqz
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_lt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_lt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_lt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_lt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_gt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_gt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_gt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_gt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_le_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_le_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_le_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_le_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_ge_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_ge_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_ge_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_ge_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_lt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_lt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_gt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_gt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_le:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_le
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_ge:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_ge
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_lt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_lt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_gt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_gt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_le:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_le
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_ge:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_ge
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_clz:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_clz
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_ctz:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_ctz
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_popcnt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_popcnt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_div_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_div_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_div_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_div_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_rem_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_rem_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_rem_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_rem_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_and_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_and_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_or_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_or_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_xor_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_xor_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_shl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_shl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_shr_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_shr_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_shr_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_shr_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_rotl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_rotl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_rotr:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_rotr
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_clz:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_clz
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_ctz:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_ctz
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_popcnt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_popcnt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_div_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_div_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_div_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_div_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_rem_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_rem_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_rem_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_rem_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_and_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_and_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_or_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_or_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_xor_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_xor_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_shl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_shl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_shr_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_shr_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_shr_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_shr_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_rotl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_rotl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_rotr:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_rotr
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_abs:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_abs
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_ceil:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_ceil
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_floor:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_floor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_trunc:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_trunc
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_nearest:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_nearest
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_sqrt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_sqrt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_div:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_div
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_min:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_min
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_max:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_max
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_copysign:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_copysign
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_abs:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_abs
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_ceil:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_ceil
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_floor:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_floor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_trunc:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_trunc
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_nearest:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_nearest
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_sqrt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_sqrt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_div:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_div
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_min:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_min
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_max:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_max
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_copysign:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_copysign
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_wrap_i64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_wrap_i64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_trunc_s_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_trunc_s_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_trunc_u_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_trunc_u_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_trunc_s_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_trunc_s_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_trunc_u_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_trunc_u_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_extend_s_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_extend_s_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_extend_u_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_extend_u_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_trunc_s_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_trunc_s_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_trunc_u_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_trunc_u_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_trunc_s_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_trunc_s_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_trunc_u_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_trunc_u_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_convert_s_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_convert_s_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_convert_u_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_convert_u_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_convert_s_i64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_convert_s_i64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_convert_u_i64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_convert_u_i64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_demote_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_demote_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_convert_s_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_convert_s_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_convert_u_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_convert_u_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_convert_s_i64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_convert_s_i64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_convert_u_i64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_convert_u_i64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_promote_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_promote_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_reinterpret_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_reinterpret_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_reinterpret_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_reinterpret_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32_reinterpret_i32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32_reinterpret_i32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64_reinterpret_i64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64_reinterpret_i64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_extend8_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_extend8_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_extend16_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_extend16_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_extend8_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_extend8_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_extend16_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_extend16_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_extend32_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_extend32_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCref_null:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCref_null
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCref_isnull:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCref_isnull
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCref_func:
			imm := IR.FunctionImm{}
			err = DecodeFunctionImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_FunctionImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCref_func
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_trunc_s_sat_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_trunc_s_sat_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_trunc_u_sat_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_trunc_u_sat_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_trunc_s_sat_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_trunc_s_sat_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_trunc_u_sat_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_trunc_u_sat_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_trunc_s_sat_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_trunc_s_sat_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_trunc_u_sat_f32:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_trunc_u_sat_f32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_trunc_s_sat_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_trunc_s_sat_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_trunc_u_sat_f64:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_trunc_u_sat_f64
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCmemory_init:
			imm := IR.DataSegmentAndMemImm{}
			err = DecodeDataSegmentAndMemImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_DataSegmentAndMemImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCmemory_init
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCmemory_drop:
			imm := IR.DataSegmentImm{}
			err = DecodeDataSegmentImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_DataSegmentImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCmemory_drop
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCmemory_copy:
			imm := IR.MemoryImm{}
			err = DecodeMemoryImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_MemoryImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCmemory_copy
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCmemory_fill:
			imm := IR.MemoryImm{}
			err = DecodeMemoryImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_MemoryImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCmemory_fill
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCtable_init:
			imm := IR.ElemSegmentAndTableImm{}
			err = DecodeElemSegmentAndTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_ElemSegmentAndTableImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCtable_init
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCtable_drop:
			imm := IR.ElemSegmentImm{}
			err = DecodeElemSegmentImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_ElemSegmentImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCtable_drop
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCtable_copy:
			imm := IR.TableImm{}
			err = DecodeTableImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_TableImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCtable_copy
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCv128_const:
			imm := IR.LiteralImm_V128{}
			err = DecodeLiteralImm_V128(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LiteralImm_V128{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCv128_const
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCv128_load:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCv128_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCv128_store:
			imm := IR.LoadOrStoreImm{}
			err = DecodeLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCv128_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_splat:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_splat
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_splat:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_splat
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_splat:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_splat
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64x2_splat:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64x2_splat
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_splat:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_splat
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_splat:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_splat
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_extract_lane_s:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_extract_lane_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_extract_lane_u:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_extract_lane_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_extract_lane_s:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_extract_lane_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_extract_lane_u:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_extract_lane_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_extract_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_extract_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64x2_extract_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64x2_extract_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_extract_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_extract_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_extract_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_extract_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_replace_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_replace_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_replace_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_replace_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_replace_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_replace_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64x2_replace_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64x2_replace_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_replace_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_replace_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_replace_lane:
			imm := IR.LaneIndexImm{}
			err = DecodeLaneIndexImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_LaneIndexImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_replace_lane
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCv8x16_shuffle:
			imm := IR.ShuffleImm_16{}
			err = DecodeShuffleImm_16(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_ShuffleImm_16{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCv8x16_shuffle
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64x2_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64x2_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64x2_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64x2_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64x2_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64x2_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_add_saturate_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_add_saturate_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_add_saturate_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_add_saturate_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_add_saturate_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_add_saturate_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_add_saturate_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_add_saturate_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_sub_saturate_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_sub_saturate_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_sub_saturate_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_sub_saturate_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_sub_saturate_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_sub_saturate_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_sub_saturate_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_sub_saturate_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_shl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_shl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_shl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_shl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_shl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_shl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64x2_shl:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64x2_shl
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_shr_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_shr_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_shr_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_shr_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_shr_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_shr_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_shr_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_shr_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_shr_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_shr_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_shr_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_shr_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64x2_shr_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64x2_shr_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64x2_shr_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64x2_shr_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCv128_and:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCv128_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCv128_or:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCv128_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCv128_xor:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCv128_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCv128_not:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCv128_not
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCv128_bitselect:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCv128_bitselect
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_any_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_any_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_any_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_any_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_any_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_any_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64x2_any_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64x2_any_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_all_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_all_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_all_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_all_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_all_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_all_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64x2_all_true:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64x2_all_true
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_eq:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_eq
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_ne:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_ne
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_lt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_lt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_lt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_lt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_lt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_lt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_lt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_lt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_lt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_lt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_lt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_lt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_lt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_lt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_lt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_lt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_le_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_le_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_le_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_le_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_le_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_le_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_le_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_le_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_le_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_le_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_le_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_le_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_le:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_le
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_le:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_le
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_gt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_gt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_gt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_gt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_gt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_gt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_gt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_gt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_gt_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_gt_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_gt_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_gt_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_gt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_gt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_gt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_gt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_ge_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_ge_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi8x16_ge_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi8x16_ge_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_ge_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_ge_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi16x8_ge_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi16x8_ge_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_ge_s:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_ge_s
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_ge_u:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_ge_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_ge:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_ge
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_ge:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_ge
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_neg:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_neg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_abs:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_abs
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_abs:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_abs
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_min:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_min
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_min:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_min
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_max:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_max
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_max:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_max
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_add:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_sub:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_div:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_div
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_div:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_div
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_mul:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_mul
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_sqrt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_sqrt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_sqrt:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_sqrt
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_convert_s_i32x4:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_convert_s_i32x4
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf32x4_convert_u_i32x4:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf32x4_convert_u_i32x4
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_convert_s_i64x2:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_convert_s_i64x2
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCf64x2_convert_u_i64x2:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCf64x2_convert_u_i64x2
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_trunc_s_sat_f32x4:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_trunc_s_sat_f32x4
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32x4_trunc_u_sat_f32x4:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32x4_trunc_u_sat_f32x4
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64x2_trunc_s_sat_f64x2:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64x2_trunc_s_sat_f64x2
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64x2_trunc_u_sat_f64x2:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64x2_trunc_u_sat_f64x2
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCatomic_wake:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCatomic_wake
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_wait:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_wait
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_wait:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_wait
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_load:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_load:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_load
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_load8_u:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_load8_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_load16_u:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_load16_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_load8_u:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_load8_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_load16_u:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_load16_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_load32_u:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_load32_u
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_store:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_store:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_store
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_store8:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_store8
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_store16:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_store16
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_store8:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_store8
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_store16:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_store16
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_store32:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_store32
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw8_u_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw8_u_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw16_u_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw16_u_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw8_u_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw8_u_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw16_u_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw16_u_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw32_u_add:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw32_u_add
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw8_u_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw8_u_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw16_u_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw16_u_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw8_u_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw8_u_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw16_u_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw16_u_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw32_u_sub:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw32_u_sub
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw8_u_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw8_u_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw16_u_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw16_u_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw8_u_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw8_u_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw16_u_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw16_u_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw32_u_and:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw32_u_and
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw8_u_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw8_u_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw16_u_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw16_u_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw8_u_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw8_u_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw16_u_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw16_u_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw32_u_or:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw32_u_or
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw8_u_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw8_u_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw16_u_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw16_u_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw8_u_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw8_u_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw16_u_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw16_u_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw32_u_xor:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw32_u_xor
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw8_u_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw8_u_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw16_u_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw16_u_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw8_u_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw8_u_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw16_u_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw16_u_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw32_u_xchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw32_u_xchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw8_u_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw8_u_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi32_atomic_rmw16_u_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi32_atomic_rmw16_u_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw8_u_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw8_u_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw16_u_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw16_u_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCi64_atomic_rmw32_u_cmpxchg:
			imm := IR.AtomicLoadOrStoreImm{}
			err = DecodeAtomicLoadOrStoreImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_AtomicLoadOrStoreImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCi64_atomic_rmw32_u_cmpxchg
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCblock:
			imm := IR.ControlStructureImm{}
			err = DecodeControlStructureImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_ControlStructureImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCblock
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCloop:
			imm := IR.ControlStructureImm{}
			err = DecodeControlStructureImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_ControlStructureImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCloop
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCif_:
			imm := IR.ControlStructureImm{}
			err = DecodeControlStructureImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_ControlStructureImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCif_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCelse_:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCelse_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCend:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCend
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCtry_:
			imm := IR.ControlStructureImm{}
			err = DecodeControlStructureImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_ControlStructureImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCtry_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCcatch_:
			imm := IR.ExceptionTypeImm{}
			err = DecodeExceptionTypeImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_ExceptionTypeImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCcatch_
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		case IR.OPCcatch_all:
			imm := IR.NoImm{}
			err = DecodeNoImm(rd, &imm, funcDef)
			if err != nil {
				return nil, err
			}
			opimm := IR.OpcodeAndImm_NoImm{}
			opimm.Imm = imm
			opimm.Opcode = IR.OPCcatch_all
			err = binary.Write(&buf, binary.LittleEndian, &opimm)
			if err != nil {
				return nil, err
			}
			ret = append(ret, buf.Bytes()...)
		}
	}
	return ret, nil
}
