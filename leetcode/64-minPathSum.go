package leetcode

func minPathSum(grid [][]int) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	states := make([][]int, m)
	for i := 0; i < m; i++ {
		states[i] = make([]int, n)
		copy(states[i], grid[i])
	}

	for i := 1; i < n; i++ {
		states[0][i] = states[0][i] + states[0][i-1]
	}
	for i := 1; i < m; i++ {
		states[i][0] = states[i][0] + states[i-1][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			states[i][j] = states[i][j] + min(states[i-1][j], states[i][j-1])
		}
	}

	return states[m-1][n-1]
}
