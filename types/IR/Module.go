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
