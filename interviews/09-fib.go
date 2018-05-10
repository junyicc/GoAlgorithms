package interviews

// Fib returns fibonacci number
func Fib(n int) (int, bool) {
	if n < 0 {
		return 0, false
	}
	if n < 2 && n >= 0 {
		return n, true
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b, true
}
