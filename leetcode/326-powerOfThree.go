package leetcode

import (
	"math"
)

func isPowerOfThree(n int) bool {
	return equal(math.Mod((math.Log10(float64(n))/math.Log10(3)), 1), 0)
}

func equal(a, b float64) bool {
	if a-b < 1e-10 {
		return true
	}
	return false
}
