package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// StackWithMin can push, pop, and get min in O(1)
type StackWithMin struct {
	data datastructure.DynamicStack
	min  datastructure.DynamicStack
}

// Push an element into stack
func (s *StackWithMin) Push(e datastructure.Elem) {
	if e == nil {
		return
	}

	s.data.Push(e)

	if s.min.Len() == 0 || datastructure.Less(e, *s.min.GetTop()) {
		s.min.Push(e)
	} else {
		s.min.Push(*s.min.GetTop())
	}
}

// Pop an el(ement from stack
func (s *StackWithMin) Pop() *datastructure.Elem {
	s.min.Pop()
	return s.data.Pop()
}

// Min of stack
func (s *StackWithMin) Min() *datastructure.Elem {
	return s.min.GetTop()
}

// Len of stack
func (s *StackWithMin) Len() int {
	return s.Len()
}

// Top of stack
func (s *StackWithMin) Top() *datastructure.Elem {
	return s.data.GetTop()
}
