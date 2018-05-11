package interviews

// NumOf1 returns number of 1 in the bits of n
func NumOf1(n int) int {
	cnt := 0
	for n != 0 {
		n = n & (n - 1)
		cnt++
	}
	return cnt
}

// IsPowerOf2 return true if n is the power of 2
func IsPowerOf2(n int) bool {
	if n > 0 {
		return n&(n-1) == 0
	}
	return false
}

// NumOfChangingBits returns number of bits that are needed to be changed from m to n
func NumOfChangingBits(m, n int) int {
	diff := m ^ n
	return NumOf1(diff)
}

// IsOdd return true if n is an odd number
func IsOdd(n int) bool {
	if n <= 0 {
		return false
	}
	return uint(n)&1 == 1
}
