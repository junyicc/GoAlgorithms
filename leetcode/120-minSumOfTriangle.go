package leetcode

import "math"

// Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.
//
// For example, given the following triangle
//
// [
// [2],
// [3,4],
// [6,5,7],
// [4,1,8,3]
// ]
// The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

func minimumTotal(triangle [][]int) int {
	if len(triangle) < 1 || len(triangle[0]) < 1 {
		return 0
	}
	n := len(triangle)
	states := make([][]int, n)
	// init state
	states[0] = make([]int, len(triangle[0]))
	states[0][0] = triangle[0][0]
	// update states
	for i := 1; i < n; i++ {
		m := len(triangle[i])
		states[i] = make([]int, m)
		for j := 0; j < m; j++ {
			states[i][j] = triangle[i][j] + validMinOfSlice(states[i-1], j-1, j)
		}
	}

	minSum := states[n-1][0]
	m := len(states[n-1])
	for j := 1; j < m; j++ {
		if states[n-1][j] < minSum {
			minSum = states[n-1][j]
		}
	}
	return minSum
}

func validMinOfSlice(nums []int, i, j int) int {
	n := len(nums)
	a, b := math.MaxInt32, math.MaxInt32
	if i >= 0 && i < n {
		a = nums[i]
	}
	if j >= 0 && j < n {
		b = nums[j]
	}

	return min(a, b)
}
