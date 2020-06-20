package leetcode

// Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

// Follow up:
// Could you solve it in linear time?

// Example:

// Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
// Output: [3,3,5,5,6,7]
// Explanation:

// Window position                Max
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < 1 || k < 1 {
		return nil
	}

	res := make([]int, 0, len(nums)-k+1)

	// init win for the first k elements
	win := make([]int, 0, k)
	for i := 0; i < k; i++ {
		for len(win) > 0 && nums[i] >= nums[win[len(win)-1]] {
			win = win[:len(win)-1]
		}
		win = append(win, i)
	}
	// the first max num
	res = append(res, nums[win[0]])

	// process the remaining elements
	for i := k; i < len(nums); i++ {
		// win sliding
		if win[0] == i-k {
			win = win[1:]
		}
		// clean win
		for len(win) > 0 && nums[i] >= nums[win[len(win)-1]] {
			win = win[:len(win)-1]
		}

		// win expanding
		win = append(win, i)

		// the current max num
		res = append(res, nums[win[0]])
	}

	return res
}
