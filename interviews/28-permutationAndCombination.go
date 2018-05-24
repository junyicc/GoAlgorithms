package interviews

import (
	"math"
)

// StringPermutation returns string permutation
func StringPermutation(str string, f func(string)) {
	if str == "" {
		return
	}
	stringPermutation(([]byte)(str), 0, f)
}

func stringPermutation(str []byte, begin int, f func(string)) {
	if begin == len(str) {
		f(string(str))
	} else {
		for i := begin; i < len(str); i++ {
			// swap
			str[begin], str[i] = str[i], str[begin]
			stringPermutation(str, begin+1, f)
			// reset
			str[begin], str[i] = str[i], str[begin]
		}
	}
}

// Permutation returns results of n permutation
func Permutation(n, m int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return Factorial(n) / Factorial(n-m)
}

// Factorial return factorial of n
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	return f
}

// Combination returns results of selecting m elements from n
func Combination(n, m int) int {
	if n == 0 {
		return 0
	}
	if m == 0 || n <= m {
		return 1
	}

	return Combination(n-1, m) + Combination(n-1, m-1)
}

// CombinationIter returns results of selecting m elements from n
func CombinationIter(n, m int) int {
	if n == 0 {
		return 0
	}
	if m == 0 || n <= m {
		return 1
	}
	// init a (n+1)*(m+1) matrix
	C := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		C[i] = make([]int, m+1)
	}
	C[0][0] = 1
	for row := 1; row < n+1; row++ {
		C[row][0] = 1
		for col := 1; col < m+1; col++ {
			C[row][col] = C[row-1][col] + C[row-1][col-1]
		}
	}
	return C[n][m]
}

// SolutionsOfNQueens returns solutions number of n queens problem
func SolutionsOfNQueens(n int) int {
	if n == 0 {
		return 0
	}
	queens := make([]int, n)
	cnt := 0
	solveNQueens(queens, 0, &cnt)
	return cnt
}

func solveNQueens(queens []int, n int, cnt *int) {
	if n == len(queens) {
		*cnt++
	} else {
		for i := 0; i < len(queens); i++ {
			queens[n] = i
			if isValid(queens, n) {
				solveNQueens(queens, n+1, cnt)
			}
		}
	}
}

func isValid(queens []int, n int) bool {
	for j := 0; j < n; j++ {
		if queens[j] == queens[n] || int(math.Abs(float64(queens[j]-queens[n]))) == n-j {
			return false
		}
	}
	return true
}
