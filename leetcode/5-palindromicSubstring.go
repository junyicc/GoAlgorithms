package leetcode

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:
Input: "cbbd"
Output: "bb"
*/

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	n := len(s)
	var res string
	// why is 2*n-1:
	// 	- n: every i as a center
	// 	- n-1: every (i, i+1) as a center
	for i := 0; i < 2*n-1; i++ {
		lo := i >> 1
		hi := lo + i%2

		for lo >= 0 && hi < n && s[lo] == s[hi] {
			tmp := s[lo : hi+1]
			if len(tmp) > len(res) {
				res = tmp
			}
			lo--
			hi++
		}
	}
	return res
}

func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}

	n := len(s)
	var res string
	for i := 0; i < n; i++ {
		s1 := extendPalindromeString(s, i, i)
		s2 := extendPalindromeString(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func extendPalindromeString(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
