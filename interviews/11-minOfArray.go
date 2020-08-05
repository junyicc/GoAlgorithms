package interviews

func minArray(numbers []int) int {
	if len(numbers) < 1 {
		return 0
	}

	if len(numbers) == 1 {
		return numbers[0]
	}

	lo, hi := 0, len(numbers)-1
	for lo+1 < hi {
		for lo < hi && numbers[lo] == numbers[lo+1] {
			lo++
		}
		for lo < hi && numbers[hi] == numbers[hi-1] {
			hi--
		}

		mi := lo + (hi-lo)>>1
		// update hi in the first place
		if numbers[mi] <= numbers[hi] {
			hi = mi
		} else if numbers[mi] >= numbers[lo] {
			lo = mi
		}
	}

	if numbers[hi] < numbers[lo] {
		return numbers[hi]
	}

	return numbers[lo]
}

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
