package leetcode

// Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.
//
// Example 1:
// Input: s1 = "ab" s2 = "eidbaooo"
// Output: True
// Explanation: s2 contains one permutation of s1 ("ba").
//
// Example 2:
// Input:s1= "ab" s2 = "eidboaoo"
// Output: False

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	need := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	left, right := 0, 0
	win := make(map[byte]int)
	var valid int
	for right < len(s2) {
		// expand win on right-side
		c := s2[right]
		right++
		// update win state
		if need[c] > 0 {
			win[c]++
			if win[c] == need[c] {
				valid++
			}
		}
		// shrink win on left-side
		for right-left >= len(s1) {
			// record result
			if valid == len(need) {
				return true
			}

			d := s2[left]
			left++
			// update win state
			if need[d] > 0 {
				if win[d] == need[d] {
					valid--
				}
				win[d]--
			}
		}
	}
	return false
}
