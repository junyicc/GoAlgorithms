package leetcode

// Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
//
// Example:
//
// Input: n = 4, k = 2
// Output:
// [
// [2,4],
// [3,4],
// [2,3],
// [1,2],
// [1,3],
// [1,4],
// ]

func combine(n int, k int) [][]int {
	if n < 1 || k < 1 {
		return nil
	}

	var res [][]int
	track := make([]int, 0)

	backtrackCombine(1, k, n, &track, &res)
	return res
}

func backtrackCombine(lo, k, n int, track *[]int, res *[][]int) {
	if len(*track) == k {
		*res = append(*res, append((*track)[:0:0], *track...))
	}

	for i := lo; i <= n; i++ {
		// make choice
		*track = append(*track, i)
		// backtrack
		backtrackCombine(i+1, k, n, track, res)
		// cancel choice
		*track = (*track)[:len(*track)-1]
	}
}
