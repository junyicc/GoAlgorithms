package leetcode

import (
	"strings"
	"unicode"
)

// Given an input string, reverse the string word by word.

// Example 1:
// Input: "the sky is blue"
// Output: "blue is sky the"

// Example 2:
// Input: "  hello world!  "
// Output: "world! hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.

// Example 3:
// Input: "a good   example"
// Output: "example good a"
// Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

func reverseWords(s string) string {
	if s == "" {
		return ""
	}
	s = strings.TrimSpace(s)

	if len(s) < 2 {
		return s
	}

	originRs := []rune(s)
	rs := make([]rune, 0, len(originRs))
	rs = append(rs, originRs[0])
	for i := 1; i < len(originRs); i++ {
		if unicode.IsSpace(originRs[i]) && unicode.IsSpace(originRs[i-1]) {
			continue
		}
		rs = append(rs, originRs[i])
	}

	lo, hi := 0, 0
	for lo < len(rs) {
		if unicode.IsSpace(rs[lo]) {
			lo++
			hi++
		} else if hi == len(rs) || unicode.IsSpace(rs[hi]) {
			reverseRuneSlice(rs, lo, hi-1)
			hi++
			lo = hi
		} else {
			hi++
		}
	}
	reverseRuneSlice(rs, 0, len(rs)-1)
	return string(rs)
}

func reverseRuneSlice(rs []rune, lo, hi int) {
	for lo < hi {
		rs[lo] = rs[lo] ^ rs[hi]
		rs[hi] = rs[lo] ^ rs[hi]
		rs[lo] = rs[lo] ^ rs[hi]
		lo++
		hi--
	}
}
