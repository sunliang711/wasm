package types

type OpcodeAndImm_AtomicLoadOrStoreImm struct {
	Opcode
	Imm AtomicLoadOrStoreImm
}
type OpcodeAndImm_BranchImm struct {
	Opcode
	Imm BranchImm
}
type OpcodeAndImm_BranchTableImm struct {
	Opcode
	Imm BranchTableImm
}
type OpcodeAndImm_CallIndirectImm struct {
	Opcode
	Imm CallIndirectImm
}
type OpcodeAndImm_ControlStructureImm struct {
	Opcode
	Imm ControlStructureImm
}
type OpcodeAndImm_DataSegmentAndMemImm struct {
	Opcode
	Imm DataSegmentAndMemImm
}
type OpcodeAndImm_DataSegmentImm struct {
	Opcode
	Imm DataSegmentImm
}
type OpcodeAndImm_ElemSegmentAndTableImm struct {
	Opcode
	Imm ElemSegmentAndTableImm
}
type OpcodeAndImm_ElemSegmentImm struct {
	Opcode
	Imm ElemSegmentImm
}
type OpcodeAndImm_ExceptionTypeImm struct {
	Opcode
	Imm ExceptionTypeImm
}
type OpcodeAndImm_FunctionImm struct {
	Opcode
	Imm FunctionImm
}
type OpcodeAndImm_GetOrSetVariableImm struct {
	Opcode
	Imm GetOrSetVariableImm
}
type OpcodeAndImm_LaneIndexImm struct {
	Opcode
	Imm LaneIndexImm
}
type OpcodeAndImm_LiteralImm_F32 struct {
	Opcode
	Imm LiteralImm_F32
}
type OpcodeAndImm_LiteralImm_F64 struct {
	Opcode
	Imm LiteralImm_F64
}
type OpcodeAndImm_LiteralImm_I32 struct {
	Opcode
	Imm LiteralImm_I32
}
type OpcodeAndImm_LiteralImm_I64 struct {
	Opcode
	Imm LiteralImm_I64
}
type OpcodeAndImm_LiteralImm_V128 struct {
	Opcode
	Imm LiteralImm_V128
}
type OpcodeAndImm_LoadOrStoreImm struct {
	Opcode
	Imm LoadOrStoreImm
}
type OpcodeAndImm_MemoryImm struct {
	Opcode
	Imm MemoryImm
}
type OpcodeAndImm_NoImm struct {
	Opcode
	Imm NoImm
}
type OpcodeAndImm_RethrowImm struct {
	Opcode
	Imm RethrowImm
}
type OpcodeAndImm_ShuffleImm_16 struct {
	Opcode
	Imm ShuffleImm_16
}
type OpcodeAndImm_TableImm struct {
	Opcode
	Imm TableImm
}
