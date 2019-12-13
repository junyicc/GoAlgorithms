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
		data:  make([]Elem, n),
		front: 0,
		rear:  0,
		lock:  sync.RWMutex{},
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
	return q.rear == q.front
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

//----------------------------------------------------------------------------------------------------------------------

//CircularQueue is a circular queue
type CircularQueue struct {
	data  []Elem
	front int
	rear  int
	lock  sync.RWMutex
}

func NewCircularQueue(n int) *CircularQueue {
	if n < 1 {
		return nil
	}
	return &CircularQueue{
		data:  make([]Elem, n),
		front: 0,
		rear:  0,
		lock:  sync.RWMutex{},
	}
}

func (cq *CircularQueue) IsEmpty() bool {
	cq.lock.RLock()
	defer cq.lock.RUnlock()
	return cq.front == cq.rear
}

func (cq *CircularQueue) IsFull() bool {
	cq.lock.RLock()
	defer cq.lock.RUnlock()
	return (cq.rear+1)%cap(cq.data) == cq.front
}

func (cq *CircularQueue) Length() int {
	cq.lock.RLock()
	defer cq.lock.RUnlock()
	return (cq.rear - cq.front + cap(cq.data)) % cap(cq.data)
}

func (cq *CircularQueue) Enqueue(e Elem) error {
	if cq.IsFull() {
		return fmt.Errorf("queue is full")
	}
	cq.lock.Lock()
	cq.data[cq.rear] = e
	cq.rear = (cq.rear + 1) % cap(cq.data)
	cq.lock.Unlock()
	return nil
}

func (cq *CircularQueue) Dequeue() (Elem, error) {
	if cq.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}
	cq.lock.Lock()
	e := cq.data[cq.front]
	cq.front = (cq.front + 1) % cap(cq.data)
	cq.lock.Unlock()
	return e, nil
}

//----------------------------------------------------------------------------------------------------------------------

//BlockingQueue can block when enqueue or dequeue
type BlockingQueue struct {
	C    chan Elem
	mu   sync.RWMutex
	done chan struct{}
}

func NewBlockingQueue(n int) *BlockingQueue {
	if n < 0 {
		return nil
	}
	return &BlockingQueue{
		C:  make(chan Elem, n),
		mu: sync.RWMutex{},
	}
}

func (bq *BlockingQueue) Length() int {
	bq.mu.RLock()
	defer bq.mu.RUnlock()
	return len(bq.C)
}

func (bq *BlockingQueue) Dequeue() (Elem, error) {
	e, ok := <-bq.C
	if !ok {
		return nil, fmt.Errorf("blocking queue has been closed")
	}
	return e, nil
}

func (bq *BlockingQueue) Enqueue(e Elem) error {
	select {
	case <-bq.done:
		return fmt.Errorf("cannot enqueue an closed blocking queue")
	default:
	}
	bq.C <- e
	return nil
}

func (bq *BlockingQueue) Close() {
	close(bq.C)
	close(bq.done)
}

//----------------------------------------------------------------------------------------------------------------------

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
func (q *DynamicQueue) Dequeue() Elem {
	q.lock.Lock()
	e := q.data[0]
	q.data = q.data[1:]
	q.lock.Unlock()
	return e
}

// Front of queue
func (q *DynamicQueue) Front() Elem {
	q.lock.RLock()
	e := q.data[0]
	q.lock.RUnlock()
	return e
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

//----------------------------------------------------------------------------------------------------------------------

const (
	doEnqueue = iota
	doDequeue
	doFront
	doEmpty
)

type queueItem struct {
	op    int
	value Elem
	res   chan interface{}
}

//ConcurrentQueue implemented in client-server mode with the help of channel
type ConcurrentQueue struct {
	data *DynamicQueue
	opC  chan *queueItem
}

func (cq *ConcurrentQueue) Enqueue(value Elem) {
	cq.doOperation(doEnqueue, value)
}

func (cq *ConcurrentQueue) Dequeue() Elem {
	e := cq.doOperation(doDequeue, nil)
	return e.(Elem)
}

func (cq *ConcurrentQueue) Front() Elem {
	e := cq.doOperation(doFront, nil)
	return e.(Elem)
}

func (cq *ConcurrentQueue) IsEmpty() bool {
	b := cq.doOperation(doEmpty, nil)
	return b.(bool)
}

func (cq *ConcurrentQueue) doOperation(op int, value interface{}) interface{} {
	item := &queueItem{
		op:    op,
		value: value,
		res:   make(chan interface{}),
	}
	cq.opC <- item
	return <-item.res
}

func (cq *ConcurrentQueue) run() {
	deferQueue := NewDynamicQueue()
	for {
		var item *queueItem
		if !deferQueue.IsEmpty() && !cq.IsEmpty() {
			it := deferQueue.Dequeue()
			item = it.(*queueItem)
		} else {
			it := <-cq.opC
			if item.op == doDequeue && cq.IsEmpty() {
				//defer dequeue operation when queue is empty
				deferQueue.Enqueue(it)
				continue
			}
		}

		switch item.op {
		case doEnqueue:
			cq.data.Enqueue(item.value)
			item.res <- nil
		case doDequeue:
			e := cq.data.Dequeue()
			item.res <- e
		case doFront:
			e := cq.data.Front()
			item.res <- e
		case doEmpty:
			b := cq.data.IsEmpty()
			item.res <- b
		}
	}
}
