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

// FindNumAppearOnce returns the number that appear once in the array where other numbers appear three times
func FindNumAppearOnce(arr []int) (int, bool) {
	if len(arr)%3 != 1 {
		return 0, false
	}

	if len(arr) == 1 {
		return arr[0], true
	}

	bitSum := [64]int{}
	for i := 0; i < len(arr); i++ {
		var bitMask int64 = 1
		for j := 63; j >= 0; j-- {
			b := int64(arr[i]) & bitMask
			if b != 0 {
				bitSum[j]++
			}
			bitMask = bitMask << 1
		}
	}

	res := 0
	for i := 0; i < 64; i++ {
		res = res << 1
		res += bitSum[i] % 3
	}
	return res, true
}
