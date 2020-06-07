package IR

import (
	"fmt"
	"sync"
)

type InterfaceValue interface {
	Type() ValueType
	Value() interface{}
}

type Stack struct {
	data []InterfaceValue
	mtx  sync.Mutex
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(v InterfaceValue) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (InterfaceValue, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if s.Empty() {
		return nil, fmt.Errorf("Empty Stack")
	}
	ret := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return ret, nil
}

func (s *Stack) Top() (InterfaceValue, error) {
	if s.Empty() {
		return nil, fmt.Errorf("Empty Stack")
	}
	return s.data[len(s.data)-1], nil
}
