package sort

// MergeSortRecur algorithm recursive version
func MergeSortRecur(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	lo, hi := 0, len(arr)
	mergeSortRecur(arr, lo, hi)
	return arr
}

func mergeSortRecur(arr []int, lo, hi int) {
	if hi-lo < 2 {
		return
	}
	mi := lo + (hi-lo)>>1

	mergeSortRecur(arr, lo, mi)
	mergeSortRecur(arr, mi, hi)

	mergeRecur(arr, lo, mi, hi)

}

func mergeRecur(arr []int, lo, mi, hi int) {
	tmp := make([]int, hi-lo)
	i, j, k := lo, mi, 0
	for ; i < mi && j < hi; k++ {
		if arr[i] < arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
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
}

// MergeSort algorithm iterative version
func MergeSort(arr []int) []int {
	length := len(arr)
	if arr == nil || length < 2 {
		return arr
	}
	seq := make([]int, length)
	for k := 1; k < length; {
		mergePass(arr, seq, k, length)
		k *= 2
		mergePass(seq, arr, k, length)
		k *= 2
	}
	return arr
}

func mergePass(source, target []int, k, n int) {
	var i int
	for i <= n-2*k {
		merge(source, target, i, i+k, i+2*k)
		i = i + 2*k
	}
	// merge the last pair
	if i < n-k {
		merge(source, target, i, i+k, n)
	} else {
		for j := i; j < n; j++ {
			target[j] = source[j]
		}
	}
}

func merge(source, target []int, lo, mi, hi int) {
	i, j, k := lo, mi, lo

	for ; i < mi && j < hi; k++ {
		if source[i] < source[j] {
			target[k] = source[i]
			i++
		} else {
			target[k] = source[j]
			j++
		}
	}
	for i < mi {
		target[k] = source[i]
		k++
		i++
	}
	for j < hi {
		target[k] = source[j]
		k++
		j++
	}
}
