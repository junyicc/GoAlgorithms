package leetcode

import "math"

// Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).
//
// Example:
// Input: S = "ADOBECODEBANC", T = "ABC"
// Output: "BANC"

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}

	// record the count of each rune in t
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	win := make(map[byte]int)
	valid := 0
	left, right := 0, 0
	// record result
	start := 0
	length := math.MaxInt32

	for right < len(s) {
		// [left, right)
		// expand right-side win
		c := s[right]
		right++

		// update window state
		if need[c] > 0 {
			win[c]++
			if win[c] == need[c] {
				valid++
			}
		}

		// shrink left-side win
		for valid == len(need) {
			// record result
			if right-left < length {
				start = left
				length = right - left
			}

			// shrink
			d := s[left]
			left++

			// update window state
			if need[d] > 0 {
				if win[d] == need[d] {
					valid--
				}
				win[d]--
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
