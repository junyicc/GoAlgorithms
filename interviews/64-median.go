package interviews

import (
	"container/heap"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// MinHeap of integer heap
type MinHeap struct {
	datastructure.IntHeap
}

// MaxHeap of integer heap
type MaxHeap struct {
	datastructure.IntHeap
}

// DynamicArray is a dynamic array that can get median in O(1)
type DynamicArray struct {
	minHeap MinHeap
	maxHeap MaxHeap
}

// Init DynamicArray
func (da *DynamicArray) Init() {
	da.minHeap = MinHeap{}
	da.minHeap.Cmp = func(i, j int) bool {
		return da.minHeap.Data[i] < da.minHeap.Data[j]
	}
	da.maxHeap = MaxHeap{}
	da.maxHeap.Cmp = func(i, j int) bool {
		return da.maxHeap.Data[i] > da.maxHeap.Data[j]
	}
}

// Size of DynamicArray
func (da DynamicArray) Size() int {
	return da.minHeap.Len() + da.maxHeap.Len()
}

// Insert an item
func (da *DynamicArray) Insert(item int) {
	if da.Size()&1 == 0 {
		// even size
		if da.maxHeap.Len() > 0 && item < da.maxHeap.Data[0] {
			heap.Push(&da.maxHeap, item)
			item = heap.Pop(&da.maxHeap).(int)
		}
		heap.Push(&da.minHeap, item)
	} else {
		// odd size
		if da.minHeap.Len() > 0 && item > da.minHeap.Data[0] {
			heap.Push(&da.minHeap, item)
			item = heap.Pop(&da.minHeap).(int)
		}
		heap.Push(&da.maxHeap, item)
	}
}

// Median of DynamicArray
func (da DynamicArray) Median() (int, bool) {
	if da.Size() < 1 {
		return 0, false
	}
	var median int
	if da.Size()&1 == 0 {
		maxItem := da.maxHeap.Data[0]
		minItem := da.minHeap.Data[0]
		median = (maxItem + minItem) >> 1
	} else {
		median = da.minHeap.Data[0]
	}
	return median, true
}
