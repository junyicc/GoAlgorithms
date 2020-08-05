package interviews

func fib(n int) int {
	if n < 2 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		c := a + b
		a = b
		b = c % (1e9 + 7)
	}
	return b
}
