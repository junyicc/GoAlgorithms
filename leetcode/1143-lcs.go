package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	m, n := len(text1), len(text2)
	states := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		states[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1]^text2[j-1] == 0 {
				states[i][j] = states[i-1][j-1] + 1
			} else {
				states[i][j] = max(states[i-1][j], states[i][j-1])
			}
		}
	}
	return states[m][n]
}
