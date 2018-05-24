package datastructure

import (
	"bytes"
	"container/heap"
	"fmt"
	"sync"
)

// Heap is a priority queue structure
type Heap struct {
	data []Elem
	lock sync.RWMutex
}

// Less comparison
func (h *Heap) Less(i, j int) bool {
	if i < 0 || j < 0 || i >= h.Len() || j >= h.Len() {
		return false
	}
	switch h.data[0].(type) {
	case nil:
		return false
	case int, int8, int16, int32, int64:
		return h.data[i].(int) < h.data[j].(int)
	case uint, uint8, uint16, uint32, uint64:
		return h.data[i].(uint) < h.data[j].(uint)
	case float32, float64:
		return h.data[i].(float64) < h.data[j].(float64)
	case string:
		return h.data[i].(string) < h.data[j].(string)
	default:
		return false
	}
}

// Swap i, j element
func (h *Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Len returns heap size
func (h *Heap) Len() int {
	return len(h.data)
}

// Heapify builds up a heap
func (h *Heap) Heapify(arr []Elem) {
	h.lock.Lock()
	defer h.lock.Unlock()
	n := len(arr)
	h.data = arr
	for i := (n - 2) >> 1; i >= 0; i-- {
		h.down(i, n)
	}
}

// Insert an element into heap
func (h *Heap) Insert(e Elem) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.data = append(h.data, e)
	h.up(h.Len() - 1)
}

// Remove element at i index
func (h *Heap) Remove(i int) Elem {
	if i < 0 || i >= h.Len() {
		return nil
	}

	h.lock.Lock()
	defer h.lock.Unlock()

	n := h.Len() - 1
	if n-1 != i {
		h.Swap(n, i)
		if !h.down(i, n) {
			h.up(i)
		}
	}
	e := h.data[n]
	h.data = h.data[:n]
	return e
}

// DelMin deletes the root of heap
func (h *Heap) DelMin() Elem {
	return h.Remove(0)
}

func (h *Heap) down(lo, hi int) bool {
	i := lo
	for j := 2*i + 1; j < hi; j = 2*i + 1 {
		if j < hi-1 && h.Less(j+1, j) {
			j++
		}
		if h.Less(i, j) {
			break
		}
		h.Swap(i, j)
		i = j
	}

	return i > lo
}

func (h *Heap) up(i int) {
	for {
		p := (i - 1) >> 1 // parent
		if i == p || !h.Less(i, p) {
			break
		}
		h.Swap(i, p)
		i = p
	}
}

func (h *Heap) String() string {
	h.lock.RLock()
	defer h.lock.RUnlock()

	var b bytes.Buffer
	for i, e := range h.data {
		b.WriteString(fmt.Sprintf("%v", e))
		if i != h.Len()-1 {
			b.WriteString(" ")
		}
	}
	return b.String()
}

// ----------------------------------------------------------------------------

// PQItem struct
type PQItem struct {
	Value    Elem
	Priority int
	Index    int
}

// PriorityQueue struct
type PriorityQueue []*PQItem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Priority > pq[j].Priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i]; pq[i].Index = i; pq[j].Index = j }

// Push item
func (pq *PriorityQueue) Push(item interface{}) {
	n := pq.Len()
	pqItem := item.(*PQItem)
	pqItem.Index = n
	*pq = append(*pq, pqItem)
}

// Pop item
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := old.Len()
	item := old[n-1]
	*pq = old[:n-1]

	return item
}

// Update item
func (pq *PriorityQueue) Update(item *PQItem, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

// ----------------------------------------------------------------------------

// IntHeap is a heap structure for integer
type IntHeap struct {
	Data []int
	Cmp  func(i, j int) bool
}

func (ih IntHeap) Len() int           { return len(ih.Data) }
func (ih IntHeap) Less(i, j int) bool { return ih.Cmp(i, j) }
func (ih IntHeap) Swap(i, j int)      { ih.Data[i], ih.Data[j] = ih.Data[j], ih.Data[i] }

// Push an integer
func (ih *IntHeap) Push(x interface{}) {
	(*ih).Data = append((*ih).Data, x.(int))
}

// Pop an integer
func (ih *IntHeap) Pop() interface{} {
	old := (*ih).Data
	n := ih.Len()
	x := old[n-1]
	(*ih).Data = old[:n-1]
	return x
}
