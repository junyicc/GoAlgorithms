package interviews

// FindTwoNumsAppearOnce returns two numbers that appear once in the array where other numbers appear twice
func FindTwoNumsAppearOnce(arr []int, n1, n2 *int) bool {
	if arr == nil || len(arr) < 2 || n1 == nil || n2 == nil {
		return false
	}

	n := arr[0]
	for i := 1; i < len(arr); i++ {
		n ^= arr[i]
	}
	index := findFirstBitIs1(n)

	*n1, *n2 = 0, 0
	for i := 0; i < len(arr); i++ {
		if isBit1(arr[i], index) {
			*n1 ^= arr[i]
		} else {
			*n2 ^= arr[i]
		}
	}

	return true
}

func findFirstBitIs1(n int) uint {
	var index uint
	m := 1
	for index < 64 && n&m == 0 {
		m = m << 1
		index++
	}
	return index
}

func isBit1(n int, index uint) bool {
	n = n >> index
	return n&1 == 1
}
