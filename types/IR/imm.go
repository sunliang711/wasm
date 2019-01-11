package IR

type ImmKind byte
type Imm interface {
	ImmKind() ImmKind
}

const (
	KindNoImm ImmKind = iota
	KindAtomicLoadOrStoreImm
	KindBranchImm
	KindBranchTableImm
	KindCallIndirectImm
	KindControlStructureImm
	KindDataSegmentAndMemImm
	KindDataSegmentImm
	KindElemSegmentAndTableImm
	KindElemSegmentImm
	KindExceptionTypeImm
	KindFunctionImm
	KindGetOrSetVariableImm
	KindLaneIndexImm
	KindLiteralImm_F32
	KindLiteralImm_F64
	KindLiteralImm_I32
	KindLiteralImm_I64
	KindLiteralImm_V128
	KindLoadOrStoreImm
	KindMemoryImm
	KindRethrowImm
	KindShuffleImm_16
	KindTableImm
)

type AtomicLoadOrStoreImm struct {
	AlignmengLog2 byte
	//AlignmengLog2 uint32
	Offset uint32
}

func (imm *AtomicLoadOrStoreImm) ImmKind() ImmKind {
	return KindAtomicLoadOrStoreImm
}

type BranchImm struct {
	TargetDepth uint64
}

func (imm *BranchImm) ImmKind() ImmKind {
	return KindBranchImm
}

type BranchTableImm struct {
	DefaultTargetDepth uint64
	BranchTableIndex   uint64
}

func (imm *BranchTableImm) ImmKind() ImmKind {
	return KindBranchTableImm
}

type CallIndirectImm struct {
	Type       IndexedFunctionType
	TableIndex uint64
}

func (imm *CallIndirectImm) ImmKind() ImmKind {
	return KindCallIndirectImm
}

type Format byte

const (
	FormatNoParametersOrResult Format = iota
	FormatOneResult
	FormatFunctionType
)

type IndexedBlockedType struct {
	Format
	ResultType ValueType
	//Format     uint32
	//ResultType uint32
	Index uint64
}
type ControlStructureImm struct {
	Type IndexedBlockedType
}

func (imm *ControlStructureImm) ImmKind() ImmKind {
	return KindControlStructureImm
}

type DataSegmentAndMemImm struct {
	DataSegmentIndex uint64
	MemoryIndex      uint64
}

func (imm *DataSegmentAndMemImm) ImmKind() ImmKind {
	return KindDataSegmentAndMemImm
}

type DataSegmentImm struct {
	DataSegmentIndex uint64
}

func (imm *DataSegmentImm) ImmKind() ImmKind {
	return KindDataSegmentImm
}

type ElemSegmentAndTableImm struct {
	ElemSegmentIndex uint64
	TableIndex       uint64
}

func (imm *ElemSegmentAndTableImm) ImmKind() ImmKind {
	return KindElemSegmentAndTableImm
}

type ElemSegmentImm struct {
	ElemSegmentIndex uint64
}

func (imm *ElemSegmentImm) ImmKind() ImmKind {
	return KindElemSegmentImm
}

type ExceptionTypeImm struct {
	ExceptionTypeIndex uint64
}

func (imm *ExceptionTypeImm) ImmKind() ImmKind {
	return KindExceptionTypeImm
}

type FunctionImm struct {
	FunctionIndex uint64
}

func (imm *FunctionImm) ImmKind() ImmKind {
	return KindFunctionImm
}

type GetOrSetVariableImm struct {
	VariableIndex uint64
}

func (imm *GetOrSetVariableImm) ImmKind() ImmKind {
	return KindGetOrSetVariableImm
}

type LaneIndexImm struct {
	LaneIndex byte
}

func (imm *LaneIndexImm) ImmKind() ImmKind {
	return KindLaneIndexImm
}

type LiteralImm_F32 struct {
	Value float32
}

func (imm *LiteralImm_F32) ImmKind() ImmKind {
	return KindLiteralImm_F32
}

type LiteralImm_F64 struct {
	Value float64
}

func (imm *LiteralImm_F64) ImmKind() ImmKind {
	return KindLiteralImm_F64
}

type LiteralImm_I32 struct {
	Value int32
}

func (imm *LiteralImm_I32) ImmKind() ImmKind {
	return KindLiteralImm_I32
}

type LiteralImm_I64 struct {
	Value int64
}

func (imm *LiteralImm_I64) ImmKind() ImmKind {
	return KindLiteralImm_I64
}

type LiteralImm_V128 struct {
	Value [16]byte
}

func (imm *LiteralImm_V128) ImmKind() ImmKind {
	return KindLiteralImm_V128
}

type LoadOrStoreImm struct {
	AlignmentLog2 byte
	//AlignmentLog2 uint32
	Offset uint32
}

func (imm *LoadOrStoreImm) ImmKind() ImmKind {
	return KindLoadOrStoreImm
}

type MemoryImm struct {
	MemoryIndex uint64
}

func (imm *MemoryImm) ImmKind() ImmKind {
	return KindMemoryImm
}

type NoImm struct {
}

func (imm *NoImm) ImmKind() ImmKind {
	return KindNoImm
}

type RethrowImm struct {
	CatchDepth uint64
}

func (imm *RethrowImm) ImmKind() ImmKind {
	return KindRethrowImm
}

type ShuffleImm_16 struct {
	LaneIndices [16]byte
}

func (imm *ShuffleImm_16) ImmKind() ImmKind {
	return KindShuffleImm_16
}

type TableImm struct {
	TableIndex uint64
}

func (imm *TableImm) ImmKind() ImmKind {
	return KindTableImm
}
