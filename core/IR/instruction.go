package IR

type Instruction struct {
	Op    *Op
	Imm   Imm
	Index int
	//for opcode 'end'
	MatchedIndex int
}

func (i *Instruction) Type() ValueType {
	return TypeIns
}

func (i *Instruction) Value() interface{} {
	return i
}
