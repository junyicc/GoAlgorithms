package leetcode

// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
//
// Example:
// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.

// Follow up:
// If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	n := len(nums)
	curSum := nums[0]
	max := nums[0]
	for i := 1; i < n; i++ {
		if curSum+nums[i] <= nums[i] {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		if curSum > max {
			max = curSum
		}
	}
	return max
}
