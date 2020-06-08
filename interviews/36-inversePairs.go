package interviews

// InversePairs returns the number of inverse pairs in the array
func InversePairs(arr []int) int {
	if arr == nil || len(arr) < 1 {
		return 0
	}

	return inversePairs(arr, 0, len(arr))
}

func inversePairs(arr []int, lo, hi int) int {
	if hi-lo < 2 {
		return 0
	}

	mi := lo + (hi-lo)>>1
	lInvPairs := inversePairs(arr, lo, mi)
	rInvPairs := inversePairs(arr, mi, hi)

	// merge tmp to dst
	tmp := make([]int, hi-lo)
	i, j, k := lo, mi, 0
	invPairs := 0
	for ; i < mi && j < hi; k++ {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			invPairs += mi - i
			tmp[k] = arr[j]
			j++
		}
	}
	for i < mi {
		tmp[k] = arr[i]
		i++
		k++
	}
	for j < hi {
		tmp[k] = arr[j]
		j++
		k++
	}
	copy(arr[lo:hi], tmp)
	return lInvPairs + rInvPairs + invPairs
}

// ReversePairs returns the number of important reverse pairs where if i < j and nums[i] > 2*nums[j]
func ReversePairs(arr []int) int {
	if arr == nil || len(arr) < 1 {
		return 0
	}

	return reversePairs(arr, arr, 0, len(arr))
}

func reversePairs(arr, dst []int, lo, hi int) int {
	if hi-lo < 2 {
		dst[lo] = arr[lo]
		return 0
	}

	tmp := make([]int, len(arr))
	mi := (lo + hi) >> 1
	cnt := reversePairs(arr, tmp, lo, mi) + reversePairs(arr, tmp, mi, hi)

	// count inverse pairs
	index := hi - 1
	for i := mi - 1; i >= lo; i-- {
		for index >= mi && tmp[i] <= 2*tmp[index] {
			index--
		}
		cnt += (index - mi + 1)
	}
	// merge tmp to dst
	i, j, k := mi-1, hi-1, hi-1
	for ; i >= lo && j >= mi; k-- {
		if tmp[i] > tmp[j] {
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
	return cnt
}
