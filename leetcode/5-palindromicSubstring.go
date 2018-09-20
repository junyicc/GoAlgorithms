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
	if s == "" || len(s) < 2 {
		return s
	}
	strLen := len(s)
	lo, maxLen := 0, 1
	for i := 0; i < strLen; i++ {
		extendPalindrome(s, i, i, &lo, &maxLen)   // odd
		extendPalindrome(s, i, i+1, &lo, &maxLen) // even
	}
	return s[lo : lo+maxLen]
}

func extendPalindrome(s string, l, r int, lo, maxLen *int) {
	strLen := len(s)
	for l >= 0 && r < strLen && s[l] == s[r] {
		l--
		r++
	}
	if *maxLen <= r-l-1 {
		*lo = l + 1
		*maxLen = r - l - 1
	}
}
