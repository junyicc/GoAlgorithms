package leetcode

func canJump(nums []int) bool {
	if len(nums) < 1 {
		return false
	}
	n := len(nums)
	states := make([]bool, n)

	// init state
	states[0] = true
	// update states
	for i := 1; i < n; i++ {
		// for every current i, calculate states from 0
		for j := 0; j < i; j++ {
			if states[j] && nums[j]+j >= i {
				states[i] = true
			}
		}
	}

	return states[n-1]
}
