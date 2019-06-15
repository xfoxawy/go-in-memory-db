package stack

import "github.com/go-in-memory-db/linkedlist"

// Stack struct
type Stack struct {
	stack *linkedlist.LinkedList
}

// NewStack generator
func NewStack() *Stack {
	return &Stack{linkedlist.NewList()}
}

// Push element ro stack
func (s *Stack) Push(e string) {
	s.stack.Shift(e)
	return
}

// Pop from stack and get element
func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	}
	top, _ := s.stack.Unshift()
	if top.Value == "" {
		return ""
	}
	return top.Value
}

// Size stack size
func (s *Stack) Size() int {
	return s.stack.Length
}

// IsEmpty check
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
