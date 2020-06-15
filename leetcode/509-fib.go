package leetcode

// F(0) = 0,   F(1) = 1
// F(N) = F(N - 1) + F(N - 2), for N > 1.

func fib(N int) int {
	if N <= 0 {
		return 0
	}
	if N == 1 {
		return 1
	}

	a, b := 0, 1
	for i := 1; i < N; i++ {
		a, b = b, a+b
	}
	return b
}
