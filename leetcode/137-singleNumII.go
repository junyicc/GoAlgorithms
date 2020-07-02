package leetcode

// Given a non-empty array of integers, every element appears three times except for one, which appears exactly once. Find that single one.
//
// Example 1:
// Input: [2,2,3,2]
// Output: 3
//
// Example 2:
// Input: [0,1,0,1,0,1,99]
// Output: 99

func singleNumberII(nums []int) int {
	// invalid nums length
	if len(nums)%3 != 1 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var res int
	for i := uint(0); i < 64; i++ {
		// counting bit counts
		var bitCnt uint
		for _, n := range nums {
			if n&(1<<i) != 0 {
				bitCnt++
			}
		}
		// restore num appearing once
		res |= int((bitCnt % 3) << i)
	}
	return res
}
