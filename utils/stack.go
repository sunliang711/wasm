package utils

import "fmt"

type Stack struct {
	data []interface{}
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, fmt.Errorf("Empty Stack")
	}
	ret := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return ret, nil
}

func (s *Stack) Top() (interface{}, error) {
	if s.Empty() {
		return nil, fmt.Errorf("Empty Stack")
	}
	return s.data[len(s.data)-1], nil
}
