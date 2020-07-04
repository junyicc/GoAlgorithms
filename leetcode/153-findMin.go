package leetcode

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
//
// Find the minimum element.
//
// You may assume no duplicate exists in the array.
//
// Example 1:
// Input: [3,4,5,1,2]
// Output: 1
//
// Example 2:
// Input: [4,5,6,7,0,1,2]
// Output: 0

func findMin(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	lo, hi := 0, len(nums)-1
	for hi-lo > 1 {
		mi := lo + (hi-lo)>>1
		if nums[mi] <= nums[hi] {
			hi = mi
		} else if nums[mi] >= nums[lo] {
			lo = mi
		}
	}
	if nums[hi] < nums[lo] {
		return nums[hi]
	}

	return nums[lo]
}
