package leetcode

// Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
//
// - Integers in each row are sorted from left to right.
// - The first integer of each row is greater than the last integer of the previous row.

// Example 1:
// Input:
// matrix = [
// [1,   3,  5,  7],
// [10, 11, 16, 20],
// [23, 30, 34, 50]
// ]
// target = 3
// Output: true

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}

	n, m := len(matrix), len(matrix[0])

	for i, j := 0, m-1; i < n && j >= 0; {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
