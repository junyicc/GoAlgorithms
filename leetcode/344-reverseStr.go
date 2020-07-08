package leetcode

func reverseString(s []byte) {
	if len(s) < 2 {
		return
	}

	lo, hi := 0, len(s)-1
	for lo < hi {
		s[lo] = s[lo] ^ s[hi]
		s[hi] = s[lo] ^ s[hi]
		s[lo] = s[lo] ^ s[hi]
		lo++
		hi--
	}
}
