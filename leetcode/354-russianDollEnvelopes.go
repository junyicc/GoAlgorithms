package leetcode

import "sort"

// You have a number of envelopes with widths and heights given as a pair of integers (w, h). One envelope can fit into another if and only if both the width and height of one envelope is greater than the width and height of the other envelope.
//
// What is the maximum number of envelopes can you Russian doll? (put one inside other)
//
// Note:
// Rotation is not allowed.
//
// Example:
// Input: [[5,4],[6,4],[6,7],[2,3]]
// Output: 3
// Explanation: The maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) < 1 || len(envelopes[0]) < 2 {
		return 0
	}
	// 对于宽度w相同的数对，要对其高度h进行降序排序
	// 因为两个宽度相同的信封不能相互包含的，而逆序排序保证在w相同的数对中最多只选取一个计入 LIS
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return false
	})

	// calculate the length of longest increasing sequence with the height of envelopes
	n := len(envelopes)
	states := make([]int, n)
	for i := 0; i < n; i++ {
		states[i] = 1
	}
	maxLen := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] {
				if states[j]+1 > states[i] {
					states[i] = states[j] + 1
				}
			}
		}
		if states[i] > maxLen {
			maxLen = states[i]
		}
	}
	return maxLen
}
