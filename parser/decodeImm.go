package parser

import (
	"io"
	"wasm/types"
	"wasm/utils"
)

func DecodeNoImm(rd io.Reader, imm *types.NoImm, funcDef *types.FunctionDef) error {
	return nil

}
func DecodeControlStructureImm(rd io.Reader, imm *types.ControlStructureImm, funcDef *types.FunctionDef) error {
	var encodedBlockType int64
	_, err := utils.DecodeVarInt(rd, 32, &encodedBlockType)
	if err != nil {
		return err
	}
	switch {
	case encodedBlockType >= 0:
		imm.Type.Format = uint32(types.FormatFunctionType)
		imm.Type.Index = uint64(encodedBlockType)
	case encodedBlockType == -64:
		imm.Type.Format = uint32(types.FormatNoParametersOrResult)
		imm.Type.ResultType = uint32(types.TypeAny)
	default:
		imm.Type.Format = uint32(types.FormatOneResult)
		vt, err := DecodeValueType(int8(encodedBlockType))
		if err != nil {
			return err
		}
		imm.Type.ResultType = uint32(vt)

	}
	return nil
}

func DecodeBranchImm(rd io.Reader, imm *types.BranchImm, funcDef *types.FunctionDef) error {
	var targetDepth uint32
	_, err := utils.DecodeVarInt(rd, 32, &targetDepth)
	if err != nil {
		return err
	}
	imm.TargetDepth = uint64(targetDepth)

	return nil
}

func DecodeAtomicLoadOrStoreImm(rd io.Reader, imm *types.AtomicLoadOrStoreImm, funcDef *types.FunctionDef) error {
	var align uint8
	var offset uint32
	_, err := utils.DecodeVarInt(rd, 7, &align)
	if err != nil {
		return err
	}

	_, err = utils.DecodeVarInt(rd, 32, &offset)
	if err != nil {
		return err
	}

	imm.AlignmengLog2 = uint32(align)
	imm.Offset = uint32(offset)
	return nil
}
func DecodeBranchTableImm(rd io.Reader, imm *types.BranchTableImm, funcDef *types.FunctionDef) error {
	var branchTable []uint64
	var targetDepth uint32

	//1. num branch
	var numBranch uint32
	_, err := utils.DecodeVarInt(rd, 32, &numBranch)
	if err != nil {
		return err
	}
	//2. all branch
	for i := 0; i < int(numBranch); i++ {
		_, err = utils.DecodeVarInt(rd, 32, &targetDepth)
		if err != nil {
			return err
		}
		branchTable = append(branchTable, uint64(targetDepth))
	}

	imm.BranchTableIndex = uint64(len(funcDef.BranchTables))
	funcDef.BranchTables = append(funcDef.BranchTables, branchTable)
	_, err = utils.DecodeVarInt(rd, 32, &targetDepth)
	if err != nil {
		return err
	}
	imm.DefaultTargetDepth = uint64(targetDepth)
	return nil
}

func DecodeCallIndirectImm(rd io.Reader, imm *types.CallIndirectImm, funcDef *types.FunctionDef) error {
	var (
		typeIndex  uint32
		tableIndex uint32
	)
	_, err := utils.DecodeVarInt(rd, 32, &typeIndex)
	if err != nil {
		return err
	}
	_, err = utils.DecodeVarInt(rd, 32, &tableIndex)
	if err != nil {
		return err
	}
	imm.Type.Index = uint64(typeIndex)
	imm.TableIndex = uint64(tableIndex)
	return nil
}
func DecodeDataSegmentAndMemImm(rd io.Reader, imm *types.DataSegmentAndMemImm, funcDef *types.FunctionDef) error {
	var (
		dataSegIndex uint32
		memoryIndex  uint32
	)
	_, err := utils.DecodeVarInt(rd, 32, &dataSegIndex)
	if err != nil {
		return err
	}
	_, err = utils.DecodeVarInt(rd, 32, &memoryIndex)
	if err != nil {
		return err
	}

	imm.DataSegmentIndex = uint64(dataSegIndex)
	imm.MemoryIndex = uint64(memoryIndex)
	return nil
}

func DecodeDataSegmentImm(rd io.Reader, imm *types.DataSegmentImm, funcDef *types.FunctionDef) error {
	var dataSegIndex uint32
	_, err := utils.DecodeVarInt(rd, 32, &dataSegIndex)
	if err != nil {
		return err
	}
	imm.DataSegmentIndex = uint64(dataSegIndex)

	return nil
}
func DecodeElemSegmentAndTableImm(rd io.Reader, imm *types.ElemSegmentAndTableImm, funcDef *types.FunctionDef) error {
	var (
		elemSegIndex uint32
		tableIndex   uint32
	)
	_, err := utils.DecodeVarInt(rd, 32, &elemSegIndex)
	if err != nil {
		return err
	}
	_, err = utils.DecodeVarInt(rd, 32, &tableIndex)
	if err != nil {
		return err
	}

	imm.ElemSegmentIndex = uint64(elemSegIndex)
	imm.TableIndex = uint64(tableIndex)
	return nil
}
func DecodeElemSegmentImm(rd io.Reader, imm *types.ElemSegmentImm, funcDef *types.FunctionDef) error {
	var elemSegIndex uint32
	_, err := utils.DecodeVarInt(rd, 32, &elemSegIndex)
	if err != nil {
		return err
	}
	imm.ElemSegmentIndex = uint64(elemSegIndex)
	return nil
}
func DecodeExceptionTypeImm(rd io.Reader, imm *types.ExceptionTypeImm, funcDef *types.FunctionDef) error {
	var typeIndex uint32
	_, err := utils.DecodeVarInt(rd, 32, &typeIndex)
	if err != nil {
		return err
	}
	imm.ExceptionTypeIndex = uint64(typeIndex)
	return nil
}
func DecodeFunctionImm(rd io.Reader, imm *types.FunctionImm, funcDef *types.FunctionDef) error {
	var functionIndex uint32
	_, err := utils.DecodeVarInt(rd, 32, &functionIndex)
	if err != nil {
		return err
	}
	imm.FunctionIndex = uint64(functionIndex)
	return nil
}
func DecodeGetOrSetVariableImm(rd io.Reader, imm *types.GetOrSetVariableImm, funcDef *types.FunctionDef) error {
	var varIndex uint32
	_, err := utils.DecodeVarInt(rd, 32, &varIndex)
	if err != nil {
		return err
	}
	imm.VariableIndex = uint64(varIndex)
	return nil
}

func DecodeLaneIndexImm(rd io.Reader, imm *types.LaneIndexImm, funcDef *types.FunctionDef) error {
	var laneIndex uint8
	_, err := utils.DecodeVarInt(rd, 7, &laneIndex)
	if err != nil {
		return err
	}
	imm.LaneIndex = laneIndex
	return nil
}
func DecodeLoadOrStoreImm(rd io.Reader, imm *types.LoadOrStoreImm, funcDef *types.FunctionDef) error {
	var (
		alignment uint8
		offset    uint32
	)
	_, err := utils.DecodeVarInt(rd, 7, &alignment)
	if err != nil {
		return err
	}
	_, err = utils.DecodeVarInt(rd, 32, &offset)
	if err != nil {
		return err
	}
	imm.AlignmentLog2 = uint32(alignment)
	imm.Offset = offset

	return nil
}

func DecodeMemoryImm(rd io.Reader, imm *types.MemoryImm, funcDef *types.FunctionDef) error {
	err := checkConstant(rd, []byte{0}, "memory.(grow|size|copy|fill) immediate reserved field must be 0")
	if err != nil {
		return err
	}
	imm.MemoryIndex = 0
	return nil
}

func DecodeRethrowImm(rd io.Reader, imm *types.RethrowImm, funcDef *types.FunctionDef) error {
	var catchDepth uint32
	_, err := utils.DecodeVarInt(rd, 32, &catchDepth)
	if err != nil {
		return err
	}
	imm.CatchDepth = uint64(catchDepth)
	return nil
}

func DecodeTableImm(rd io.Reader, imm *types.TableImm, funcDef *types.FunctionDef) error {
	var tableIndex uint32
	_, err := utils.DecodeVarInt(rd, 32, &tableIndex)
	if err != nil {
		return err
	}
	imm.TableIndex = uint64(tableIndex)
	return nil
}

func DecodeShuffleImm_16(rd io.Reader, imm *types.ShuffleImm_16, funcDef *types.FunctionDef) error {
	var index uint8
	for laneIndex := 0; laneIndex < 16; laneIndex++ {
		_, err := utils.DecodeVarInt(rd, 7, &index)
		if err != nil {
			return nil
		}
		imm.LaneIndices[laneIndex] = index
	}
	return nil
}

func DecodeLiteralImm_F32(rd io.Reader, imm *types.LiteralImm_F32, funcDef *types.FunctionDef) error {
	f32Bytes, err := utils.ReadNByte(rd, 4)
	if err != nil {
		return err
	}
	f32, err := utils.Bytes2float32(f32Bytes, true)
	if err != nil {
		return err
	}
	imm.Value = f32
	return nil
}

func DecodeLiteralImm_F64(rd io.Reader, imm *types.LiteralImm_F64, funcDef *types.FunctionDef) error {
	f64Bytes, err := utils.ReadNByte(rd, 8)
	if err != nil {
		return err
	}

	f64, err := utils.Bytes2float64(f64Bytes, true)
	if err != nil {
		return err
	}
	imm.Value = f64
	return nil
}

func DecodeLiteralImm_I32(rd io.Reader, imm *types.LiteralImm_I32, funcDef *types.FunctionDef) error {
	var v int32
	_, err := utils.DecodeVarInt(rd, 32, &v)
	if err != nil {
		return err
	}
	imm.Value = v
	return nil
}

func DecodeLiteralImm_I64(rd io.Reader, imm *types.LiteralImm_I64, funcDef *types.FunctionDef) error {
	var v int64
	_, err := utils.DecodeVarInt(rd, 64, &v)
	if err != nil {
		return err
	}
	imm.Value = v

	return nil
}

func DecodeLiteralImm_V128(rd io.Reader, imm *types.LiteralImm_V128, funcDef *types.FunctionDef) error {
	v128Bytes, err := utils.ReadNByte(rd, 16)
	if err != nil {
		return err
	}
	copy(imm.Value[:], v128Bytes)

	return nil
}
