package dynamic_programming

func LongestCommonSubsequence(s1, s2 string) int {
	if s1 == "" || s2 == "" {
		return 0
	}
	return lcs([]rune(s1), []rune(s2))
}

func lcs(r1, r2 []rune) int {
	n := len(r1)
	m := len(r2)
	states := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		states[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if r1[i-1] == r2[j-1] {
				states[i][j] = states[i-1][j-1] + 1
			} else {
				states[i][j] = max(states[i-1][j], states[i][j-1])
			}
		}
	}
	return states[n][m]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
