package search

import (
	"sort"
	"testing"
)

// TestLinearSearch tests linear search algorithm

func TestLinearSearch(t *testing.T) {
	items := []int{5, 8, 46, 58, 59, 86, 99, 251, 320}
	if i, _ := BinarySearch(items, 5); i != 0 {
		t.Errorf("expected 0 and got %d", i)
	}
}

// TestBinarySearch tests binary search algorithm
func TestBinarySearch(t *testing.T) {
	items := []int{5, 8, 46, 58, 59, 86, 99, 251, 320}
	if i, ok := BinarySearch(items, -1); i != -1 || ok {
		t.Errorf("expected -1 and got %d", i)
	}

	items = []int{1, 3, 4, 5, 6, 8, 8, 8, 11, 11, 18}
	if i := BinarySearchFirst(items, 8); i != 5 {
		t.Errorf("expected 5, got %d", i)
	}
	if i := BinarySearchLast(items, 11); i != 9 {
		t.Errorf("expected 9, got %d", i)
	}
	if i := BinarySearchFirstGt(items, 0); i != 0 {
		t.Errorf("expected 0, got %d", i)
	}
	if i := BinarySearchLastLt(items, 19); i != 10 {
		t.Errorf("expected 10, got %d", i)
	}
	// built-in search
	// first gt
	if i := sort.SearchInts(items, 2); i != 1 {
		t.Errorf("expected 1, got %d", i)
	}

	if i := sort.Search(len(items), func(i int) bool {
		return items[i] >= 8
	}); i != 5 {
		t.Errorf("expected 5, got %d", i)
	}
}

// TestInterpolationSearch tests binary search algorithm
func TestInterpolationSearch(t *testing.T) {
	items := []int{5, 8, 46, 58, 59, 86, 99, 251, 320}
	if i, _ := InterpolationSearch(items, 320); i != 8 {
		t.Errorf("expected 8 and got %d", i)
	}
}

// TestFibSearch tests fibonacci search
func TestFibSearch(t *testing.T) {
	items := []int{5, 8, 46, 58, 59, 86, 99, 251, 320}
	if i, _ := FibSearch(items, 251); i != 7 {
		t.Errorf("expected 1 and got %d", i)
	}
}
