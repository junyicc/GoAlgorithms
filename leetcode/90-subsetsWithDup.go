package leetcode

import "sort"

// Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).
//
// Note: The solution set must not contain duplicate subsets.
//
// Example:
//
// Input: [1,2,2]
// Output:
// [
// [2],
// [1],
// [1,2,2],
// [2,2],
// [1,2],
// []
// ]

func subsetsWithDup(nums []int) [][]int {
	if len(nums) < 1 {
		return nil
	}

	var res [][]int
	track := make([]int, 0)

	sort.Ints(nums)
	backtrackSubsetsWithDup(nums, 0, &track, &res)
	return res
}

func backtrackSubsetsWithDup(nums []int, lo int, track *[]int, res *[][]int) {
	*res = append(*res, append((*track)[:0:0], *track...))

	for i := lo; i < len(nums); i++ {
		// i>lo && nums[i] == nums[i-1]: the same element in the next round
		if i > lo && nums[i] == nums[i-1] {
			continue
		}
		// make choice
		*track = append(*track, nums[i])
		// backtrack
		backtrackSubsetsWithDup(nums, i+1, track, res)
		// cancel choice
		*track = (*track)[:len(*track)-1]
	}
}
