package search

// LinearSearch searches element e linearly in the array
func LinearSearch(arr []int, e int) (int, bool) {
	if arr == nil || len(arr) < 1 {
		return -1, false
	}
	for i, item := range arr {
		if item == e {
			return i, true
		}
	}
	return -1, false
}

// LinearSearchOpt is an improved linear search algorithm
func LinearSearchOpt(arr []int, e int) (int, bool) {
	n := len(arr)
	if arr == nil || n < 1 {
		return -1, false
	}

	arr = append(arr, e)
	i := 0
	for arr[i] != e {
		i++
	}
	if i != n {
		return i, true
	}
	return -1, false
}
