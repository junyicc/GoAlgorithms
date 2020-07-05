package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	if m < 1 || n < 1 {
		return 1
	}

	states := make([][]int, m)
	for i := 0; i < m; i++ {
		states[i] = make([]int, n)
	}

	// init states
	if obstacleGrid[0][0] == 0 {
		states[0][0] = 1
	} else {
		states[0][0] = 0
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			states[i][0] = min(states[i-1][0], 1)
		}
	}
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 0 {
			states[0][i] = min(states[0][i-1], 1)
		}
	}

	// update states
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				states[i][j] = states[i-1][j] + states[i][j-1]
			}
		}
	}
	return states[m-1][n-1]
}
