package datastructure

import (
	"bytes"
	"fmt"
	"sync"
)

// Stack struct
type Stack struct {
	data []Elem
	top  int
	lock sync.RWMutex
}

// NewStack new a stack
func NewStack(n int) (*Stack, error) {
	if n < 1 {
		return nil, fmt.Errorf("stack size must be larger than 0")
	}
	s := &Stack{
		data: make([]Elem, 0, n),
		top:  -1,
	}

	return s, nil
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
	return s.top == -1
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
		_, _ = s.Pop()
	}
	return nil
}

// DynamicStack without limited size
type DynamicStack struct {
	data []Elem
	lock sync.RWMutex
}

// NewDynamicStack new an infinite stack
func NewDynamicStack() *DynamicStack {
	return &DynamicStack{}
}

// Push an element into stack
func (s *DynamicStack) Push(e Elem) {
	s.lock.Lock()
	s.data = append(s.data, e)
	s.lock.Unlock()
}

// Pop an element from stack
func (s *DynamicStack) Pop() Elem {
	if s.IsEmpty() {
		return nil
	}
	s.lock.Lock()
	n := len(s.data)
	e := s.data[n-1]
	s.data = s.data[:n-1]
	defer s.lock.Unlock()
	return e
}

// IsEmpty returns true if the stack is empty
func (s *DynamicStack) IsEmpty() bool {
	return len(s.data) == 0
}

// GetTop of the stack
func (s *DynamicStack) GetTop() Elem {
	if s.IsEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

// Len returns the number of elements in the stack
func (s *DynamicStack) Len() int {
	return len(s.data)
}

// String returns string that are composed of stack elements from bottom to top
func (s *DynamicStack) String() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	var b bytes.Buffer
	for _, e := range s.data {
		b.WriteString(fmt.Sprintf("%v ", e))
	}
	return b.String()
}
