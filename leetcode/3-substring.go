package leetcode

/*
Given a string, find the length of the longest substring without repeating characters.

Examples:
Given "abcabcbb", the answer is "abc", which the length is 3.
Given "bbbbb", the answer is "b", with the length of 1.
Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

// d = the distance between position of i char and the last position of i char
// f(i) = f(i-1) + 1  // the first i or d > f(i-1)
// f(i) = d  // d <= f(i-1), record the temporate max len, and update non-repeating length

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
		prePos := strMap[c]
		if strMap[c] == -1 || i-prePos > noRepeatLen {
			noRepeatLen++
		} else {
			if noRepeatLen > maxLen {
				maxLen = noRepeatLen
			}
			noRepeatLen = i - prePos
		}
		strMap[c] = i
	}
	if noRepeatLen > maxLen {
		maxLen = noRepeatLen
	}
	return maxLen
}

func lengthOfLongestSubstringWithSlidingWin(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var res int
	left, right := 0, 0
	win := make(map[byte]int)
	for right < len(s) {
		// expand win on right-side
		c := s[right]
		right++
		win[c]++

		// shrink win on left-side when there is repeating byte
		for win[c] > 1 {
			d := s[left]
			left++
			win[d]--
		}

		// record result when there is no repeating byte
		if right-left > res {
			res = right - left
		}
	}
	return res
}
