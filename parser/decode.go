package parser

import (
	"fmt"
	"io"
	"wasm/types"
	"wasm/utils"
)

func DecodeReferenceType(rd io.Reader) (types.RefType, error) {
	rt, err := utils.ReadByte(rd)
	if err != nil {
		return types.RTInvalid, err
	}
	switch rt {
	case 0x70:
		return types.RTAnyFunc, nil
	case 0x6f:
		return types.RTAnyRef, nil
	default:
		return types.RTInvalid, fmt.Errorf(types.ErrReferenceTypeByte)
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

func DecodeTableType(rd io.Reader) (types.TableType, error) {
	refType, err := DecodeReferenceType(rd)
	if err != nil {
		return types.TableType{}, err
	}
	isShared, min, max, err := DecodeFlags(rd)
	if err != nil {
		return types.TableType{}, err
	}
	return types.TableType{ElementType: refType, IsShared: isShared, Size: types.SizeConstraints{min, max}}, nil

}

func DecodeGlobalType(rd io.Reader) (types.GlobalType, error) {
	var (
		globalType types.GlobalType
	)
	// A. valueType
	vType, err := DecodeValueType(rd)
	if err != nil {
		return types.GlobalType{}, err
	}
	globalType.ValType = vType

	//B. isMutable
	var isMutable byte
	_, err = utils.DecodeVarInt(rd, 1, &isMutable)
	if err != nil {
		return types.GlobalType{}, err
	}
	globalType.IsMutable = isMutable != 0

	return globalType, nil
}

func DecodeInitializer(rd io.Reader) (types.InitializerExpression, error) {
	initExpression := types.InitializerExpression{}

	//1. opcode
	opcode, err := DecodeOpcode(rd)
	if err != nil {
		return initExpression, err
	}
	initExpression.Type = types.InitializerType(opcode)

	//2. switch initializer
	switch initExpression.Type {
	case types.I32_const:
		_, err = utils.DecodeVarInt(rd, 32, &initExpression.I32)
		if err != nil {
			return initExpression, err
		}
	case types.I64_const:
		_, err = utils.DecodeVarInt(rd, 64, &initExpression.I64)
		if err != nil {
			return initExpression, err
		}
	case types.F32_const:
		f32Bytes, err := utils.ReadNByte(rd, 4)
		if err != nil {
			return initExpression, err
		}
		f32Value, err := utils.Bytes2float32(f32Bytes, true)
		if err != nil {
			return initExpression, err
		}
		initExpression.F32 = f32Value
	case types.F64_const:
		f64Bytes, err := utils.ReadNByte(rd, 8)
		if err != nil {
			return initExpression, err
		}
		f64Value, err := utils.Bytes2float64(f64Bytes, true)
		if err != nil {
			return initExpression, err
		}
		initExpression.F64 = f64Value
	case types.V128_const:
		v128Bytes, err := utils.ReadNByte(rd, 16)
		if err != nil {
			return initExpression, err
		}
		copy(initExpression.V128[:], v128Bytes)
	case types.Get_global:
		var gref uint32
		_, err := utils.DecodeVarInt(rd, 32, &gref)
		if err != nil {
			return initExpression, err
		}
		initExpression.GlobalRef = uint64(gref)
	case types.Ref_null:
	default:
		return initExpression, fmt.Errorf(types.ErrInvalidInitializerExpressionOpcode)
	}

	return initExpression, nil
}

func DecodeOpcode(rd io.Reader) (int16, error) {
	var (
		opcode int16
	)
	byte0, err := utils.ReadByte(rd)
	if err != nil {
		return 0, err
	}
	opcode = int16(byte0)
	if opcode > int16(types.OPCMaxSingleByteOpcode) {
		byte1, err := utils.ReadByte(rd)
		if err != nil {
			return 0, err
		}
		opcode = opcode << 8
		opcode = opcode | int16(byte1)
	}
	return opcode, nil
}

func DecodeLocalSet(rd io.Reader, ls *types.LocalSet) (int, error) {
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
	vType, err := DecodeValueType(rd)
	if err != nil {
		return 0, err
	}
	ls.Type = vType

	//used n + 1 bytes in total
	return n + 1, nil
}
