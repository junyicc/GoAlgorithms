package math

// IsOdd return true if n is an odd number
func IsOdd(n int) bool {
	if n <= 0 {
		return false
	}
	return uint(n)&1 == 1
}
