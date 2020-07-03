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
	min := nums[lo]
	for nums[hi] <= nums[lo] {
		if hi-lo == 1 {
			min = nums[hi]
			break
		}

		mi := lo + (hi-lo)>>1

		// handle duplication
		if nums[lo] == nums[mi] && nums[mi] == nums[hi] {
			return minOfSlice(nums)
		}

		if nums[mi] >= nums[lo] {
			lo = mi
		} else if nums[mi] <= nums[hi] {
			hi = mi
		}
	}

	return min
}

func minOfSlice(nums []int) int {
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
