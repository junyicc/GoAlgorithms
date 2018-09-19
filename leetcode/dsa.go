package leetcode

import (
	"bytes"
	"fmt"
	"sync"
)

// ListNode singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Stack without limited size
type Stack struct {
	data []interface{}
	lock sync.RWMutex
}

// New an infinite stack
func (s *Stack) New() *Stack {
	s.data = []interface{}{}
	return s
}

// Push an element into stack
func (s *Stack) Push(e interface{}) {
	s.lock.Lock()
	s.data = append(s.data, e)
	s.lock.Unlock()
}

// Pop an element from stack
func (s *Stack) Pop() interface{} {
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
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Top of the stack
func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

// Len returns the number of elements in the stack
func (s *Stack) Len() int {
	return len(s.data)
}

// String returns string that are composed of stack elements from bottom to top
func (s *Stack) String() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	var b bytes.Buffer
	for _, e := range s.data {
		b.WriteString(fmt.Sprintf("%v ", e))
	}
	return b.String()
}
