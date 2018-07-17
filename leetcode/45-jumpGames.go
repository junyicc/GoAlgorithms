package leetcode

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
