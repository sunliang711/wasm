package IR

import (
	"fmt"
)

type InterfaceValue interface {
	Type() ValueType
	Value() interface{}
}

type Stack struct {
	data []InterfaceValue
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(v InterfaceValue) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (InterfaceValue, error) {
	if s.Empty() {
		return nil, fmt.Errorf("Empty Stack")
	}
	ret := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return ret, nil
}
