package leetcode

// Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.
//
// Example 1:
// Input: [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.
//
// Example 2:
// Input: [-2,0,-1]
// Output: 0
// Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

func maxProduct(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	n := len(nums)
	maxProd := nums[0]
	minProd := nums[0]
	res := nums[0]
	for i := 1; i < n; i++ {
		curMaxProd := maxProd * nums[i]
		curMinProd := minProd * nums[i]

		maxProd = max(nums[i], max(curMaxProd, curMinProd))
		minProd = min(nums[i], min(curMaxProd, curMinProd))

		res = max(res, maxProd)
	}

	return res
}
