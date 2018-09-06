package leetcode

import (
	"math"
)

/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:
Input: 123
Output: 321

Example 2:
Input: -123
Output: -321

Example 3:
Input: 120
Output: 21
*/
func reverse(o int) int {
	negative := false
	if o < 0 {
		o = 0 - o
		negative = true
	}
	d := 0
	for o > 0 {
		num := o % 10
		d = d*10 + num
		o /= 10
	}
	if negative {
		d = 0 - d
	}
	if d > math.MaxInt32 || d < math.MinInt32 {
		return 0
	}

	return d
}
