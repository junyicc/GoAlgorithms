package leetcode

// Given a 2-dimensional grid of integers, each value in the grid represents the color of the grid square at that location.
//
// Two squares belong to the same connected component if and only if they have the same color and are next to each other in any of the 4 directions.
//
// The border of a connected component is all the squares in the connected component that are either 4-directionally adjacent to a square not in the component, or on the boundary of the grid (the first or last row or column).
//
// Given a square at location (r0, c0) in the grid and a color, color the border of the connected component of that square with the given color, and return the final grid.
//
// Example 1:
// Input: grid = [[1,1],[1,2]], r0 = 0, c0 = 0, color = 3
// Output: [[3, 3], [3, 2]]
//
// Example 2:
// Input: grid = [[1,2,2],[2,3,2]], r0 = 0, c0 = 1, color = 3
// Output: [[1, 3, 3], [2, 3, 3]]
//
// Example 3:
// Input: grid = [[1,1,1],[1,1,1],[1,1,1]], r0 = 1, c0 = 1, color = 2
// Output: [[2, 2, 2], [2, 1, 2], [2, 2, 2]]

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return grid
	}

	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	oldColor := grid[r0][c0]
	colorBorderDFS(grid, r0, c0, oldColor, color, visited)

	return grid
}

func colorBorderDFS(grid [][]int, r, c int, oldColor, newColor int, visited [][]bool) int {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return 0
	}

	if visited[r][c] {
		return 1
	}

	if grid[r][c] != oldColor {
		return 0
	}

	// grid[r][c] == oldColor
	visited[r][c] = true

	surround := colorBorderDFS(grid, r-1, c, oldColor, newColor, visited) +
		colorBorderDFS(grid, r+1, c, oldColor, newColor, visited) +
		colorBorderDFS(grid, r, c-1, oldColor, newColor, visited) +
		colorBorderDFS(grid, r, c+1, oldColor, newColor, visited)

	if surround < 4 {
		grid[r][c] = newColor
	}
	return 1
}
