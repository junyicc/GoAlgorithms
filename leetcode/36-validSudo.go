package leetcode

// Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
// - Each row must contain the digits 1-9 without repetition.
// - Each column must contain the digits 1-9 without repetition.
// - Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 || len(board[0]) != 9 {
		return false
	}
	m, n := len(board), len(board[0])
	// validate each 3*3 sub-box
	validSubBox := true
	for i := 1; i < m; i += 3 {
		for j := 1; j < n; j += 3 {
			numStates := make([]byte, 10)
			for y := i - 1; y <= i+1; y++ {
				for x := j - 1; x <= j+1; x++ {
					if y >= 0 && y < m && x >= 0 && x < n {
						if board[y][x] == '.' {
							continue
						}
						// counting
						numStates[board[y][x]-'0']++
					}
				}
			}
			for c := 1; c < len(numStates); c++ {
				if numStates[c] > 1 {
					validSubBox = false
					break
				}
			}
			if !validSubBox {
				break
			}
		}
	}

	// validate rows
	validRow := true
	for i := 0; i < m; i++ {
		numStates := make([]byte, n+1)
		for j := 0; j < n; j++ {
			if board[i][j] == '.' {
				continue
			}
			numStates[board[i][j]-'0']++
		}
		for c := 1; c < len(numStates); c++ {
			if numStates[c] > 1 {
				validRow = false
				break
			}
		}
		if !validRow {
			break
		}
	}
	// validate columns
	validCol := true
	for j := 0; j < n; j++ {
		numStates := make([]byte, m+1)
		for i := 0; i < m; i++ {
			if board[i][j] == '.' {
				continue
			}
			numStates[board[i][j]-'0']++
		}
		for c := 1; c < len(numStates); c++ {
			if numStates[c] > 1 {
				validCol = false
				break
			}
		}
		if !validCol {
			break
		}
	}

	return validSubBox && validRow && validCol
}
