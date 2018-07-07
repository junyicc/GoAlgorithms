package search

// InterpolationSearch searches element e in the array
func InterpolationSearch(arr []int, e int) (int, bool) {
	if arr == nil || len(arr) < 1 {
		return -1, false
	}

	lo, hi := 0, len(arr)-1
	for lo < hi {
		mi := lo + (hi-lo)*(e-arr[lo])/(arr[hi]-arr[lo])
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
