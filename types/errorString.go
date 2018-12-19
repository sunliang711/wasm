package types

const (
	ErrReadCount      = "Read count error"
	ErrIncorrectOrder = "Incorrect order for known section"

	ErrMagicNumber = "Magic number error"
	ErrVersion     = "Version number error"

	ErrInsufficientBytes = "Got insufficient bytes"

	ErrUnknownSection = "Error unknown section with raw id: %d"

	ErrSectionType = `Not a "%s"`
	ErrSectionNum  = `Num of "%s" is invalid`

	ErrFunctionTag = "Not a function tag"

	ErrFunctionTypeIndexOutOfRange = "Function type index out of range"

	ErrReferenceTypeByte = "Invalid reference type byte"

	ErrNotPtr = "Parameter isn't pointer"
	ErrIntPtr = "Parameter must pointer of (u)int8 or (u)int16 or (u)int32 or (u)int64"

	ErrInvalidElemFlags    = "Invalid elem flags"
	ErrInvalidDataSegFlags = "Invalid data segment flags"

	ErrInvalidInitializerExpressionOpcode = "Invalid initializer expression opcode"

	ErrInvalidFloat32Format = "Invalid float32 format"
	ErrInvalidFloat64Format = "Invalid float64 format"

	ErrInvalidParameter = "Invalid parameter"
)
