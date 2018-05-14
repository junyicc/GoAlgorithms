package interviews

// VisitMatrixClockwisely uses f to visit matrix clockwisely
func VisitMatrixClockwisely(matrix [][]int, rows, cols int, f func(int)) {
	if matrix == nil || rows < 1 || cols < 1 {
		return
	}

	start := 0
	for rows > 2*start && cols > 2*start {
		visitMatrixClockwisely(matrix, rows, cols, start, f)
		start++
	}
}

func visitMatrixClockwisely(matrix [][]int, rows, cols, start int, f func(int)) {
	endX := cols - 1 - start
	endY := rows - 1 - start
	// visit entry from left to right: necessary
	for i := start; i <= endX; i++ {
		f(matrix[start][i])
	}
	// visit entry from top to bottom
	if endY > start {
		for i := start + 1; i <= endY; i++ {
			f(matrix[i][endX])
		}
	}
	// visit entry from right to left
	if endX > start && endY > start {
		for i := endX - 1; i >= start; i-- {
			f(matrix[endY][i])
		}
	}
	// visit entry from bottom to top
	if endY-start > 1 && endX > start {
		for i := endY - 1; i > start; i-- {
			f(matrix[i][start])
		}
	}
}
