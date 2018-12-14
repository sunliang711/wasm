package types

type ExternKind byte

const(
	Function ExternKind = iota
	Table
	Memory
	Global
	Exception

	Max = 4
	Invalid = 0xff
)
