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

func (m *Module) GetAllFuncIns() string {
	ret := ""
	for _, f := range m.Functions.Defs {
		ret += "function: " + f.Name + "{"
		for i, ins := range f.Instruction {
			ret += fmt.Sprintf("[%d]: %s", i, ins.Op.Name)
			if i != len(f.Instruction)-1 {
				ret += ","
			}
			if (i+1)%8 == 0 {
				ret += "\n"
			}
		}
		ret += "}\n"
	}
	return ret
}

func (m *Module) GetAllFunc() string {
	ret := ""
	for i, f := range m.Functions.Defs {
		ret += "function: " + f.Name
		if i != len(m.Functions.Defs)-1 {
			ret += ","
		}
		if (i+1)%8 == 0 {
			ret += "\n"
		}
	}

	return ret
}
