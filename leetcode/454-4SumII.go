package leetcode

/*
Given four lists A, B, C, D of integer values, compute how many tuples (i, j, k, l) there are such that A[i] + B[j] + C[k] + D[l] is zero.
*/

func fourSumCount(A []int, B []int, C []int, D []int) int {
	if A == nil || B == nil || C == nil || D == nil || len(A) < 1 || len(B) < 1 || len(C) < 1 || len(D) < 1 {
		return 0
	}
	aN, bN, cN, dN := len(A), len(B), len(C), len(D)
	numMap := make(map[int]int, aN*aN)
	for i := 0; i < aN; i++ {
		for j := 0; j < bN; j++ {
			num := A[i] + B[j]
			numMap[num]++
		}
	}
	cnt := 0
	for i := 0; i < cN; i++ {
		for j := 0; j < dN; j++ {
			num := -(C[i] + D[j])
			if c, ok := numMap[num]; ok {
				cnt += c
			}
		}
	}
	return cnt
}
