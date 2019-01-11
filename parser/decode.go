package parser

import (
	"fmt"
	"io"
	"wasm/types"
	"wasm/types/IR"
	"wasm/utils"
)

func DecodeReferenceType(rd io.Reader) (IR.RefType, error) {
	rt, err := utils.ReadByte(rd)
	if err != nil {
		return IR.RTInvalid, err
	}
	switch rt {
	case 0x70:
		return IR.RTAnyFunc, nil
	case 0x6f:
		return IR.RTAnyRef, nil
	default:
		return IR.RTInvalid, fmt.Errorf(types.ErrReferenceTypeByte)
	}
}

func DecodeFlags(rd io.Reader) (bool, uint64, uint64, error) {
	var flags uint32
	_, err := utils.DecodeVarInt(rd, 32, &flags)
	if err != nil {
		return false, 0, 0, err
	}
	isShared := flags&0x02 != 0
	var min uint32
	_, err = utils.DecodeVarInt(rd, 32, &min)
	if err != nil {
		return false, 0, 0, err
	}
	var (
		max uint64
	)
	hasMax := flags&0x01 != 0
	if hasMax {
		var max uint64
		_, err = utils.DecodeVarInt(rd, 64, &max)
		if err != nil {
			return false, 0, 0, err
		}
	} else {
		max = types.UINT64_MAX
	}
	return isShared, uint64(min), max, nil
}

func DecodeTableType(rd io.Reader) (IR.TableType, error) {
	refType, err := DecodeReferenceType(rd)
	if err != nil {
		return IR.TableType{}, err
	}
	isShared, min, max, err := DecodeFlags(rd)
	if err != nil {
		return IR.TableType{}, err
	}
	return IR.TableType{ElementType: refType, IsShared: isShared, Size: IR.SizeConstraints{min, max}}, nil

}

func DecodeGlobalType(rd io.Reader) (IR.GlobalType, error) {
	var (
		globalType IR.GlobalType
	)
	// A. valueType
	vType, err := DecodeValueTypeFromReader(rd)
	if err != nil {
		return IR.GlobalType{}, err
	}
	globalType.ValType = vType

	//B. isMutable
	var isMutable byte
	_, err = utils.DecodeVarInt(rd, 1, &isMutable)
	if err != nil {
		return IR.GlobalType{}, err
	}
	globalType.IsMutable = isMutable != 0

	return globalType, nil
}

func DecodeInitializer(rd io.Reader) (IR.InitializerExpression, error) {
	initExpression := IR.InitializerExpression{}

	//1. opcode
	opcode, err := DecodeOpcode(rd)
	if err != nil {
		return initExpression, err
	}
	initExpression.Type = IR.InitializerType(opcode)

	//2. switch initializer
	switch initExpression.Type {
	case IR.I32_const:
		_, err = utils.DecodeVarInt(rd, 32, &initExpression.I32)
		if err != nil {
			return initExpression, err
		}
	case IR.I64_const:
		_, err = utils.DecodeVarInt(rd, 64, &initExpression.I64)
		if err != nil {
			return initExpression, err
		}
	case IR.F32_const:
		f32Bytes, err := utils.ReadNByte(rd, 4)
		if err != nil {
			return initExpression, err
		}
		f32Value, err := utils.Bytes2float32(f32Bytes, true)
		if err != nil {
			return initExpression, err
		}
		initExpression.F32 = f32Value
	case IR.F64_const:
		f64Bytes, err := utils.ReadNByte(rd, 8)
		if err != nil {
			return initExpression, err
		}
		f64Value, err := utils.Bytes2float64(f64Bytes, true)
		if err != nil {
			return initExpression, err
		}
		initExpression.F64 = f64Value
	case IR.V128_const:
		v128Bytes, err := utils.ReadNByte(rd, 16)
		if err != nil {
			return initExpression, err
		}
		copy(initExpression.V128[:], v128Bytes)
	case IR.Get_global:
		var gref uint32
		_, err := utils.DecodeVarInt(rd, 32, &gref)
		if err != nil {
			return initExpression, err
		}
		initExpression.GlobalRef = uint64(gref)
	case IR.Ref_null:
	default:
		return initExpression, fmt.Errorf(types.ErrInvalidInitializerExpressionOpcode)
	}
	endOp,err := utils.ReadByte(rd)
	if err != nil || IR.Opcode(endOp) != IR.OPCend{
		return initExpression,fmt.Errorf("initialzer end error")
	}

	return initExpression, nil
}

func DecodeOpcode(rd io.Reader) (uint16, error) {
	var (
		opcode uint16
	)
	byte0, err := utils.ReadByte(rd)
	if err != nil {
		return 0, err
	}
	opcode = uint16(byte0)
	if opcode > uint16(IR.OPCMaxSingleByteOpcode) {
		byte1, err := utils.ReadByte(rd)
		if err != nil {
			return 0, err
		}
		opcode = opcode << 8
		opcode = opcode | uint16(byte1)
	}
	return opcode, nil
}

func DecodeLocalSet(rd io.Reader, ls *IR.LocalSet) (int, error) {
	if ls == nil {
		return 0, fmt.Errorf(types.ErrInvalidParameter)
	}
	var num uint32
	// n bytes
	n, err := utils.DecodeVarInt(rd, 32, &num)
	if err != nil {
		return 0, err
	}
	ls.Num = uint64(num)

	// 1 byte
	vType, err := DecodeValueTypeFromReader(rd)
	if err != nil {
		return 0, err
	}
	ls.Type = vType

	//used n + 1 bytes in total
	return n + 1, nil
}
//func DecodeOpcodeAndImm(opcodeBytes []byte, funcDef *types.FunctionDef) ([]byte, error) {
//	rd := bytes.NewReader(opcodeBytes)
//	var ret []byte
//
//	for {
//		opc, err := DecodeOpcode(rd)
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			return nil, err
//		}
//		var buf bytes.Buffer
//		switch types.Opcode(opc) {
//		case types.OPCbr: //
//			imm := types.BranchImm{}//
//			err = DecodeBranchImm(rd, &imm, funcDef)//
//			if err != nil {
//				return nil, err
//			}
//			opimm := types.OpcodeAndImm_BranchImm{}//
//			opimm.Imm = imm
//			opimm.Opcode = types.OPCbr//
//			err = binary.Write(&buf, binary.LittleEndian, &opimm)
//			if err != nil {
//				return nil, err
//			}
//			ret = append(ret, buf.Bytes()...)
//
//		case types.OPCbr_if:
//		}
//	}
//	return ret, nil
//}
