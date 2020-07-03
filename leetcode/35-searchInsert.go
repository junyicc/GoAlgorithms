package leetcode

// Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

// search for the first position that greater than or equal to target

func searchInsert(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}

	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if nums[mi] < target {
			lo = mi + 1
		} else {
			if mi == 0 || nums[mi-1] < target {
				return mi
			} else {
				hi = mi - 1
			}
		}
	}
	if lo > hi {
		return len(nums)
	}
	return -1
}
