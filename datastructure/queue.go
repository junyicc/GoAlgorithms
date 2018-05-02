package datastructure

import (
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

// New queue
func (q *Queue) New(n int) error {
	if n < 1 {
		return fmt.Errorf("queue size must larger than 0")
	}
	q.data = make([]Elem, n, n)
	q.front = 0
	q.rear = 0
	return nil
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

// EnQueue an element
func (q *Queue) EnQueue(e Elem) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	if (q.rear+1)%(cap(q.data)) == q.front {
		return fmt.Errorf("queue is full")
	}
	q.data[q.rear] = e
	q.rear = (q.rear + 1) % cap(q.data)
	return nil
}

// DeQueue an element
func (q *Queue) DeQueue() (Elem, error) {
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
