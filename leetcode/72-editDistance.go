package leetcode

import "math"

// Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.
//
// You have the following 3 operations permitted on a word:
// - Insert i character
//     - insert character word2[j] before word1[i], compare word1[i] and word2[j+1]
//     - insert character word1[i] before word2[j], compare word1[i+1] and word2[j]
// - Delete i character
//     - delete word1[i], compare word1[i+1] and word2[j]
//     - delete word2[j], compare word1[i] and word2[j+1]
// - Replace i character
//     - replace word1[i] with word2[j], or replace word2[j] with word1[i], compare word1[i+1] and word2[j+1]

func minDistance(word1 string, word2 string) int {
	if len(word1) < 1 {
		return len(word2)
	}
	if len(word2) < 1 {
		return len(word1)
	}

	m, n := len(word1), len(word2)
	states := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		states[i] = make([]int, n+1)
	}
	// init states
	for j := 1; j < n+1; j++ {
		states[0][j] = j

	}
	for i := 1; i < m+1; i++ {
		states[i][0] = i
	}
	// update states
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				states[i][j] = states[i-1][j-1]
			} else {
				states[i][j] = minOfThree(states[i-1][j-1], states[i-1][j], states[i][j-1]) + 1
			}
		}
	}
	return states[m][n]
}

func minOfThree(x, y, z int) int {
	minNum := math.MaxInt64
	if x < minNum {
		minNum = x
	}
	if y < minNum {
		minNum = y
	}
	if z < minNum {
		minNum = z
	}
	return minNum
}
