package IR

import "fmt"

type ImmKind byte
type Imm interface {
	ImmKind() ImmKind
	fmt.Stringer
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
func (imm *AtomicLoadOrStoreImm) String() string {
	return fmt.Sprintf("{alignment: %v,offset: %v}", imm.AlignmengLog2, imm.Offset)
}

type BranchImm struct {
	TargetDepth uint64
}

func (imm *BranchImm) ImmKind() ImmKind {
	return KindBranchImm
}
func (imm *BranchImm) String() string {
	return fmt.Sprintf("{targetDepth: %v}", imm.TargetDepth)
}

type BranchTableImm struct {
	DefaultTargetDepth uint64
	BranchTableIndex   uint64
}

func (imm *BranchTableImm) ImmKind() ImmKind {
	return KindBranchTableImm
}
func (imm *BranchTableImm) String() string {
	return fmt.Sprintf("{defaultTargetDepth: %v,branchTableIndex: %v}", imm.DefaultTargetDepth, imm.BranchTableIndex)
}

type CallIndirectImm struct {
	Type       IndexedFunctionType
	TableIndex uint64
}

func (imm *CallIndirectImm) ImmKind() ImmKind {
	return KindCallIndirectImm
}
func (imm *CallIndirectImm) String() string {
	return fmt.Sprintf("{type: %v,tableIndex: %v}", imm.Type, imm.TableIndex)
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
func (imm *ControlStructureImm) String() string {
	return fmt.Sprintf("{type: %v}", imm.Type)
}

type DataSegmentAndMemImm struct {
	DataSegmentIndex uint64
	MemoryIndex      uint64
}

func (imm *DataSegmentAndMemImm) ImmKind() ImmKind {
	return KindDataSegmentAndMemImm
}
func (imm *DataSegmentAndMemImm) String() string {
	return fmt.Sprintf("{dataSegmentIndex: %v,memoryIndex: %v}", imm.DataSegmentIndex, imm.MemoryIndex)
}

type DataSegmentImm struct {
	DataSegmentIndex uint64
}

func (imm *DataSegmentImm) ImmKind() ImmKind {
	return KindDataSegmentImm
}

func (imm *DataSegmentImm) String() string {
	return fmt.Sprintf("{dataSegmentIndex: %v}", imm.DataSegmentIndex)
}

type ElemSegmentAndTableImm struct {
	ElemSegmentIndex uint64
	TableIndex       uint64
}

func (imm *ElemSegmentAndTableImm) ImmKind() ImmKind {
	return KindElemSegmentAndTableImm
}
func (imm *ElemSegmentAndTableImm) String() string {
	return fmt.Sprintf("{elementSegmentIndex: %v,tableIndex: %v}")
}

type ElemSegmentImm struct {
	ElemSegmentIndex uint64
}

func (imm *ElemSegmentImm) ImmKind() ImmKind {
	return KindElemSegmentImm
}
func (imm *ElemSegmentImm) String() string {
	return fmt.Sprintf("{elemSegmentIndex: %v}", imm.ElemSegmentIndex)
}

type ExceptionTypeImm struct {
	ExceptionTypeIndex uint64
}

func (imm *ExceptionTypeImm) ImmKind() ImmKind {
	return KindExceptionTypeImm
}
func (imm *ExceptionTypeImm) String() string {
	return fmt.Sprintf("{exceptionTypeIndex: %v}", imm.ExceptionTypeIndex)
}

type FunctionImm struct {
	FunctionIndex uint64
}

func (imm *FunctionImm) ImmKind() ImmKind {
	return KindFunctionImm
}
func (imm *FunctionImm) String() string {
	return fmt.Sprintf("{functionIndex: %v}", imm.FunctionIndex)
}

type GetOrSetVariableImm struct {
	VariableIndex uint64
}

func (imm *GetOrSetVariableImm) ImmKind() ImmKind {
	return KindGetOrSetVariableImm
}
func (imm *GetOrSetVariableImm) String() string {
	return fmt.Sprintf("{variableIndex: %v}", imm.VariableIndex)
}

type LaneIndexImm struct {
	LaneIndex byte
}

func (imm *LaneIndexImm) ImmKind() ImmKind {
	return KindLaneIndexImm
}
func (imm *LaneIndexImm) String() string {
	return fmt.Sprintf("{laneIndex: %v}", imm.LaneIndex)
}

type LiteralImm_F32 struct {
	Value float32
}

func (imm *LiteralImm_F32) ImmKind() ImmKind {
	return KindLiteralImm_F32
}
func (imm *LiteralImm_F32) String() string {
	return fmt.Sprintf("{value: %v}", imm.Value)
}

type LiteralImm_F64 struct {
	Value float64
}

func (imm *LiteralImm_F64) ImmKind() ImmKind {
	return KindLiteralImm_F64
}
func (imm *LiteralImm_F64) String() string {
	return fmt.Sprintf("{value: %v}", imm.Value)
}

type LiteralImm_I32 struct {
	Value int32
}

func (imm *LiteralImm_I32) ImmKind() ImmKind {
	return KindLiteralImm_I32
}
func (imm *LiteralImm_I32) String() string {
	return fmt.Sprintf("{value: %v}", imm.Value)
}

type LiteralImm_I64 struct {
	Value int64
}

func (imm *LiteralImm_I64) ImmKind() ImmKind {
	return KindLiteralImm_I64
}
func (imm *LiteralImm_I64) String() string {
	return fmt.Sprintf("{value: %v}", imm.Value)
}

type LiteralImm_V128 struct {
	Value [16]byte
}

func (imm *LiteralImm_V128) ImmKind() ImmKind {
	return KindLiteralImm_V128
}
func (imm *LiteralImm_V128) String() string {
	return fmt.Sprintf("{value: %v}", imm.Value)
}

type LoadOrStoreImm struct {
	AlignmentLog2 byte
	//AlignmentLog2 uint32
	Offset uint32
}

func (imm *LoadOrStoreImm) ImmKind() ImmKind {
	return KindLoadOrStoreImm
}
func (imm *LoadOrStoreImm) String() string {
	return fmt.Sprintf("{alignment: %v,offset: %v}", imm.AlignmentLog2, imm.Offset)
}

type MemoryImm struct {
	MemoryIndex uint64
}

func (imm *MemoryImm) ImmKind() ImmKind {
	return KindMemoryImm
}
func (imm *MemoryImm) String() string {
	return fmt.Sprintf("{memoryIndex: %v}", imm.MemoryIndex)
}

type NoImm struct {
}

func (imm *NoImm) ImmKind() ImmKind {
	return KindNoImm
}
func (imm *NoImm) String() string {
	return "{noimm}"
}

type RethrowImm struct {
	CatchDepth uint64
}

func (imm *RethrowImm) ImmKind() ImmKind {
	return KindRethrowImm
}
func (imm *RethrowImm) String() string {
	return fmt.Sprintf("{catchDepth: %v}", imm.CatchDepth)
}

type ShuffleImm_16 struct {
	LaneIndices [16]byte
}

func (imm *ShuffleImm_16) ImmKind() ImmKind {
	return KindShuffleImm_16
}
func (imm *ShuffleImm_16) String() string {
	return fmt.Sprintf("{laneIndices: %v}", imm.LaneIndices)
}

type TableImm struct {
	TableIndex uint64
}

func (imm *TableImm) ImmKind() ImmKind {
	return KindTableImm
}
func (imm *TableImm) String() string {
	return fmt.Sprintf("{tableIndex: %v}", imm.TableIndex)
}
