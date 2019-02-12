package IR

import (
	"fmt"
	"wasm/types"
)

type Module struct {
	FeatureSpec *FeatureSpec

	Types []FunctionType

	Functions      IndexSpaceFunction
	Tables         IndexSpaceTable
	Memories       IndexSpaceMemory
	Globals        IndexSpaceGlobal
	ExceptionTypes IndexSpaceExceptionType

	Exports         []*Export
	ExportFunctions []*Export
	DataSegments    []DataSegment
	ElemSegments    []ElemSegment
	UserSections    []UserSection

	StartFunctionIndex int

	//TODO: Add all export function names array
}

func (m Module) String() string {
	//TODO
	return "TODO Module::String()"
}

func NewModule() *Module {
	return &Module{
		FeatureSpec: NewFeatureSepc(),
	}
}

func (m *Module) GetFuncIndexWithName(name string) (int, error) {
	for _, f := range m.ExportFunctions {
		if f.Name == name {
			return int(f.Index), nil
		}
	}
	return -1, fmt.Errorf(types.ErrFuncNotFound)
}

func (m *Module) GetAllFuncIns(detail bool) string {
	ret := ""
	for _, f := range m.Functions.Defs {
		ret += fmt.Sprintf("Function: %s { \n", f.Name)
		for i, ins := range f.Instruction {
			if detail {
				ret += fmt.Sprintf("[%06d]:%-12s %s\n", i, ins.Op.Name, ins.Imm.String())
			} else {
				ret += fmt.Sprintf("[%06d]:%-12s\n", i, ins.Op.Name)
			}
			//if i != len(f.Instruction)-1 {
			//	ret += ", "
			//}
			//if (i+1)%8 == 0 && i != len(f.Instruction)-1 {
			//	ret += "\n"
			//}
		}
		ret += "}\n"
	}
	return ret
}

func (m *Module) GetAllFunc() string {
	ret := ""
	for i, f := range m.Functions.Defs {
		ret += fmt.Sprintf("function: %s %s", f.Name, f.FunctionType)
		if i != len(m.Functions.Defs)-1 {
			ret += "\n"
		}
	}
	return ret
}

func (m *Module) GetAllSections() string {
	ret := ""
	//types
	ret += "type section:\n"
	for _, funcType := range m.Types {
		ret += funcType.String() + "\n"
	}

	//functions
	ret += "\nimport functions:\n"
	for _, imFunc := range m.Functions.Imports {
		ret += imFunc.String() + "\n"
	}

	ret += "defined functions:\n"
	for _, defFunc := range m.Functions.Defs {
		ret += defFunc.String() + "\n"
	}

	//tables
	ret += "\nimport tables:\n"
	for _, imTab := range m.Tables.Imports {
		ret += imTab.String() + "\n"
	}

	ret += "defined tables:\n"
	for _, defTab := range m.Tables.Defs {
		ret += defTab.String() + "\n"
	}

	//memories
	ret += "\nimport memories:\n"
	for _, imMem := range m.Memories.Imports {
		ret += imMem.String() + "\n"
	}
	ret += "defined memories:\n"
	for _, defMem := range m.Memories.Defs {
		ret += defMem.String() + "\n"
	}

	//globals
	ret += "\nimport globals:\n"
	for _, imGlobal := range m.Globals.Imports {
		ret += imGlobal.String() + "\n"
	}
	for _, defGlobal := range m.Globals.Defs {
		ret += defGlobal.String() + "\n"
	}

	//exports
	ret += "\nexport section:\n"
	for _, ex := range m.Exports {
		ret += ex.String() + "\n"
	}
	return ret
}
