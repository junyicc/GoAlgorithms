package leetcode

// Given two lists of closed intervals, each list of intervals is pairwise disjoint and in sorted order.
//
// Return the intersection of these two interval lists.
//
// (Formally, a closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.  The intersection of two closed intervals is a set of real numbers that is either empty, or can be represented as a closed interval.  For example, the intersection of [1, 3] and [2, 4] is [2, 3].)
//
// Example 1:
// Input: A = [[0,2],[5,10],[13,23],[24,25]], B = [[1,5],[8,12],[15,24],[25,26]]
// Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]

func intervalIntersection(A [][]int, B [][]int) [][]int {
	if len(A) < 1 || len(B) < 1 {
		return nil
	}

	var res [][]int
	var i, j int
	for i < len(A) && j < len(B) {
		a1, a2 := A[i][0], A[i][1]
		b1, b2 := B[j][0], B[j][1]

		// intersects
		if a2 >= b1 && b2 >= a1 {
			lo := max(a1, b1)
			hi := min(a2, b2)
			res = append(res, []int{lo, hi})
		}

		if a2 > b2 {
			j++
		} else {
			i++
		}
	}

	return res
}
