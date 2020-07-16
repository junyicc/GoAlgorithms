package leetcode

// Given a set of distinct integers, nums, return all possible subsets (the power set).
//
// Note: The solution set must not contain duplicate subsets.
//
// Example:
//
// Input: nums = [1,2,3]
// Output:
// [
// [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

func subsets(nums []int) [][]int {
	if len(nums) < 1 {
		return nil
	}

	var res [][]int
	track := make([]int, 0)
	backtrackSubsets(nums, 0, &track, &res)
	return res
}

func backtrackSubsets(nums []int, lo int, track *[]int, res *[][]int) {
	*res = append(*res, append((*track)[:0:0], *track...))

	for i := lo; i < len(nums); i++ {
		// make choice
		*track = append(*track, nums[i])

		// backtrack
		backtrackSubsets(nums, i+1, track, res)

		*track = (*track)[:len(*track)-1]
	}
}
