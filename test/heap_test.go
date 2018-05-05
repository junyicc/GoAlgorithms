package test

import (
	"testing"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

var heap datastructure.Heap

func TestHeapify(t *testing.T) {
	data := []datastructure.Elem{9, 6, 5, 2, 3}
	exp := "2 3 5 6 9"
	heap.Heapify(data)
	result := heap.String()
	if result != exp {
		t.Errorf("expected %s and got %s", exp, result)
	}
}

func TestHeapInsert(t *testing.T) {
	data := []datastructure.Elem{9, 6, 5, 2, 3}
	heap.Heapify(data)
	heap.Insert(4)
	exp := "2 3 4 6 9 5"
	result := heap.String()
	if result != exp {
		t.Errorf("expected %s and got %s", exp, result)
	}
}

func TestHeapRemove(t *testing.T) {
	data := []datastructure.Elem{9, 6, 5, 2, 3}
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
