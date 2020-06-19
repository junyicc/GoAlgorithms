package leetcode

// Given an unsorted integer array, find the smallest missing positive integer.
//
// Example 1:
//
// Input: [1,2,0]
// Output: 3
// Example 2:
//
// Input: [3,4,-1,1]
// Output: 2
// Example 3:
//
// Input: [7,8,9,11,12]
// Output: 1

// Note:
// Your algorithm should run in O(n) time and uses constant extra space.

func firstMissingPositive(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	// check min positive 1 exists
	exists1 := false
	for _, num := range nums {
		if num == 1 {
			exists1 = true
			break
		}
	}
	if !exists1 {
		return 1
	}

	// valid num: [1, n]
	// use 1 to replace invalid num:
	// 	- negative
	// 	- num gt n
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = 1
		}
	}

	// use index as hash value
	// traverse array
	// make it dualistic:
	// 	- get nums[i], make the index nums[i] negative only once
	// 	- use index 0 to represent n exists
	for i := 0; i < n; i++ {
		idx := nums[i]
		if idx == n {
			idx = 0
		}
		if idx >= 0 && nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	return n + 1
}
