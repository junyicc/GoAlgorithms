package sort

// BubbleSort algorithm
func BubbleSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	lo, hi := 0, len(arr)
	for lo < hi {
		hi = bubble(arr, lo, hi)
	}
	return arr
}

// bubble returns the unordered index to the right
func bubble(arr []int, lo, hi int) int {
	last := lo
	for lo++; lo < hi; lo++ {
		if arr[lo-1] > arr[lo] {
			arr[lo-1], arr[lo] = arr[lo], arr[lo-1]
			last = lo
		}
	}
	return last
}
