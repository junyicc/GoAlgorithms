package interviews

// FindFirstNotRepeatingChar returns the first not repeating char in the string
func FindFirstNotRepeatingChar(s string) byte {
	if s == "" {
		return 0
	}
	cnt := [256]int{}
	strArr := []byte(s)
	for _, c := range strArr {
		cnt[c]++
	}

	for _, c := range strArr {
		if cnt[c] == 1 {
			return c
		}
	}
	return 0
}
