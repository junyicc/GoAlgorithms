package leetcode

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}
	cnt := 0
	for i := 5; n/i > 0; i *= 5 {
		cnt += n / i
	}
	return cnt
}

func trailingZeroesRecur(n int) int {
	if n == 0 {
		return 0
	}
	return n/5 + trailingZeroesRecur(n/5)
}
