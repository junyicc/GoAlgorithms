package search

// BinarySearch searches element e in binary in the array
func BinarySearch(arr []int, e int) (int, bool) {
	if arr == nil || len(arr) < 1 {
		return -1, false
	}

	lo, hi := 0, len(arr)
	for hi-lo > 1 {
		mi := (lo + hi) >> 1
		if e < arr[mi] {
			hi = mi
		} else {
			lo = mi
		}
	}
	if arr[lo] == e {
		return lo, true
	}
	return -1, false
}
