package test

import (
	"algorithm/sort"
	"math/rand"
	"testing"
)

// TestQuickSort tests QuickSort algorithm
func TestBubbleSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := sort.BubbleSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := sort.BubbleSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := sort.BubbleSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestSelectSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := sort.SelectSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := sort.SelectSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := sort.SelectSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestInsertSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := sort.InsertSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := sort.InsertSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := sort.InsertSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestShellSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := sort.ShellSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := sort.ShellSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := sort.ShellSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestMergeSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := sort.MergeSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := sort.MergeSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := sort.MergeSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestQuickSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := sort.QuickSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := sort.QuickSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := sort.QuickSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestHeapSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := sort.HeapSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := sort.HeapSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := sort.HeapSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	sort.BubbleSort(data)
}

func BenchmarkSelectSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	sort.SelectSort(data)
}
func BenchmarkInsertSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	sort.InsertSort2(data)
}
func BenchmarkShellSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	sort.ShellSort(data)
}

func BenchmarkHeapSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	sort.HeapSort(data)
}

func BenchmarkMergeSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	sort.MergeSortRecur(data)
}

func BenchmarkQuickSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	sort.QuickSort(data)
}
