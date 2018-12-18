package types

import (
	"fmt"
)

type RefType byte

func (r RefType) String() string {
	switch r {
	case RTInvalid:
		return fmt.Sprintf("{ReferenceType: Invalid(%d)}", byte(r))
	case RTAnyRef:
		return fmt.Sprintf("{ReferenceType: AnyRef(%d)}", byte(r))
	case RTAnyFunc:
		return fmt.Sprintf("{ReferenceType: AnyFunc(%d)}", byte(r))
	}
	return fmt.Sprintf("{ReferenceType: Unknown(%d)}", byte(r))
}

const (
	RTInvalid RefType = 0
	RTAnyRef  RefType = 7
	RTAnyFunc RefType = 8
)

type SizeConstraints struct {
	Min uint64
	Max uint64
}

func (s SizeConstraints) String() string {
	return fmt.Sprintf("{Min: %d,Max: %d}", s.Min, s.Max)
}

type FunctionType struct {
	//TODO inner struct 'Encoding' Used to represent a function type as an abstract pointer-sized value in the runtime.
	Results *TypeTuple
	Params  *TypeTuple
}

func (ft FunctionType) String() string {
	return fmt.Sprintf("{Params: %v , Results: %v}", ft.Params, ft.Results)
}

//basicType: IndexedFunctionType,TableType,MemoryType,GlobalType,ExceptionType
//basicType -> import
//basicType -> basicDef
//import,basicDef -> IndexSpace

//{{basicType BEGIN
type IndexedFunctionType struct {
	Index uint64
}

func (i IndexedFunctionType) String() string {
	return fmt.Sprintf("{Index: %d}", i.Index)
}

type TableType struct {
	ElementType RefType
	IsShared    bool
	Size        SizeConstraints
}

func (t TableType) String() string {
	return fmt.Sprintf("{ElementType: %v,IsShared: %v,Size: %v}", t.ElementType, t.IsShared, t.Size)
}

type MemoryType struct {
	IsShared bool
	Size     SizeConstraints
}

func (m MemoryType) String() string {
	return fmt.Sprintf("{IsShared: %v,Size: %v}", m.IsShared, m.Size)
}

type GlobalType struct {
	IsMutable bool
	ValType   ValueType
}

func (g GlobalType) String() string {
	return fmt.Sprintf("{IsMutable:%v,ValType:%v}", g.IsMutable, g.ValType)
}

type ExceptionType struct {
	Params TypeTuple
}

func (e ExceptionType) String() string {
	return fmt.Sprintf("{Params: %v}", e.Params)
}

//}}basicType END

type ImportCommon struct {
	ModuleName string
	ExportName string
}

func (i ImportCommon) String() string {
	return fmt.Sprintf("{ModuleName: %s,ExportName: %s}", i.ModuleName, i.ExportName)
}

//{{import BEGIN
type ImportIndexedFunctionType struct {
	Type IndexedFunctionType
	ImportCommon
}

func (i ImportIndexedFunctionType) String() string {
	return fmt.Sprintf("{Type: %v,ImportCommon: %v}", i.Type, i.ImportCommon)
}

type ImportTableType struct {
	Type TableType
	ImportCommon
}

func (i ImportTableType) String() string {
	return fmt.Sprintf("{Type: %v,ImportCommon: %v}", i.Type, i.ImportCommon)
}

type ImportMemoryType struct {
	Type MemoryType
	ImportCommon
}

func (i ImportMemoryType) String() string {
	return fmt.Sprintf("{Type: %v,Import Common: %v}", i.Type, i.ImportCommon)
}

type ImportGlobalType struct {
	Type GlobalType
	ImportCommon
}

func (i ImportGlobalType) String() string {
	return fmt.Sprintf("{Type: %v,ImportCommon: %v}", i.Type, i.ImportCommon)
}

type ImportExceptionType struct {
	Type ExceptionType
	ImportCommon
}

func (i ImportExceptionType) String() string {
	return fmt.Sprintf("{Type: %v,ImportCommon: %v}", i.Type, i.ImportCommon)
}

//}}import END

//{{basicDef BEGIN
type FunctionDef struct {
	Type                   IndexedFunctionType
	NonParameterLocalTypes []ValueType
	Code                   []byte
	BranchTables           [][]uint64
}

type TableDef struct {
	Type TableType
}

type MemoryDef struct {
	Type MemoryType
}

type GlobalDef struct {
	Type GlobalType
	//TODO: InitializerExpression initializer;
}

type ExceptionTypeDef struct {
	Type ExceptionType
}

//}}basicDef END

//{{IndexSpace BEGIN
type IndexSpaceFunction struct {
	Imports []ImportIndexedFunctionType
	Defs    []FunctionDef
}
type IndexSpaceTable struct {
	Imports []ImportTableType
	Defs    []TableDef
}
type IndexSpaceMemory struct {
	Imports []ImportMemoryType
	Defs    []MemoryDef
}
type IndexSpaceGlobal struct {
	Imports []ImportGlobalType
	Defs    []GlobalDef
}
type IndexSpaceExceptionType struct {
	Imports []ImportExceptionType
	Defs    []ExceptionTypeDef
}

//}}IndexSpace END

type Export struct {
	Name  string
	Kind  ExternKind
	Index uint64
}

type DataSegment struct {
	IsActive    bool
	MemoryIndex uint64
	//TODO InitializerExpression baseOffset;
	Data []byte
}

type ElemSegment struct {
	IsActive   bool
	TableIndex uint64
	//TODO InitializerExpression baseOffset;
	Indices []uint64
}

type UserSection struct {
	Name string
	Data []byte
}

type FeatureSpec struct {
	//TODO // go : field of struct is false
	//// A feature flag for the MVP, just so the MVP operators can reference it as the required
	//// feature flag.
	//bool mvp = true;
	//
	//// Proposed standard extensions
	//bool importExportMutableGlobals = true;
	//bool extendedNamesSection = true;
	//bool simd = true;
	//bool atomics = true;
	//bool exceptionHandling = true;
	//bool nonTrappingFloatToInt = true;
	//bool extendedSignExtension = true;
	//bool multipleResultsAndBlockParams = true;
	//bool bulkMemoryOperations = true;
	//bool referenceTypes = true;
	//bool quotedNamesInTextFormat = true; // Enabled by default for everything but wavm-disas,
	//// where a command-line flag is required to enable it
	//// to ensure the default output uses standard syntax.
	//
	//// WAVM-specific extensions
	//bool sharedTables = true;
	//bool functionRefInstruction = true;
	//bool requireSharedFlagForAtomicOperators = false; // (true is standard)
	//
	//Uptr maxLocals = 65536;
	//Uptr maxLabelsPerFunction = UINTPTR_MAX;
}
