package leetcode

import "sort"

// Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
// Each number in candidates may only be used once in the combination.
//
// Note:
//
// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
// Example 1:
//
// Input: candidates = [10,1,2,7,6,1,5], target = 8,
// A solution set is:
// [
// [1, 7],
// [1, 2, 5],
// [2, 6],
// [1, 1, 6]
// ]
// Example 2:
//
// Input: candidates = [2,5,2,1,2], target = 5,
// A solution set is:
// [
//   [1,2,2],
//   [5]
// ]

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) < 1 {
		return nil
	}
	sort.Ints(candidates)

	var res [][]int
	track := make([]int, 0)
	backtrackCombinationSum2(candidates, 0, 0, target, &track, &res)
	return res
}

func backtrackCombinationSum2(candidates []int, lo, curSum, target int, track *[]int, res *[][]int) {
	if curSum == target {
		*res = append(*res, append((*track)[:0:0], *track...))
		return
	}

	for i := lo; i < len(candidates); i++ {
		if curSum+candidates[i] > target {
			break
		}

		if i > lo && candidates[i] == candidates[i-1] {
			continue
		}

		// make choice
		*track = append(*track, candidates[i])
		curSum += candidates[i]

		// backtrack
		backtrackCombinationSum2(candidates, i+1, curSum, target, track, res)

		// cancel choice
		*track = (*track)[:len(*track)-1]
		curSum -= candidates[i]
	}
}
