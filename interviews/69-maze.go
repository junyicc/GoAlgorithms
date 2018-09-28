package interviews

import "math"

// FindShortestPath returns the shortest path from (0,0) to (n-1, m-1)
func FindShortestPath(matrix [][]int, n, m int) int {
	if matrix == nil || n < 1 || m < 1 {
		return 0
	}
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	steps := leastSteps(matrix, n, m, 0, 0, visited)
	return steps
}

func leastSteps(matrix [][]int, nrows, ncols, row, col int, visited [][]bool) int {
	if row == nrows-1 && col == ncols-1 {
		return matrix[row][col]
	}
	steps := math.MaxInt32
	if row >= 0 && row < nrows && col >= 0 && col < ncols && matrix[row][col] != -1 && !visited[row][col] {
		visited[row][col] = true
		steps = matrix[row][col] + min([]int{
			leastSteps(matrix, nrows, ncols, row-1, col, visited),
			leastSteps(matrix, nrows, ncols, row, col-1, visited),
			leastSteps(matrix, nrows, ncols, row+1, col, visited),
			leastSteps(matrix, nrows, ncols, row, col+1, visited),
		})
		visited[row][col] = false
	}
	return steps
}

func min(nums []int) int {
	minNum := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}
	return minNum
}
