package string

func equal(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	n := len(s1)
	b1, b2 := []byte(s1), []byte(s2)
	res := 0
	for i := 0; i < n; i++ {
		res |= int(b1[i] ^ b2[i])
	}
	return res == 0
}
