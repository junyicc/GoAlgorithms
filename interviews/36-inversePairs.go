package interviews

// InversePairs returns the number of inverse pairs in the array
func InversePairs(arr []int) int {
	if arr == nil || len(arr) < 1 {
		return 0
	}

	return inversePairs(arr, arr, 0, len(arr))
}

func inversePairs(arr, dst []int, lo, hi int) int {
	if hi-lo < 2 {
		dst[lo] = arr[lo]
		return 0
	}

	tmp := make([]int, len(arr))
	mi := (lo + hi) >> 1
	lInvPairs := inversePairs(arr, tmp, lo, mi)
	rInvPairs := inversePairs(arr, tmp, mi, hi)

	// merge tmp to dst
	invPairs := 0
	i, j, k := mi-1, hi-1, hi-1
	for ; i >= lo && j >= mi; k-- {
		if tmp[i] > tmp[j] {
			invPairs += (j - mi + 1)
			dst[k] = tmp[i]
			i--
		} else {
			dst[k] = tmp[j]
			j--
		}
	}
	for ; i >= lo; i-- {
		dst[k] = tmp[i]
		k--
	}
	for ; j >= mi; j-- {
		dst[k] = tmp[j]
		k--
	}
	return lInvPairs + rInvPairs + invPairs
}
