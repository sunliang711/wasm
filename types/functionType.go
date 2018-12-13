package types

import "fmt"

type FunctionType struct {
	//TODO inner struct 'Encoding' Used to represent a function type as an abstract pointer-sized value in the runtime.
	Results *TypeTuple
	Params  *TypeTuple
}

func (ft FunctionType) String() string{
	return fmt.Sprintf("{Params: %v , Results: %v}",ft.Params,ft.Results)
}