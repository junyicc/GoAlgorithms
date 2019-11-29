package datastructure

import (
	"bytes"
	"fmt"
	"sync"
)

// ListNode of linkedlist
type ListNode struct {
	Data Elem
	Next *ListNode
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{
		Data: v,
		Next: nil,
	}
}

// ComplexListNode of LinkedList
type ComplexListNode struct {
	Data    Elem
	Next    *ComplexListNode
	Sibling *ComplexListNode
}

// Less comparison
func (n *ListNode) Less(node *ListNode) bool {
	if node == nil {
		return false
	}

	switch node.Data.(type) {
	case nil:
		return false
	case int, int8, int16, int32, int64:
		return n.Data.(int) < node.Data.(int)
	case uint, uint8, uint16, uint32, uint64:
		return n.Data.(uint) < node.Data.(uint)
	case float32, float64:
		return n.Data.(float64) < node.Data.(float64)
	case string:
		return n.Data.(string) < node.Data.(string)
	default:
		return false
	}
}

func (n *ListNode) String() string {
	return String(n.Data)
}

//------------------------------------------------------------------------------------------------

// LinkedList struct
type LinkedList struct {
	head   *ListNode
	length int
	lock   sync.RWMutex
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:   NewListNode(0),
		length: 0,
	}
}

// Append an element to the end of the linked list
func (l *LinkedList) Append(e Elem) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	last := l.head
	for last.Next != nil {
		last = last.Next
	}
	return l.InsertAfter(last, e)
}

// Insert an element at k index
func (l *LinkedList) Insert(k int, e Elem) error {
	if k < 0 || k >= l.length {
		return fmt.Errorf("index out of range")
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	var node *ListNode
	if k == 0 {
		node = l.head
	} else {
		last := l.head.Next
		for i := 0; i < k-1; i++ {
			last = last.Next
		}
		node = last
	}
	return l.InsertAfter(node, e)
}

func (l *LinkedList) InsertAfter(p *ListNode, e Elem) error {
	if p == nil {
		return fmt.Errorf("invalid list node")
	}

	n := NewListNode(e)
	n.Next = p.Next
	p.Next = n
	l.length++
	return nil
}

// Remove an element at k index
func (l *LinkedList) Remove(k int) (*Elem, error) {
	if k < 0 || k >= l.length {
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
	// delete node
	data := node.Data
	node.Next = nil
	node = nil
	l.length--
	return &data, nil
}

func (l *LinkedList) DeleteNode(node *ListNode) (*Elem, error) {
	if node == nil {
		return nil, fmt.Errorf("invalid list node")
	}
	delData := node.Data
	//delete node in O(1)
	if node.Next != nil {
		next := node.Next
		node.Data = next.Data
		node.Next = next.Next
		next.Data = nil
		next = nil
	} else if node == l.head.Next {
		//delete the first node in O(1)
		l.head.Next = nil
	} else {
		//delete the last node in O(n)
		last := l.head.Next
		for last.Next != node {
			last = last.Next
		}
		last.Next = nil
	}
	return &delData, nil
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
	if k < 0 || k >= l.length {
		return nil, false
	}

	l.lock.RLock()
	defer l.lock.RUnlock()

	last := l.head.Next
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
	last := l.head.Next
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
	return l.head.Next == nil
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

	node := l.head.Next
	var b bytes.Buffer
	for node.Next != nil {
		b.WriteString(fmt.Sprintf("%v", node.Data))
		b.WriteString("->")
		node = node.Next
	}
	b.WriteString(fmt.Sprintf("%v", node.Data))
	return b.String()
}

// FirstNode of the linked list
func (l *LinkedList) FirstNode() *ListNode {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.head.Next
}

func (l *LinkedList) Head() *ListNode {
	return l.head
}
