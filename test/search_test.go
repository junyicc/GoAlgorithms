package test

import (
	"testing"

	"github.com/CAIJUNYI/GoAlgorithms/search"
)

// TestLinearSearch tests linear search algorithm

func TestLinearSearch(t *testing.T) {
	items := []int{5, 8, 46, 58, 59, 86, 99, 251, 320}
	if i, _ := search.BinarySearch(items, 5); i != 0 {
		t.Errorf("expected 0 and got %d", i)
	}
}

// TestBinarySearch tests binary search algorithm
func TestBinarySearch(t *testing.T) {
	items := []int{5, 8, 46, 58, 59, 86, 99, 251, 320}
	if i, ok := search.BinarySearch(items, -1); i != -1 || ok {
		t.Errorf("expected -1 and got %d", i)
	}
}

// TestInterpolationSearch tests binary search algorithm
func TestInterpolationSearch(t *testing.T) {
	items := []int{5, 8, 46, 58, 59, 86, 99, 251, 320}
	if i, _ := search.InterpolationSearch(items, 320); i != 8 {
		t.Errorf("expected 8 and got %d", i)
	}
}

// TestFibSearch tests fibonacci search
func TestFibSearch(t *testing.T) {
	items := []int{5, 8, 46, 58, 59, 86, 99, 251, 320}
	if i, _ := search.FibSearch(items, 251); i != 7 {
		t.Errorf("expected 1 and got %d", i)
	}
}
