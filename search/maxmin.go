package search

func findMaxMin(nums []int, lo, hi int) (max int, min int) {
	// base case
	if hi-lo == 2 {
		if nums[lo] > nums[hi-1] {
			return nums[lo], nums[hi-1]
		} else {
			return nums[hi-1], nums[lo]
		}
	}

	if hi-lo < 2 {
		return nums[lo], nums[lo]
	}

	mi := lo + (hi-lo)>>1
	leftMax, leftMin := findMaxMin(nums, lo, mi)
	rightMax, rightMin := findMaxMin(nums, mi, hi)

	if leftMax > rightMax {
		max = leftMax
	} else {
		max = rightMax
	}
	if leftMin > rightMin {
		min = rightMin
	} else {
		min = leftMin
	}
	return max, min
}
