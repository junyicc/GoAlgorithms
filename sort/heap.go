package sort

// HeapSort algorithm
func HeapSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	hi := len(arr)
	var i int
	// build heap
	for i = (hi - 2) >> 1; i > -1; i-- {
		heapAdjust(arr, i, hi)
	}
	// heap sort
	for i = hi - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapAdjust(arr, 0, i)
	}
	return arr
}

// heapAdjust arr[lo...hi] to be a max heap
func heapAdjust(arr []int, lo, hi int) {
	tmp := arr[lo]
	for i := (2*lo + 1); i < hi; i = 2*lo + 1 {
		// bigger child
		if i < hi-1 && arr[i] < arr[i+1] {
			i++
		}
		if tmp >= arr[i] {
			break
		}
		arr[lo] = arr[i]
		lo = i
	}
	arr[lo] = tmp
}
