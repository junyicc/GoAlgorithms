package interviews

import (
	"container/heap"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// FindKthMin returns the kth min number in the array
func FindKthMin(arr []int, k int) (int, bool) {
	if arr == nil || len(arr) < 1 || k < 0 || k >= len(arr) {
		return 0, false
	}
	lo, hi := 0, len(arr)-1
	index := increasingPartition(arr, lo, hi)
	for index != k-1 {
		if index < k-1 {
			lo = index + 1
		} else {
			hi = index - 1
		}
		index = increasingPartition(arr, lo, hi)
	}
	return arr[index], true
}

// FindKthMax returns the kth max number in the array
func FindKthMax(arr []int, k int) (int, bool) {
	if arr == nil || len(arr) < 1 || k < 0 || k >= len(arr) {
		return 0, false
	}
	lo, hi := 0, len(arr)-1
	index := decreasingPartition(arr, lo, hi)
	for index != k-1 {
		if index < k-1 {
			lo = index + 1
		} else {
			hi = index - 1
		}
		index = decreasingPartition(arr, lo, hi)
	}
	return arr[index], true
}

// GetKthLeast returns the kth least numbers in the array
func GetKthLeast(arr []int, k int) []int {
	if arr == nil || len(arr) < 1 || k < 0 || k >= len(arr) {
		return nil
	}
	lo, hi := 0, len(arr)-1
	index := increasingPartition(arr, lo, hi)
	for index != k-1 {
		if index < k-1 {
			lo = index + 1
		} else {
			hi = index - 1
		}
		index = increasingPartition(arr, lo, hi)
	}
	result := arr[:k]
	return result
}

func increasingPartition(arr []int, lo, hi int) int {
	pivot := getPivot(arr, lo, hi)
	leftMark := lo + 1
	rightMark := hi
	for {
		for leftMark <= rightMark && pivot <= arr[rightMark] {
			rightMark--
		}
		for leftMark <= rightMark && arr[leftMark] <= pivot {
			leftMark++
		}
		if rightMark < leftMark {
			break
		} else {
			arr[leftMark], arr[rightMark] = arr[rightMark], arr[leftMark]
		}
	}
	arr[lo], arr[rightMark] = arr[rightMark], arr[lo]
	return rightMark
}

func decreasingPartition(arr []int, lo, hi int) int {
	pivot := getPivot(arr, lo, hi)
	leftMark := lo + 1
	rightMark := hi
	for {
		for leftMark <= rightMark && pivot >= arr[rightMark] {
			rightMark--
		}
		for leftMark <= rightMark && arr[leftMark] >= pivot {
			leftMark++
		}
		if rightMark < leftMark {
			break
		} else {
			arr[leftMark], arr[rightMark] = arr[rightMark], arr[leftMark]
		}
	}
	arr[lo], arr[rightMark] = arr[rightMark], arr[lo]
	return rightMark
}

// getPivot returns median of three
func getPivot(arr []int, lo, hi int) int {
	mi := (lo + hi) >> 1

	if arr[lo] > arr[hi] {
		arr[lo], arr[hi] = arr[hi], arr[lo]
	}
	if arr[mi] > arr[hi] {
		arr[mi], arr[hi] = arr[hi], arr[mi]
	}
	if arr[mi] > arr[lo] {
		arr[mi], arr[lo] = arr[lo], arr[mi]
	}
	return arr[lo]
}

// GetMinKth returns kth min numbers in the array using heap
func GetMinKth(arr []int, k int) []int {
	if arr == nil || len(arr) < 1 || k < 0 || k >= len(arr) {
		return nil
	}

	intHeap := datastructure.IntHeap{
		Data: []int{},
	}
	intHeap.Cmp = func(i, j int) bool {
		return intHeap.Data[i] > intHeap.Data[j]
	}
	heap.Init(&intHeap)

	for _, e := range arr {
		if intHeap.Len() < k {
			heap.Push(&intHeap, e)
		} else {
			item := heap.Pop(&intHeap).(int)
			if e < item {
				heap.Push(&intHeap, e)
			} else {
				heap.Push(&intHeap, item)
			}
		}
	}
	return intHeap.Data
}

// GetMaxKth returns kth max numbers in the array using heap
func GetMaxKth(arr []int, k int) []int {
	if arr == nil || len(arr) < 1 || k < 0 || k >= len(arr) {
		return nil
	}

	intHeap := datastructure.IntHeap{
		Data: []int{},
	}
	intHeap.Cmp = func(i, j int) bool {
		return intHeap.Data[i] < intHeap.Data[j]
	}
	heap.Init(&intHeap)

	for _, e := range arr {
		if intHeap.Len() < k {
			heap.Push(&intHeap, e)
		} else {
			item := heap.Pop(&intHeap).(int)
			if e > item {
				heap.Push(&intHeap, e)
			} else {
				heap.Push(&intHeap, item)
			}
		}
	}
	return intHeap.Data
}
