package sort

func CountingSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	return countingSort(arr)
}

func countingSort(arr []int) []int {
	n := len(arr)
	max := maxValue(arr, 0, n-1)

	// counting
	cnts := make([]int, max+1)
	for _, e := range arr {
		cnts[e]++
	}
	// accumulating
	for i := 1; i < max+1; i++ {
		cnts[i] += cnts[i-1]
	}
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		idx := cnts[arr[i]] - 1
		res[idx] = arr[i]
		cnts[arr[i]]--
	}
	return res
}

func SimpleCountingSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	return simpleCountingSort(arr)
}

func simpleCountingSort(arr []int) []int {
	n := len(arr)
	max := maxValue(arr, 0, n-1)

	// counting
	cnts := make([]int, max+1)
	for _, e := range arr {
		cnts[e]++
	}
	res := make([]int, 0)
	for i, cnt := range cnts {
		for ; cnt > 0; cnt-- {
			res = append(res, i)
		}
	}
	return res
}
