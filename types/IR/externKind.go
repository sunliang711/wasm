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
