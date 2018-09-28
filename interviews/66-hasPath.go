package interviews

// HasPath returns true if s can be found in matrix
func HasPath(matrix [][]byte, nrows, ncols int, s string) bool {
	if matrix == nil || nrows < 1 || ncols < 1 || s == "" {
		return false
	}
	visited := make([][]bool, nrows)
	for i := 0; i < nrows; i++ {
		visited[i] = make([]bool, ncols)
	}

	for i := 0; i < nrows; i++ {
		for j := 0; j < ncols; j++ {
			if hasPath(matrix, nrows, ncols, i, j, s, 0, visited) {
				return true
			}
		}
	}
	return false
}

func hasPath(matrix [][]byte, nrows, ncols, row, col int, s string, strIdx int, visited [][]bool) bool {
	if strIdx == len(s) {
		return true
	}
	has := false
	if row >= 0 && row < nrows && col >= 0 && col < ncols && matrix[row][col] == s[strIdx] && !visited[row][col] {
		strIdx++
		visited[row][col] = true

		has = hasPath(matrix, nrows, ncols, row-1, col, s, strIdx, visited) ||
			hasPath(matrix, nrows, ncols, row, col-1, s, strIdx, visited) ||
			hasPath(matrix, nrows, ncols, row+1, col, s, strIdx, visited) ||
			hasPath(matrix, nrows, ncols, row, col+1, s, strIdx, visited)

		if !has {
			strIdx--
			visited[row][col] = false
		}
	}

	return has
}
