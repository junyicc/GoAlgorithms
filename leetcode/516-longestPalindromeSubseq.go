package leetcode

// Given a string s, find the longest palindromic subsequence's length in s. You may assume that the maximum length of s is 1000.
//
// Example 1:
// Input:
// "bbbab"
// Output:
// 4
// One possible longest palindromic subsequence is "bbbb".
//
//
// Example 2:
// Input:
// "cbbd"
// Output:
// 2
// One possible longest palindromic subsequence is "bb".

func longestPalindromeSubseq(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	n := len(s)

	// states[i][j]: longest palindrome length between s[i] and s[j]
	// if s[i] == s[j] {
	// 	states[i][j] = states[i+1][j-1] + 2
	// } else {
	// 	states[i][j] = max(states[i+1][j], states[i][j-1])
	// }

	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, n)
	}

	// init states
	for i := 0; i < n; i++ {
		states[i][i] = 1
	}

	// update states in the half top-right triangle
	// from bottom-left to top-right
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				states[i][j] = states[i+1][j-1] + 2
			} else {
				states[i][j] = max(states[i+1][j], states[i][j-1])
			}
		}
	}

	return states[0][n-1]
}
