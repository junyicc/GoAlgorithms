package search

// BinarySearch searches element e in binary in the array
func BinarySearch(arr []int, e int) (int, bool) {
	if arr == nil || len(arr) < 1 {
		return -1, false
	}

	lo, hi := 0, len(arr)
	for hi-lo > 1 {
		mi := lo + (hi-lo)>>1 // lo+hi may overflow
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

func SimpleBinarySearch(arr []int, e int) int {
	if len(arr) < 1 {
		return -1
	}

	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if arr[mi] == e {
			return mi
		} else if arr[mi] < e {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	return -1
}

// BinarySearchFirst search the first e, and multiple e may exist
func BinarySearchFirst(arr []int, e int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	lo, hi := 0, n-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if arr[mi] < e {
			lo = mi + 1
		} else if e < arr[mi] {
			hi = mi - 1
		} else {
			if mi == 0 || arr[mi-1] != e {
				return mi
			} else {
				hi = mi - 1
			}
		}
	}
	return -1
}

// BinarySearchLast search the last e, and multiple e may exist
func BinarySearchLast(arr []int, e int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	lo, hi := 0, n-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if arr[mi] < e {
			lo = mi + 1
		} else if e < arr[mi] {
			hi = mi - 1
		} else {
			if mi == n-1 || arr[mi+1] != e {
				return mi
			} else {
				lo = mi + 1
			}
		}
	}
	return -1
}

// BinarySearchFirstGt search the first element greater than e,
// e may not exist
func BinarySearchFirstGt(arr []int, e int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	lo, hi := 0, n-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if arr[mi] < e {
			lo = mi + 1
		} else {
			if mi == 0 || arr[mi-1] < e {
				return mi
			} else {
				hi = mi - 1
			}
		}
	}
	return -1
}

// BinarySearchLastLt search the last element less than e,
// e may not exist
func BinarySearchLastLt(arr []int, e int) int {
	n := len(arr)
	if n < 1 {
		return -1
	}
	lo, hi := 0, n-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if e < arr[mi] {
			hi = mi - 1
		} else {
			if mi == n-1 || arr[mi+1] > e {
				return mi
			} else {
				lo = mi + 1
			}
		}
	}
	return -1
}
