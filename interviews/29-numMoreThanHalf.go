package interviews

// FindNumMoreThanHalf returns the number that appears more than half
func FindNumMoreThanHalf(arr []int) (int, bool) {
	if arr == nil || len(arr) < 1 {
		return 0, false
	}
	n := len(arr)

	num := arr[0]
	times := 1
	for i := 1; i < n; i++ {
		if times == 0 {
			times = 1
			num = arr[i]
		} else if arr[i] == num {
			times++
		} else {
			times--
		}
	}
	if times < 1 {
		return 0, false
	}

	return num, true
}
