package leetcode

// Given a string, your task is to count how many palindromic substrings in this string.
//
// The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.
//
// Example 1:
// Input: "abc"
// Output: 3
// Explanation: Three palindromic strings: "a", "b", "c".
//
//
// Example 2:
// Input: "aaa"
// Output: 6
// Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

func countSubstrings(s string) int {
	if s == "" {
		return 0
	}

	n := len(s)
	var cnt int
	for i := 0; i < 2*n-1; i++ {
		lo := i >> 1
		hi := lo + i%2
		for lo >= 0 && hi < n && s[lo] == s[hi] {
			cnt++
			lo--
			hi++
		}
	}
	return cnt
}
