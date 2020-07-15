package backtracking

import "fmt"

func CalcNQueen(n int, f func([]int)) {
	if n < 4 {
		return
	}
	// index: row, value: column
	res := make([]int, n, n)
	calcNQueen(res, 0, f)
}

func calcNQueen(res []int, row int, f func([]int)) {
	if row == len(res) {
		f(res)
		return
	}
	for col := 0; col < len(res); col++ {
		if ok(res, row, col) {
			// make choice
			res[row] = col
			// backtracking
			calcNQueen(res, row+1, f)
			// cancel choice
			res[row] = 0
		}
	}
}

func ok(res []int, row, col int) bool {
	leftUp, rightUp := col-1, col+1
	n := len(res)
	for i := row - 1; i >= 0; i-- {
		if res[i] == col {
			return false
		}
		if leftUp >= 0 && res[i] == leftUp {
			return false
		}
		if rightUp < n && res[i] == rightUp {
			return false
		}
		leftUp--
		rightUp++
	}
	return true
}

func printQueen(res []int) {
	n := len(res)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if res[i] == j {
				fmt.Print(" Q")
			} else {
				fmt.Print(" *")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
