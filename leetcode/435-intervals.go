package leetcode

import "sort"

// Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.
//
// Example 1:
// Input: [[1,2],[2,3],[3,4],[1,3]]
// Output: 1
// Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.
//
// Example 2:
// Input: [[1,2],[1,2],[1,2]]
// Output: 2
// Explanation: You need to remove two [1,2] to make the rest of intervals non-overlapping.
//
// Example 3:
// Input: [[1,2],[2,3]]
// Output: 0
// Explanation: You don't need to remove any of the intervals since they're already non-overlapping.

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	return n - nonOverlapIntervals(intervals)
}

func nonOverlapIntervals(intervals [][]int) int {
	if len(intervals) < 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// one at least
	cnt := 1
	hi := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval[0] >= hi {
			cnt++
			hi = interval[1]
		}
	}
	return cnt
}
