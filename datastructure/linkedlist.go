package datastructure

import (
	"bytes"
	"fmt"
	"sync"
)

// Node of linkedlist
type Node struct {
	Data Elem
	Next *Node
}

// LinkedList struct
type LinkedList struct {
	head   *Node
	length int
	lock   sync.RWMutex
}

// Append an element to the end of the linked list
func (l *LinkedList) Append(e Elem) {
	l.lock.Lock()
	defer l.lock.Unlock()
	node := Node{e, nil}
	if l.head == nil {
		l.head = &node
	} else {
		last := l.head
		for last.Next != nil {
			last = last.Next
		}
		last.Next = &node
	}
	l.length++
}

// Insert an element at k index
func (l *LinkedList) Insert(k int, e Elem) error {
	if k < 0 || k > l.length {
		return fmt.Errorf("index out of range")
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	node := Node{e, nil}
	if k == 0 {
		node.Next = l.head
		l.head = &node
		return nil
	}
	last := l.head
	for i := 0; i < k-1; i++ {
		last = last.Next
	}
	node.Next = last.Next
	last.Next = &node
	l.length++
	return nil
}

// Remove an element at k index
func (l *LinkedList) Remove(k int) (*Elem, error) {
	if k < 0 || k > l.length {
		return nil, fmt.Errorf("index out of range")
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	last := l.head
	for i := 0; i < k-1; i++ {
		last = last.Next
	}
	node := last.Next
	last.Next = node.Next
	l.length--
	return &node.Data, nil
}

// Clear linked list
func (l *LinkedList) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()

	last := l.head
	for last.Next != nil {
		last.Next = last.Next.Next
	}
	l.head = nil
	l.length = 0
}

// Get the element at k index
func (l *LinkedList) Get(k int) (*Elem, bool) {
	if k < 0 || k > l.length {
		return nil, false
	}

	l.lock.RLock()
	defer l.lock.RUnlock()

	last := l.head
	for i := 0; i < k && last != nil; i++ {
		last = last.Next
	}
	if last == nil {
		return nil, false
	}
	return &last.Data, true
}

// IndexOf an element
func (l *LinkedList) IndexOf(e Elem) (int, bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()
	last := l.head
	for i := 0; last != nil; i++ {
		if last.Data == e {
			return i, true
		}
		last = last.Next
	}
	return -1, false
}

// IsEmpty the linked list
func (l *LinkedList) IsEmpty() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	if l.head == nil {
		return true
	}
	return false
}

// Len of the linked list
func (l *LinkedList) Len() int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.length
}

func (l *LinkedList) String() string {
	l.lock.RLock()
	defer l.lock.RUnlock()

	node := l.head
	var b bytes.Buffer
	for node.Next != nil {
		b.WriteString(fmt.Sprintf("%v", node.Data))
		b.WriteString("->")
		node = node.Next
	}
	b.WriteString(node.Data.(string))
	return b.String()
}

// Head of the linked list
func (l *LinkedList) Head() *Node {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.head
}
