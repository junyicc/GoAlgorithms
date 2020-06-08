package dynamic_programming

func ShortestPathLengthInMatrix(matrix [][]int) int {
	if len(matrix) < 1 {
		return 0
	}
	return shortestPathLengthInMatrix(matrix)
}

func shortestPathLengthInMatrix(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, m)
		copy(states[i], matrix[i])
	}

	// init state
	for j := 1; j < m; j++ {
		states[0][j] = states[0][j] + states[0][j-1]
	}
	for i := 1; i < n; i++ {
		states[i][0] = states[i-1][0] + states[i][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			states[i][j] = min(states[i-1][j]+states[i][j], states[i][j-1]+states[i][j])
		}
	}
	return states[n-1][m-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
