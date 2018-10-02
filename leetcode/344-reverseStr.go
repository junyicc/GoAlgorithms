package leetcode

func reverseString(s string) string {
	if s == "" || len(s) < 2 {
		return s
	}

	sb := []byte(s)
	lo, hi := 0, len(sb)-1
	for lo < hi {
		sb[lo] = sb[lo] ^ sb[hi]
		sb[hi] = sb[lo] ^ sb[hi]
		sb[lo] = sb[lo] ^ sb[hi]
		lo++
		hi--
	}
	return string(sb)
}
