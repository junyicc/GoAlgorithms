package sort

// InsertSort algorithm
func InsertSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			tmp := arr[i]
			j := i - 1
			// move backward
			for ; 0 <= j && tmp < arr[j]; j-- {
				arr[j+1] = arr[j]
			}
			arr[j+1] = tmp
		}
	}
	return arr
}

// InsertSort2 faster algorithm
func InsertSort2(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	lo, hi := 0, len(arr)
	for i := lo + 1; i < hi; i++ {
		pos := i
		for lo < pos && arr[pos] < arr[pos-1] {
			arr[pos], arr[pos-1] = arr[pos-1], arr[pos]
			pos--
		}
	}
	return arr
}
