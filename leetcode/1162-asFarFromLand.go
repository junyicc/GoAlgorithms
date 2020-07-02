package leetcode

import "math"

// Given an N x N grid containing only values 0 and 1, where 0 represents water and 1 represents land, find a water cell such that its distance to the nearest land cell is maximized and return the distance.
//
// The distance used in this problem is the Manhattan distance: the distance between two cells (x0, y0) and (x1, y1) is |x0 - x1| + |y0 - y1|.
//
// If no land or water exists in the grid, return -1.

func maxDistance(grid [][]int) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return -1
	}
	n, m := len(grid), len(grid[0])
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				dist[i][j] = math.MaxInt32
			}
		}
	}

	maxDist := -1
	// value from left, up
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				left, up := math.MaxInt32, math.MaxInt32
				if i-1 >= 0 {
					up = dist[i-1][j]
				}
				if j-1 >= 0 {
					left = dist[i][j-1]
				}
				dist[i][j] = min(dist[i][j], min(left, up)+1)
			}
		}
	}

	// value from right, down
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if grid[i][j] == 0 {
				right, down := math.MaxInt32, math.MaxInt32
				if i+1 < n {
					down = dist[i+1][j]
				}
				if j+1 < m {
					right = dist[i][j+1]
				}
				dist[i][j] = min(dist[i][j], min(right, down)+1)
				if dist[i][j] > maxDist {
					maxDist = dist[i][j]
				}
			}
		}
	}
	if maxDist == math.MaxInt32 {
		maxDist = -1
	}
	return maxDist
}
