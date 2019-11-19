package datastructure

import (
	heapUtil "container/heap"
	"testing"
)

func TestHeapify(t *testing.T) {
	var heap Heap
	data := []Elem{9, 6, 5, 2, 3}
	exp := "2 3 5 6 9"
	heap.Heapify(data)
	result := heap.String()
	if result != exp {
		t.Errorf("expected %s and got %s", exp, result)
	}
}

func TestHeapInsert(t *testing.T) {
	var heap Heap
	data := []Elem{9, 6, 5, 2, 3}
	heap.Heapify(data)
	heap.Insert(4)
	exp := "2 3 4 6 9 5"
	result := heap.String()
	if result != exp {
		t.Errorf("expected %s and got %s", exp, result)
	}
}

func TestHeapRemove(t *testing.T) {
	var heap Heap
	data := []Elem{9, 6, 5, 2, 3}
	heap.Heapify(data)
	heap.Remove(1)

	exp1 := "2 6 5 9"
	result1 := heap.String()
	if result1 != exp1 {
		t.Errorf("expected %s and got %s", exp1, result1)
	}

	heap.DelMin()
	exp2 := "5 6 9"
	result2 := heap.String()
	if result2 != exp2 {
		t.Errorf("expected %s and got %s", exp2, result2)
	}
}

func TestPriorityQueue(t *testing.T) {
	company := map[string]int{
		"MS":     9,
		"Google": 11,
		"Apple":  12,
		"Amazon": 10,
		"IBM":    9,
	}
	pq := PriorityQueue{}
	i := 0
	for v, p := range company {
		pq.Push(&PQItem{
			Value:    v,
			Priority: p,
			Index:    i,
		})
		i++
	}

	heapUtil.Init(&pq)

	if item := (heapUtil.Pop(&pq)).(*PQItem); item.Value != "Apple" || item.Priority != 12 {
		t.Errorf("expected Apple and got %v", item.Value)
	}
	if item := (heapUtil.Pop(&pq)).(*PQItem); item.Value != "Google" || item.Priority != 11 {
		t.Errorf("expected Google and got %v", item.Value)
	}
	if item := (heapUtil.Pop(&pq)).(*PQItem); item.Value != "Amazon" || item.Priority != 10 {
		t.Errorf("expected Amazon and got %v", item.Value)
	}
}
