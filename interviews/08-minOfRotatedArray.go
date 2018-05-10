package interviews

// MinOfRotatedArray where min of first ascending part >= max of second ascending part
func MinOfRotatedArray(arr []int) (int, bool) {
	if arr == nil || len(arr) < 1 {
		return -1, false
	}

	lo, hi := 0, len(arr)-1
	min := arr[lo]
	for arr[hi] <= arr[lo] {
		if hi-lo == 1 {
			min = arr[hi]
			break
		}
		mi := (lo + hi) >> 1
		if arr[lo] == arr[mi] && arr[mi] == arr[hi] {
			min = minInOrder(arr)
			break
		}

		if arr[mi] >= arr[lo] {
			lo = mi
		} else if arr[mi] <= arr[hi] {
			hi = mi
		}
	}
	return min, true
}

func minInOrder(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] <= min {
			min = arr[i]
		}
	}
	return min
}
