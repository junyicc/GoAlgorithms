package interviews

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
