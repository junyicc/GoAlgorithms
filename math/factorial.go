package math

// Factorial method
func Factorial(n int) int {
	if n < 0 {
		panic("Factorial: n must be ge 0")
	}
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
