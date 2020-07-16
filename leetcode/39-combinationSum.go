package leetcode

import "sort"

// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
// The same repeated number may be chosen from candidates unlimited number of times.
//
// Note:
//
// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.

// Example 1:
// Input: candidates = [2,3,6,7], target = 7,
// A solution set is:
// [
// [7],
// [2,2,3]
// ]

// Example 2:
// Input: candidates = [2,3,5], target = 8,
// A solution set is:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) < 1 {
		return nil
	}

	sort.Ints(candidates)

	var res [][]int
	track := make([]int, 0)
	backtrackCombinationSum(candidates, 0, 0, target, &track, &res)

	return res
}

func backtrackCombinationSum(candidates []int, lo, curSum, target int, track *[]int, res *[][]int) {
	if curSum == target {
		*res = append(*res, append((*track)[:0:0], *track...))
		return
	}

	for i := lo; i < len(candidates); i++ {
		// pruning
		if curSum+candidates[i] > target {
			break
		}
		// make choice
		curSum += candidates[i]
		*track = append(*track, candidates[i])
		// backtrack
		backtrackCombinationSum(candidates, i, curSum, target, track, res)
		// cancel choice
		curSum -= candidates[i]
		*track = (*track)[:len(*track)-1]
	}
}
