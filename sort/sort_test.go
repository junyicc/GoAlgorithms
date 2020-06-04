package sort

import (
	"math/rand"
	"testing"
)

// TestQuickSort tests QuickSort algorithm
func TestBubbleSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := BubbleSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := BubbleSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := BubbleSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestSelectSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := SelectSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := SelectSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := SelectSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestInsertSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := InsertSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := InsertSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := InsertSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestShellSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := ShellSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := ShellSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := ShellSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestMergeSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := MergeSortRecur(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := MergeSortRecur(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := MergeSortRecur(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestQuickSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	QuickSort(data1)

	for i := 0; i < len(data1); i++ {
		if data1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, data1)
		}
	}

	var data2 []int
	QuickSort(data2)
	if len(data2) > 1 {
		t.Errorf("expected nil and got %v", data2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	QuickSort(data3)
	if data3[0] != 3 || len(data3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, data3)
	}
}

func TestHeapSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := HeapSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := HeapSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := HeapSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestBucketSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := BucketSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := BucketSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := BucketSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestSimpleCountingSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := SimpleCountingSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Errorf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := SimpleCountingSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := SimpleCountingSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

func TestCountingSort(t *testing.T) {
	data1 := []int{3, 1, 7, 0, 3, 5, 23, 56, 19, 59}
	ans := []int{0, 1, 3, 3, 5, 7, 19, 23, 56, 59}
	sortedResult1 := CountingSort(data1)

	for i := 0; i < len(sortedResult1); i++ {
		if sortedResult1[i] != ans[i] {
			t.Fatalf("expected %v\n got %v", ans, sortedResult1)
		}
	}

	var data2 []int
	sortedResult2 := CountingSort(data2)
	if sortedResult2 != nil {
		t.Errorf("expected nil and got %v", sortedResult2)
	}

	data3 := []int{3}
	ans3 := []int{3}
	sortedResult3 := CountingSort(data3)
	if sortedResult3[0] != 3 || len(sortedResult3) != 1 {
		t.Errorf("expected %v\n got %v", ans3, sortedResult3)
	}
}

// --- benchmark ---

func BenchmarkBubbleSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	BubbleSort(data)
}

func BenchmarkSelectSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	SelectSort(data)
}
func BenchmarkInsertSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	InsertSort2(data)
}
func BenchmarkShellSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	ShellSort(data)
}

func BenchmarkHeapSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	HeapSort(data)
}

func BenchmarkMergeSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	MergeSort(data)
}

func BenchmarkQuickSort(b *testing.B) {
	data := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		data[i] = rand.Intn(b.N)
	}
	QuickSort(data)
}
