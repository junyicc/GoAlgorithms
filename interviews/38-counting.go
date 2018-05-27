package interviews

// CountingK returns the number of k in the sorted array
func CountingK(arr []int, k int) int {
	if arr == nil || len(arr) < 1 {
		return 0
	}
	if len(arr) == 1 {
		if arr[0] == k {
			return 1
		}
		return 0
	}

	firstIndex := getFirstK(arr, k)
	lastIndex := getLastK(arr, k)
	cnt := 0
	if firstIndex >= 0 && lastIndex >= 0 {
		cnt = lastIndex - firstIndex + 1
	}
	return cnt
}

func getFirstK(arr []int, k int) int {
	lo, hi := 0, len(arr)
	var mi int
	for lo < hi {
		mi = (lo + hi) >> 1
		if arr[mi] == k && arr[mi-1] < k {
			break
		} else if arr[mi] < k {
			lo = mi + 1
		} else {
			hi = mi
		}
	}
	if lo < hi {
		return mi
	}
	return -1
}
func getLastK(arr []int, k int) int {
	lo, hi := 0, len(arr)
	var mi int
	for lo < hi {
		mi = (lo + hi) >> 1
		if arr[mi] == k && arr[mi+1] > k {
			break
		} else if arr[mi] > k {
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	if lo < hi {
		return mi
	}
	return -1
}
