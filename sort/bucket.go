package sort

import "math"

func BucketSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	bucketSort(arr)
	return arr
}

func bucketSort(arr []int) {
	n := len(arr)
	bucket := make([][]int, n)
	max := maxValue(arr, 0, n-1)

	// put element into bucket
	for i := 0; i < n; i++ {
		idx := arr[i] * (n - 1) / max
		bucket[idx] = append(bucket[idx], arr[i])
	}

	// sort bucket one by one
	idx := 0
	for i := 0; i < n; i++ {
		if len(bucket[i]) > 0 {
			QuickSort(bucket[i])
			copy(arr[idx:], bucket[i])
			idx += len(bucket[i])
		}
	}
}

func maxValue(arr []int, lo, hi int) int {
	max := math.MinInt64
	for i := lo; i <= hi; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
