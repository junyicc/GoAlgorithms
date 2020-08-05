package interviews

func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		c := a + b
		a = b
		b = c % (1e9 + 7)
	}

	return b
}
