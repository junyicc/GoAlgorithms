package datastructure

import (
	"bytes"
	"fmt"
	"sync"
)

// Queue struct
type Queue struct {
	data  []Elem
	front int
	rear  int
	lock  sync.RWMutex
}

// NewQueue new a queue
func NewQueue(n int) (*Queue, error) {
	if n < 1 {
		return nil, fmt.Errorf("queue size must larger than 0")
	}
	q := &Queue{
		data: make([]Elem, n, n),
	}

	return q, nil
}

// Length of queue: elment number
func (q *Queue) Length() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	n := (q.rear - q.front + cap(q.data)) % cap(q.data)
	return n
}

// IsEmpty queue
func (q *Queue) IsEmpty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if q.rear == q.front {
		return true
	}
	return false
}

// Enqueue an element
func (q *Queue) Enqueue(e Elem) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	if (q.rear+1)%(cap(q.data)) == q.front {
		return fmt.Errorf("queue is full")
	}
	q.data[q.rear] = e
	q.rear = (q.rear + 1) % cap(q.data)
	return nil
}

// Dequeue an element
func (q *Queue) Dequeue() (Elem, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	q.lock.Lock()
	defer q.lock.Unlock()
	e := q.data[q.front]
	q.front = (q.front + 1) % cap(q.data)
	return e, nil
}

// Front of queue
func (q *Queue) Front() Elem {
	q.lock.RLock()
	defer q.lock.RUnlock()

	e := q.data[q.front]
	return e
}

// DynamicQueue without limited size
type DynamicQueue struct {
	data []Elem
	lock sync.RWMutex
}

// NewDynamicQueue new an infinite queue
func NewDynamicQueue() *DynamicQueue {
	return &DynamicQueue{}
}

// Enqueue an element at rear
func (q *DynamicQueue) Enqueue(e Elem) {
	q.lock.Lock()
	q.data = append(q.data, e)
	q.lock.Unlock()
}

// Dequeue an element at front
func (q *DynamicQueue) Dequeue() *Elem {
	q.lock.Lock()
	e := q.data[0]
	q.data = q.data[1:]
	q.lock.Unlock()
	return &e
}

// Front of queue
func (q *DynamicQueue) Front() *Elem {
	q.lock.RLock()
	e := q.data[0]
	q.lock.RUnlock()
	return &e
}

// IsEmpty returns true if the queue is empty
func (q *DynamicQueue) IsEmpty() bool {
	return len(q.data) == 0
}

// Len returns the number of elements in the queue
func (q *DynamicQueue) Len() int {
	return len(q.data)
}

// String returns string that are composed of queue elements from front to rear
func (q *DynamicQueue) String() string {
	q.lock.RLock()
	defer q.lock.RUnlock()
	var b bytes.Buffer
	for _, e := range q.data {
		b.WriteString(fmt.Sprintf("%v ", e))
	}
	return b.String()
}
