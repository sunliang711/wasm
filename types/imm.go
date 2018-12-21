package types

type AtomicLoadOrStoreImm struct {
	AlignmengLog2 byte
	Offset        uint32
}
type BranchImm struct {
	TargetDepth uint64
}
type BranchTableImm struct {
	DefaultTargetDepth uint64
	BranchTableIndex   uint64
}
type CallIndirectImm struct {
	Type       IndexedFunctionType
	TableIndex uint64
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
	Index      uint64
}
type ControlStructureImm struct {
	Type IndexedBlockedType
}
type DataSegmentAndMemImm struct {
	DataSegmentIndex uint64
	MemoryIndex      uint64
}
type DataSegmentImm struct {
	DataSegmentIndex uint64
}
type ElemSegmentAndTableImm struct {
	ElemSegmentIndex uint64
	TableIndex       uint64
}
type ElemSegmentImm struct {
	ElemSegmentIndex uint64
}
type ExceptionTypeImm struct {
	ExceptionTypeIndex uint64
}
type FunctionImm struct {
	FunctionIndex uint64
}
type GetOrSetVariableImm struct {
	VariableIndex uint64
}
type LaneIndexImm struct {
	LaneIndex byte
}
type LiteralImm_F32 struct {
	Value float32
}
type LiteralImm_F64 struct {
	Value float64
}
type LiteralImm_I32 struct {
	Value int32
}
type LiteralImm_I64 struct {
	Value int64
}
type LiteralImm_V128 struct {
	Value [16]byte
}
type LoadOrStoreImm struct {
	AlignmentLog2 byte
	Offset        uint32
}
type MemoryImm struct {
	MemoryIndex uint64
}
type NoImm struct {
}
type RethrowImm struct {
	CatchDepth uint64
}
type ShuffleImm_16 struct {
	LaneIndices [16]byte
}
type TableImm struct {
	TableIndex uint64
}
