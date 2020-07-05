package leetcode

import "math"

// Given an array of non-negative integers, you are initially positioned at the first index of the array.
// Each element in the array represents your maximum jump length at that position.
// Your goal is to reach the last index in the minimum number of jumps.

// Jump games
func Jump(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}

	return jumpnext(nums, len(nums))
}

func jumpnext(nums []int, n int) int {
	var i, level, nextMax, curMax int
	for i <= curMax {
		level++
		for ; i <= curMax; i++ {
			// update the max reach of next level
			if nums[i]+i > nextMax {
				nextMax = nums[i] + i
			}
			if nextMax >= n-1 {
				return level
			}
		}
		curMax = nextMax
	}
	return -1
}

func jump(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	n := len(nums)

	states := make([]int, n)
	// init states
	for i := 1; i < n; i++ {
		states[i] = math.MaxInt32
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if states[j] >= 0 && nums[j]+j >= i {
				if states[j]+1 < states[i] {
					states[i] = states[j] + 1
				}
			}
		}
	}
	return states[n-1]
}
