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
)
