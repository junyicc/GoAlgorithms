package interviews

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	for i, j := 0, n-1; i < m && j >= 0; {
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

// FindInMatrix where numbers are increasing both in column and row
func FindInMatrix(matrix [][]int, rows, cols int, e int) bool {
	if matrix == nil {
		return false
	}
	found := false
	for i, j := 0, cols-1; i < rows && j >= 0; {
		entry := matrix[i][j]
		if entry == e {
			found = true
			break
		} else if entry < e {
			i++
		} else if e < entry {
			j--
		}
	}
	return found
}
