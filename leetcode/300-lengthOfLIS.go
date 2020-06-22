package leetcode

func lengthOfLIS(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	n := len(nums)
	states := make([]int, n)
	for i := 0; i < n; i++ {
		states[i] = 1
	}
	maxLen := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if states[j]+1 > states[i] {
					states[i] = states[j] + 1
				}
			}
		}
		if states[i] > maxLen {
			maxLen = states[i]
		}
	}
	return maxLen
}
