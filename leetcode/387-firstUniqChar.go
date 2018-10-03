package leetcode

func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}
	var charMap [256]int
	for _, c := range s {
		charMap[c]++
	}
	for i, c := range s {
		if charMap[c] == 1 {
			return i
		}
	}
	return -1
}
