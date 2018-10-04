package leetcode

func hammingWeight(n uint32) int {
	cnt := 0
	for n != 0 {
		n = n & (n - 1)
		cnt++
	}
	return cnt
}
