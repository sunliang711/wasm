package types

type Module struct {
	FeatureSpec FeatureSpec

	Types []FunctionType

	Functions      IndexSpaceFunction
	Tables         IndexSpaceTable
	Memories       IndexSpaceMemory
	Globals        IndexSpaceGlobal
	ExceptionTypes IndexSpaceExceptionType

	Exports      []Export
	DataSegments []DataSegment
	ElemSegments []ElemSegment
	UserSections []UserSection

	StartFunctionIndex int
}

func (m Module) String() string {
	return "TODO Module::String()"
}
