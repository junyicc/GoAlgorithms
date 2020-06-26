package dynamic_programming

import "math"

// the path sum of triangle[i][j] may be the result of min{sums[i-1][j-1], sums[i-1][j], sums[i-1][j+1]}
func minPathSumOfTriangle(triangle [][]int) int {
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
			states[i][j] = triangle[i][j] + validMinOfSlice(states[i-1], j-1, j, j+1)
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

func validMinOfSlice(nums []int, i, j, k int) int {
	n := len(nums)
	a, b, c := math.MaxInt32, math.MaxInt32, math.MaxInt32
	if i >= 0 && i < n {
		a = nums[i]
	}
	if j >= 0 && j < n {
		b = nums[j]
	}
	if k >= 0 && k < n {
		c = nums[k]
	}

	return min(a, min(b, c))
}
