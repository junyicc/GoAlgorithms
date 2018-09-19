package leetcode

/*
Given a string, find the length of the longest substring without repeating characters.

Examples:
Given "abcabcbb", the answer is "abc", which the length is 3.
Given "bbbbb", the answer is "b", with the length of 1.
Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	strLen := len(s)
	strMap := make([]int, 256)
	for i := 0; i < 256; i++ {
		strMap[i] = -1
	}
	maxLen := 0
	noRepeatLen := 0
	for i := 0; i < strLen; i++ {
		c := s[i]
		if strMap[c] == -1 {
			strMap[c] = i
			noRepeatLen++
		} else {
			if noRepeatLen > maxLen {
				maxLen = noRepeatLen
			}
			prePos := strMap[c]
			for j := i - noRepeatLen; j < prePos; j++ {
				strMap[s[j]] = -1
			}
			strMap[c] = i
			noRepeatLen = i - prePos
		}
	}
	if noRepeatLen > maxLen {
		maxLen = noRepeatLen
	}
	return maxLen
}
