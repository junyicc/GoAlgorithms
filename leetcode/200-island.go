package leetcode

// Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
//
// Example 1:
// Input:
// 11110
// 11010
// 11000
// 00000
// Output: 1
//
// Example 2:
// Input:
// 11000
// 11000
// 00100
// 00011
// Output: 3

func numIslands(grid [][]byte) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])
	// init states
	states := make([][]byte, m)
	for i := 0; i < m; i++ {
		states[i] = make([]byte, n)
	}
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				landNum := validIsland(grid, states, i, j)
				if landNum > 0 {
					cnt++
				}
			}
		}
	}
	return cnt
}

func validIsland(grid [][]byte, states [][]byte, i, j int) int {
	if grid[i][j] == '0' {
		return 0
	}

	var landNum int
	if states[i][j] == 0 {
		states[i][j] = 1
		landNum++
	}

	// update other states
	// left
	if j-1 >= 0 {
		if grid[i][j-1] == '1' && states[i][j-1] == 0 {
			states[i][j-1] = 1
			landNum += validIsland(grid, states, i, j-1)
		}
	}
	// right
	if j+1 < len(grid[0]) {
		if grid[i][j+1] == '1' && states[i][j+1] == 0 {
			states[i][j+1] = 1
			landNum += validIsland(grid, states, i, j+1)
		}
	}
	// up
	if i-1 >= 0 {
		if grid[i-1][j] == '1' && states[i-1][j] == 0 {
			states[i-1][j] = 1
			landNum += validIsland(grid, states, i-1, j)
		}
	}
	// down
	if i+1 < len(grid) {
		if grid[i+1][j] == '1' && states[i+1][j] == 0 {
			states[i+1][j] = 1
			landNum += validIsland(grid, states, i+1, j)
		}
	}

	return landNum
}
