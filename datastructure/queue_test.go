package datastructure

import (
	"testing"
	"time"
)

func TestCircularQueue(t *testing.T) {
	q := NewCircularQueue(5)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	if n, _ := q.Dequeue(); n.(int) != 1 {
		t.Errorf("expected %d, got %d", 6, n)
	}
	q.Enqueue(5)
	if n, _ := q.Dequeue(); n.(int) != 2 {
		t.Errorf("expected %d, got %d", 2, n)
	}
	q.Enqueue(6)
	if n, _ := q.Dequeue(); n.(int) != 3 {
		t.Errorf("expected %d, got %d", 3, n)
	}
	if n := q.Length(); n != 3 {
		t.Errorf("expected length %d, got %d", 3, n)
	}
}

func TestBlockingQueue(t *testing.T) {
	q := NewBlockingQueue(5)
	go func() {
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		q.Enqueue(4)
		q.Enqueue(5)
		q.Enqueue(6)
	}()
	time.Sleep(1 * time.Second)
	if n := q.Length(); n != 5 {
		t.Errorf("expected %d, got %d", 5, n)
	}
	t.Log("queue is blocking")
	for i := 1; i <= 6; i++ {
		n, err := q.Dequeue()
		if err != nil {
			t.Errorf("dequeue error: %v", err)
			continue
		}
		if n.(int) != i {
			t.Errorf("expected %d, got %d", i, n)
		}
	}
}
