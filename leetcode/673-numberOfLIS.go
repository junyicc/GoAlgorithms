package leetcode

// Given an unsorted array of integers, find the number of longest increasing subsequence.
//
// Example 1:
// Input: [1,3,5,4,7]
// Output: 2
// Explanation: The two longest increasing subsequence are [1, 3, 4, 7] and [1, 3, 5, 7].
//
// Example 2:
// Input: [2,2,2,2,2]
// Output: 5
// Explanation: The length of longest continuous increasing subsequence is 1, and there are 5 subsequences' length is 1, so output 5.

func findNumberOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	n := len(nums)
	states := make([]int, n)
	for i := 0; i < n; i++ {
		states[i] = 1
	}
	count := make([]int, n)
	for i := 0; i < n; i++ {
		count[i] = 1
	}

	// calculate states
	maxLen := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if states[j]+1 > states[i] {
					states[i] = states[j] + 1
					count[i] = count[j]
				} else if states[j]+1 == states[i] {
					count[i] += count[j]
				}
			}
		}
		if states[i] > maxLen {
			maxLen = states[i]
		}
	}
	// find the max len of LIS,
	// and sum the count of combinations
	cnt := 0
	for i := 0; i < n; i++ {
		if states[i] == maxLen {
			cnt += count[i]
		}
	}

	return cnt
}
