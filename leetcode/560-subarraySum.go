package leetcode

// Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.
//
// Example 1:
// Input:nums = [1,1,1], k = 2
// Output: 2

func subarraySumWithPresum(nums []int, k int) int {
	if len(nums) < 1 {
		return 0
	}

	n := len(nums)

	// presum[i]: the sum of 0~(i-1) nums
	presum := make([]int, n+1)
	for i := 0; i < n; i++ {
		presum[i+1] = presum[i] + nums[i]
	}
	var cnt int
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// the sum of j~(i-1) nums
			if presum[i]-presum[j] == k {
				cnt++
			}
		}
	}
	return cnt
}

func subarraySum(nums []int, k int) int {
	if len(nums) < 1 {
		return 0
	}

	n := len(nums)

	// presum[i]: the cnt of i
	presum := make(map[int]int)
	presum[0] = 1

	var cnt int
	var cum int
	for i := 0; i < n; i++ {
		cum += nums[i]
		subSum := cum - k
		// exists such sum of subarray
		if c, ok := presum[subSum]; ok {
			cnt += c
		}

		presum[cum]++
	}
	return cnt
}
