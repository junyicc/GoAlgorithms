package leetcode

// An image is represented by a 2-D array of integers, each integer representing the pixel value of the image (from 0 to 65535).
//
// Given a coordinate (sr, sc) representing the starting pixel (row and column) of the flood fill, and a pixel value newColor, "flood fill" the image.
//
// To perform a "flood fill", consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color as the starting pixel), and so on. Replace the color of all of the aforementioned pixels with the newColor.
//
// At the end, return the modified image.
//
// Example 1:
// Input:
// image = [[1,1,1],[1,1,0],[1,0,1]]
// sr = 1, sc = 1, newColor = 2
// Output: [[2,2,2],[2,2,0],[2,0,1]]
// Explanation:
// From the center of the image (with position (sr, sc) = (1, 1)), all pixels connected
// by a path of the same color as the starting pixel are colored with the new color.
// Note the bottom corner is not colored 2, because it is not 4-directionally connected
// to the starting pixel.

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if len(image) < 1 || len(image[0]) < 1 {
		return image
	}

	m, n := len(image), len(image[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	oldColor := image[sr][sc]
	floodFillDFS(image, sr, sc, oldColor, newColor, visited)

	return image
}

func floodFillDFS(image [][]int, sr int, sc int, oldColor, newColor int, visited [][]bool) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) {
		return
	}

	if visited[sr][sc] {
		return
	}

	if image[sr][sc] != oldColor {
		return
	}

	// image[sr][sc] == oldColor
	image[sr][sc] = newColor
	visited[sr][sc] = true

	floodFillDFS(image, sr-1, sc, oldColor, newColor, visited)
	floodFillDFS(image, sr+1, sc, oldColor, newColor, visited)
	floodFillDFS(image, sr, sc-1, oldColor, newColor, visited)
	floodFillDFS(image, sr, sc+1, oldColor, newColor, visited)
}
