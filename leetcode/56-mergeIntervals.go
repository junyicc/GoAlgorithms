package leetcode

import "sort"

// Given a collection of intervals, merge all overlapping intervals.
//
// Example 1:
// Input: [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
//
// Example 2:
// Input: [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.

func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return nil
	}

	// sort by start
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval[0] <= cur[1] {
			// intersect
			cur[1] = max(cur[1], interval[1])
		} else {
			res = append(res, cur)
			cur = interval
		}
	}
	res = append(res, cur)
	return res
}
