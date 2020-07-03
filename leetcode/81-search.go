package leetcode

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
// (i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).
//
// You are given a target value to search. If found in the array return true, otherwise return false.

func searchInRotatedSortedArrayII(nums []int, target int) bool {
	if len(nums) < 1 {
		return false
	}
	n := len(nums)
	lo, hi := 0, n-1
	for hi-lo > 1 {
		for lo < hi && nums[lo] == nums[lo+1] {
			lo++
		}
		for lo < hi && nums[hi] == nums[hi-1] {
			hi--
		}

		mi := lo + (hi-lo)>>1
		if nums[lo] <= nums[mi] {
			if nums[lo] <= target && nums[mi] >= target {
				hi = mi
			} else {
				lo = mi
			}
		} else if nums[mi] <= nums[hi] {
			if nums[hi] >= target && nums[mi] <= target {
				lo = mi
			} else {
				hi = mi
			}
		}
	}
	if nums[lo] == target || nums[hi] == target {
		return true
	}

	return false
}
