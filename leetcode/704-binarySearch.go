package leetcode

// Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}

	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if nums[mi] == target {
			return mi
		} else if nums[mi] < target {
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	return -1
}
