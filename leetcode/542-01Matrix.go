package leetcode

import "math"

// Given a matrix consists of 0 and 1, find the distance of the nearest 0 for each cell.
//
// The distance between two adjacent cells is 1.
//
// Example 1:
// Input:
// [[0,0,0],
// [0,1,0],
// [0,0,0]]
//
// Output:
// [[0,0,0],
// [0,1,0],
// [0,0,0]]
//
// Example 2:
// Input:
// [[0,0,0],
// [0,1,0],
// [1,1,1]]
//
// Output:
// [[0,0,0],
// [0,1,0],
// [1,2,1]]

// 「Tree 的 BFS」是「单源 BFS」
// 「图 的 BFS」 是「多源 BFS」

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return nil
	}

	n, m := len(matrix), len(matrix[0])
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
				dist[i][j] = math.MaxInt32
			}
		}
	}
	// from up, left
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 1 {
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
	// from down, right
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if matrix[i][j] == 1 {
				right, down := math.MaxInt32, math.MaxInt32
				if i+1 < n {
					down = dist[i+1][j]
				}
				if j+1 < m {
					right = dist[i][j+1]
				}
				dist[i][j] = min(dist[i][j], min(down, right)+1)
			}
		}
	}
	return dist
}
