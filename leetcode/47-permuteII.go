package leetcode

import "sort"

// Given a collection of numbers that might contain duplicates, return all possible unique permutations.
//
// Example:
// Input: [1,1,2]
// Output:
// [
// [1,1,2],
// [1,2,1],
// [2,1,1]
// ]

func permuteUnique(nums []int) [][]int {
	if len(nums) < 1 {
		return nil
	}

	sort.Ints(nums)

	var res [][]int
	f := func(track []int) {
		tmp := make([]int, 0, len(track))
		for _, i := range track {
			tmp = append(tmp, nums[i])
		}
		res = append(res, tmp)
	}
	track := make([]int, 0, len(nums))
	used := make(map[int]bool)
	permuteUniqueWithBacktracking(nums, track, used, f)
	return res
}

func permuteUniqueWithBacktracking(nums []int, track []int, used map[int]bool, f func([]int)) {
	if len(track) == len(nums) {
		clone := append(track[:0:0], track...)
		f(clone)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		// pruning
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		// make choice
		track = append(track, i)
		used[i] = true

		permuteUniqueWithBacktracking(nums, track, used, f)

		track = track[:len(track)-1]
		used[i] = false
	}
}
