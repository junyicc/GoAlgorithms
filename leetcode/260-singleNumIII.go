package leetcode

// Given an array of numbers nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once.
//
// Example:
// Input:  [1,2,1,3,2,5]
// Output: [3,5]

func singleNumberIII(nums []int) []int {
	if len(nums) < 1 {
		return nil
	}

	var diff int
	for _, n := range nums {
		diff ^= n
	}

	// find the first bit 1 from low position of diff
	diff = (diff & (diff - 1)) ^ diff
	res := make([]int, 2)
	for _, n := range nums {
		if n&diff == 0 {
			res[0] ^= n
		} else {
			res[1] ^= n
		}
	}
	return res
}
