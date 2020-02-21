package parser

import "errors"

type stack struct {
	data []*Operation
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) Push(v *Operation) {
	s.data = append(s.data, v)
}

func (s *stack) Pop() (*Operation, error) {
	l := len(s.data)
	if l == 0 {
		return nil, errors.New("empty stack")
	}

	result := s.data[l-1]
	s.data = s.data[:l-1]

	return result, nil
}

func (s *stack) Top() (*Operation, error) {
	l := len(s.data)
	if l == 0 {
		return nil, errors.New("empty stack")
	}

	return s.data[l-1], nil
}
