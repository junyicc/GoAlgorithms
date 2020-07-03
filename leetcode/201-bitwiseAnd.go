package leetcode

// Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.
//
// Example 1:
// Input: [5,7]
// Output: 4
//
// Example 2:
// Input: [0,1]
// Output: 0

func rangeBitwiseAnd(m int, n int) int {
	// count the different bits from low position
	var cnt int
	for m != n {
		cnt++

		m >>= 1
		n >>= 1
	}
	// set different bits of origin m as 0
	return m << cnt
}

// Simple solution may cause timeout if the range is large
func rangeBitwiseAndSimple(m int, n int) int {
	res := m
	for i := m + 1; i <= n; i++ {
		res &= i
	}
	return res
}
