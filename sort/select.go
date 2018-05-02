package sort

// SelectSort algorithm
func SelectSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	lo, hi := 0, len(arr)-1
	for ; lo <= hi; hi-- {
		maxIndex := selectMax(arr, lo, hi)
		arr[hi], arr[maxIndex] = arr[maxIndex], arr[hi]
	}
	return arr
}

func selectMax(arr []int, lo, hi int) int {
	maxIndex := lo
	for lo++; lo <= hi; lo++ {
		if arr[lo] >= arr[maxIndex] {
			maxIndex = lo
		}
	}
	return maxIndex
}
