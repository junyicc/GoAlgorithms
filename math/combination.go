package math

// Combination returns possibilities of selecting m elements from n
func Combination(n, m int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	if m == 0 || n <= m {
		return 1
	}
	return Factorial(n) / (Factorial(n-m) * Factorial(m))
}

// CombinationRecur returns possibilities of selecting m elements from n
func CombinationRecur(n, m int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	if m == 0 || n <= m {
		return 1
	}
	return CombinationRecur(n-1, m) + CombinationRecur(n-1, m-1)
}

// CombinationIter returns possibilities of selecting m elements from n
func CombinationIter(n, m int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
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
