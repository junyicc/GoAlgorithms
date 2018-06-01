package interviews

// FindRepeatingNum returns any repeating num in the array whose numbers are lt its length n
func FindRepeatingNum(arr []int) (int, bool) {
	if !isValidArray(arr) {
		return -1, false
	}
	for i := 0; i < len(arr); {
		if arr[i] != i {
			m := arr[i]
			if arr[m] != m {
				arr[i], arr[m] = arr[m], arr[i]
			} else {
				return m, true
			}
		} else {
			i++
		}
	}
	return -1, false
}

func isValidArray(arr []int) bool {
	if arr == nil || len(arr) < 2 {
		return false
	}
	n := len(arr)
	for _, e := range arr {
		if e < 0 || e >= n {
			return false
		}
	}
	return true
}
