package main

import (
	"errors"
)

type Stack struct {
	items []interface{}
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is Empty")
	} else {
		lastElemIndex := len(s.items) - 1
		flight := s.items[lastElemIndex]
		s.items = s.items[:lastElemIndex]
		return flight, nil
	}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is Empty")
	} else {
		lastElemIndex := len(s.items) - 1
		return s.items[lastElemIndex], nil
	}
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}
