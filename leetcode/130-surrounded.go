package leetcode

import "github.com/CAIJUNYI/GoAlgorithms/datastructure"

// Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
//
// A region is captured by flipping all 'O's into 'X's in that surrounded region.
//
// Example:
// X X X X
// X O O X
// X X O X
// X O X X
//
// After running your function, the board should be:
// X X X X
// X X X X
// X X X X
// X O X X
//
// Explanation:
// Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.

func solve(board [][]byte) {
	if len(board) < 1 || len(board[0]) < 1 {
		return
	}

	// use union-find to separate two parts
	// the one cannot be surrounded by X
	// the other is surrounded by X

	// convert 2D-matrix to 1D-array
	// [i][j] -> i*n + j

	m := len(board)
	n := len(board[0])

	// the m*n element is a virtual node that is connected to the nodes cannot be surrounded by X
	// the m*n node is the sentinel node
	uf := datastructure.NewUnionFindInt(m*n + 1)

	// process the 0 and n-1 column
	// connect to the sentinel node
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			uf.Union(i*n, m*n)
		}
		if board[i][n-1] == 'O' {
			uf.Union(i*n+(n-1), m*n)
		}
	}
	// process the 0 and m-1 row
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			uf.Union(j, m*n)
		}
		if board[m-1][j] == 'O' {
			uf.Union((m-1)*n+j, m*n)
		}
	}

	// connect each remaining element to its surround element if it is 'O'
	direction := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				for k := 0; k < 4; k++ {
					x := i + direction[k][0]
					y := j + direction[k][1]
					if board[x][y] == 'O' {
						uf.Union(i*n+j, x*n+y)
					}
				}
			}
		}
	}

	// fill the surrounded 'O'
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if !uf.Connected(i*n+j, m*n) {
				board[i][j] = 'X'
			}
		}
	}
}
