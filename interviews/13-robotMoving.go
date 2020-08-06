package interviews

// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
//
// 示例 1：
// 输入：m = 2, n = 3, k = 1
// 输出：3
//
// 示例 2：
// 输入：m = 3, n = 1, k = 0
// 输出：1
//
// 提示：
// 1 <= n,m <= 100
// 0 <= k<= 20

func movingCount(m int, n int, k int) int {
	if m < 1 || n < 1 {
		return 0
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	return movingCountCore(m, n, k, 0, 0, visited)
}

func movingCountCore(m, n int, k int, i, j int, visited [][]bool) int {
	if i < 0 || i >= m || j < 0 || j >= n {
		return 0
	}

	if visited[i][j] {
		return 0
	}

	if !canMoveIn(i, j, k) {
		return 0
	}

	visited[i][j] = true

	return 1 +
		movingCountCore(m, n, k, i-1, j, visited) +
		movingCountCore(m, n, k, i+1, j, visited) +
		movingCountCore(m, n, k, i, j-1, visited) +
		movingCountCore(m, n, k, i, j+1, visited)
}

func canMoveIn(i, j int, k int) bool {
	var s int
	for i != 0 {
		s += i % 10
		i = i / 10
	}
	for j != 0 {
		s += j % 10
		j = j / 10
	}
	return s <= k
}
