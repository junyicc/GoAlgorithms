package leetcode

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
//
// You are given a target value to search. If found in the array return its index, otherwise return -1.
//
// You may assume no duplicate exists in the array.
//
// Your algorithm's runtime complexity must be in the order of O(log n).
//
// Example 1:
// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
//
// Example 2:
// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1

func searchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}

	lo, hi := 0, len(nums)-1
	for hi-lo > 1 {
		mi := lo + (hi-lo)>>1
		if nums[mi] == target {
			return mi
		}

		// there are 4 possibilities
		if nums[lo] < nums[mi] {
			// left part
			if nums[lo] <= target && nums[mi] >= target {
				// left part and less than nums[mi]
				hi = mi
			} else {
				lo = mi
			}
		} else if nums[mi] < nums[hi] {
			// right part
			if nums[hi] >= target && nums[mi] <= target {
				// right part, and greater than nums[mi]
				lo = mi
			} else {
				hi = mi
			}
		}
	}

	if nums[lo] == target {
		return lo
	} else if nums[hi] == target {
		return hi
	}
	return -1
}
