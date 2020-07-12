package leetcode

// Given a string s, partition s such that every substring of the partition is a palindrome.
//
// Return the minimum cuts needed for a palindrome partitioning of s.
//
// Example:
//
// Input: "aab"
// Output: 1
// Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.

func minCut(s string) int {
	if len(s) < 2 {
		return 0
	}

	n := len(s)
	cnt := n - 1
	var stack []string
	backtrackingMinCut(s, 0, n, stack, &cnt)
	return cnt
}

func backtrackingMinCut(s string, lo, hi int, stack []string, cnt *int) {
	if lo == hi {
		if len(stack)-1 < *cnt {
			*cnt = len(stack) - 1
		}
		return
	}

	for i := lo; i < hi; i++ {

		if !checkPalindrome(s, lo, i) {
			continue
		}

		stack = append(stack, s[lo:i+1])

		backtrackingMinCut(s, i+1, hi, stack, cnt)

		stack = stack[:len(stack)-1]
	}
}
