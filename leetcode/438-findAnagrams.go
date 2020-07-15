package leetcode

// Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.
//
// Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.
//
// The order of output does not matter.
//
// Example 1:
// Input:
// s: "cbaebabacd" p: "abc"
//
// Output:
// [0, 6]
//
// Explanation:
// The substring with start index = 0 is "cba", which is an anagram of "abc".
// The substring with start index = 6 is "bac", which is an anagram of "abc".
//
// Example 2:
// Input:
// s: "abab" p: "ab"
//
// Output:
// [0, 1, 2]
//
// Explanation:
// The substring with start index = 0 is "ab", which is an anagram of "ab".
// The substring with start index = 1 is "ba", which is an anagram of "ab".
// The substring with start index = 2 is "ab", which is an anagram of "ab".

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)

	need := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	left, right := 0, 0
	win := make(map[byte]int)
	var valid int
	for right < len(s) {
		// expand win on right-side
		c := s[right]
		right++
		// update win state
		if need[c] > 0 {
			win[c]++
			if win[c] == need[c] {
				valid++
			}
		}

		// shrink win
		if right-left >= len(p) {
			// record result
			if valid == len(need) {
				res = append(res, left)
			}

			d := s[left]
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
	return res
}
