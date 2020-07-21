package leetcode

// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.
//
// Example:
// Input: [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6

func trap(height []int) int {
	if len(height) < 1 {
		return 0
	}

	n := len(height)
	leftMax := height[0]
	rightMax := height[n-1]
	lo, hi := 0, n-1
	var water int
	for lo <= hi {
		// update left max and right max
		leftMax = max(leftMax, height[lo])
		rightMax = max(rightMax, height[hi])

		// calculate water
		if leftMax < rightMax {
			water += leftMax - height[lo]
			lo++
		} else {
			water += rightMax - height[hi]
			hi--
		}
	}
	return water
}

func trapSimple(height []int) int {
	if len(height) < 1 {
		return 0
	}

	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	water := 0
	for i := 0; i < n; i++ {
		water += min(leftMax[i], rightMax[i]) - height[i]
	}
	return water
}
