package types

import "fmt"

type TypeTuple struct {
	NumElems uint32
	Elems    []ValueType
}

func (t TypeTuple) String() string {
	return fmt.Sprintf("{NumElems: %d ,Elems: %v}", t.NumElems, t.Elems)
}
