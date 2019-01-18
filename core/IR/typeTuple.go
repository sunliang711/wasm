package IR

type TypeTuple struct {
	NumElems uint32
	Elems    []ValueType
}

func (t TypeTuple) String() string {
	ret := ""
	for i, e := range t.Elems {
		ret += e.String()
		if i != len(t.Elems)-1 {
			ret += ","
		}
	}
	return ret
}
