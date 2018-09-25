package math

// IsEven return true if n is an odd number
func IsEven(n int) bool {
	if n < 0 {
		return false
	}
	return uint(n)&1 == 0
}
