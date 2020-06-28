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

	n, m := len(word1), len(word2)
	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, m)
	}
	// init states
	if word1[0] != word2[0] {
		states[0][0] = 1
	}
	for j := 1; j < m; j++ {
		if word1[0] == word2[j] {
			states[0][j] = j
		} else {
			states[0][j] = states[0][j-1] + 1
		}
	}
	for i := 1; i < n; i++ {
		if word2[0] == word1[i] {
			states[i][0] = i
		} else {
			states[i][0] = states[i-1][0] + 1
		}
	}
	// update states
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if word1[i] == word2[j] {
				states[i][j] = minOfThree(states[i-1][j-1], states[i-1][j]+1, states[i][j-1]+1)
			} else {
				states[i][j] = minOfThree(states[i-1][j-1]+1, states[i-1][j]+1, states[i][j-1]+1)
			}
		}
	}
	return states[n-1][m-1]
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
