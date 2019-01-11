package IR

import "fmt"

type ExternKind byte

func (e ExternKind) String() string {
	switch e {
	case Function:
		return fmt.Sprintf("ExternKind:Function")
	case Table:
		return fmt.Sprintf("ExternKind:Table")
	case Memory:
		return fmt.Sprintf("ExternKind:Memory")
	case Global:
		return fmt.Sprintf("ExternKind:Global")
	case Exception:
		return fmt.Sprintf("ExternKind:Exception")
	default:
		return fmt.Sprintf("ExternKind:Unknown")
	}
}

const (
	Function ExternKind = iota
	Table
	Memory
	Global
	Exception

	Max     = 4
	Invalid = 0xff
)

type ExternType struct {
	Kind ExternKind

	Function *FunctionType
	Table    *TableType
	Memory   *MemoryType
	*ExceptionType
}

func ExternTypeF(f *FunctionType) *ExternType {
	return &ExternType{
		Kind:     Function,
		Function: f,
	}
}

func ExternTypeT(t *TableType) *ExternType {
	return &ExternType{
		Kind:  Table,
		Table: t,
	}
}

func ExternTypeM(m *MemoryType) *ExternType {
	return &ExternType{
		Kind:   Memory,
		Memory: m,
	}
}

func ExternTypeE(e *ExceptionType) *ExternType {
	return &ExternType{
		Kind:          Exception,
		ExceptionType: e,
	}
}
