package leetcode

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
// (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
//
// Find the minimum element.
//
// The array may contain duplicates.

func findMinII(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	lo, hi := 0, len(nums)-1
	for hi-lo > 1 {
		// de-duplication
		for lo < hi && nums[hi] == nums[hi-1] {
			hi--
		}
		for lo < hi && nums[lo] == nums[lo+1] {
			lo++
		}

		mi := lo + (hi-lo)>>1
		if nums[mi] <= nums[hi] {
			hi = mi
		} else if nums[lo] <= nums[mi] {
			lo = mi
		}
	}
	if nums[hi] < nums[lo] {
		return nums[hi]
	}

	return nums[lo]
}
