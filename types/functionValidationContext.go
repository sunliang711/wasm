package types

type ControlContextType byte

const (
	CCTFunction ControlContextType = iota
	CCTBlock
	CCTIfThen
	CCTIfElse
	CCTLoop
	CCTTry
	CCTCatch
)

type ControlContext struct {
	Type           ControlContextType
	OuterStackSize uint64

	Params      TypeTuple
	Results     TypeTuple
	IsReachable bool

	elseParams TypeTuple
}

type FunctionValidationContext struct {
	*Module
	*FunctionDef
	*DeferredCodeValidationState
	*FunctionType

	Locals       []ValueType
	ControlStack []ControlContext
	Stack        []ValueType
}
