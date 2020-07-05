package interviews

// CountingK returns the number of k in the sorted array
func CountingK(arr []int, k int) int {
	if len(arr) < 1 {
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
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if arr[mi] < k {
			lo = mi + 1
		} else {
			if mi == 0 || arr[mi-1] < k {
				return mi
			} else {
				hi = mi - 1
			}
		}
	}
	return -1
}
func getLastK(arr []int, k int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if arr[mi] > k {
			hi = mi - 1
		} else {
			if mi == len(arr)-1 || arr[mi+1] > k {
				return mi
			} else {
				lo = mi + 1
			}
		}
	}
	return -1
}
