package datastructure

import (
	"fmt"
	"sync"
)

// Stack struct
type Stack struct {
	data []Elem
	top  int
	lock sync.RWMutex
}

// New a stack
func (s *Stack) New(n int) error {
	if n < 1 {
		return fmt.Errorf("stack size must be larger than 0")
	}
	s.data = make([]Elem, 0, n)
	s.top = -1
	return nil
}

// Push an element into stack
func (s *Stack) Push(e Elem) error {
	if s.top == cap(s.data)-1 {
		return fmt.Errorf("stack is full")
	}
	n := len(s.data)
	s.lock.Lock()
	s.data = s.data[:n+1]
	s.data[n] = e
	s.top++
	s.lock.Unlock()
	return nil
}

// Pop an element from stack
func (s *Stack) Pop() (Elem, error) {
	if s.top == -1 {
		return nil, fmt.Errorf("stack is empty")
	}
	n := len(s.data)
	s.lock.Lock()
	e := s.data[n-1]
	s.data = s.data[:n-1]
	s.top--
	s.lock.Unlock()

	return e, nil
}

// IsEmpty stack
func (s *Stack) IsEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

// GetTop element of stack
func (s *Stack) GetTop() Elem {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.data[s.top]
}

// Len of stack: element number
func (s *Stack) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.data)
}

// Clear stack
func (s *Stack) Clear() error {
	if s.top == -1 {
		return fmt.Errorf("stack is already empty")
	}

	for !s.IsEmpty() {
		s.Pop()
	}
	return nil
}

// InfStack without limited size
type InfStack struct {
	data []Elem
	lock sync.RWMutex
}

// New an infinite stack
func (s *InfStack) New() *InfStack {
	s.data = []Elem{}
	return s
}

// Push an element into stack
func (s *InfStack) Push(e Elem) {
	s.lock.Lock()
	s.data = append(s.data, e)
	s.lock.Unlock()
}

// Pop an element from stack
func (s *InfStack) Pop() *Elem {
	s.lock.Lock()
	n := len(s.data)
	e := s.data[n-1]
	s.data = s.data[:n-1]
	defer s.lock.Unlock()
	return &e
}

// IsEmpty returns true if the stack is empty
func (s *InfStack) IsEmpty() bool {
	return len(s.data) == 0
}

// GetTop of the stack
func (s *InfStack) GetTop() *Elem {
	if s.IsEmpty() {
		return nil
	}
	return &s.data[len(s.data)-1]
}

// Len returns the number of elements in the stack
func (s *InfStack) Len() int {
	return len(s.data)
}
