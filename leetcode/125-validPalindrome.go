package leetcode

import (
	"strings"
	"unicode"
)

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:
Input: "A man, a plan, a canal: Panama"
Output: true

Example 2:
Input: "race a car"
Output: false
*/

func validPalindrome(s string) bool {
	if s == "" {
		return true
	}
	lo, hi := 0, len(s)-1
	for lo <= hi {
		if !unicode.IsLetter(rune(s[lo])) && !unicode.IsNumber(rune(s[lo])) {
			lo++
			continue
		}
		if !unicode.IsLetter(rune(s[hi])) && !unicode.IsNumber(rune(s[hi])) {
			hi--
			continue
		}
		if strings.EqualFold(string(s[lo]), string(s[hi])) {
			lo++
			hi--
		} else {
			return false
		}
	}
	return true
}
