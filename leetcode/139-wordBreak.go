package leetcode

// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.
//
// Note:
// The same word in the dictionary may be reused multiple times in the segmentation.
// You may assume the dictionary does not contain duplicate words.
//
// Example 1:
// Input: s = "leetcode", wordDict = ["leet", "code"]
// Output: true
// Explanation: Return true because "leetcode" can be segmented as "leet code".
//
// Example 2:
// Input: s = "applepenapple", wordDict = ["apple", "pen"]
// Output: true
// Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
//              Note that you are allowed to reuse a dictionary word.
//
// Example 3:
// Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// Output: false

func wordBreak(s string, wordDict []string) bool {
	if s == "" || len(wordDict) < 1 {
		return false
	}

	n := len(s)
	states := make([]bool, n+1)
	states[0] = true

	maxLen := 0
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}
	for i := 0; i < n; i++ {
		// find start index
		for j := i + 1; j < n+1; j++ {
			// start from i, end at j-1
			seg := s[i:j]
			// use max len to skip reduce loop
			if len(seg) > maxLen {
				break
			}

			if states[i] && dict[seg] {
				states[j] = true
			}
		}
	}
	return states[n]
}
