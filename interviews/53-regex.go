package interviews

// Match returns true if str matches pattern
func Match(str, pattern string) bool {
	if str == "" && pattern == "" {
		return true
	}
	if str == "" || pattern == "" {
		return false
	}

	ustr := []rune(str)
	upattern := []rune(pattern)

	return match(ustr, upattern, 0, 0)
}

func match(ustr, upattern []rune, i, j int) bool {
	if i == len(ustr) && j == len(upattern) {
		return true
	}
	if i == len(ustr) || j == len(upattern) {
		return false
	}

	if j < len(upattern)-2 && upattern[j+1] == '*' {
		if ustr[i] == upattern[j] || upattern[j] == '.' && i < len(ustr) {
			return match(ustr, upattern, i+1, j+2) || // match once
				match(ustr, upattern, i+1, j) || // continue matching
				match(ustr, upattern, i, j+2) // ignore
		}
		// ustr[i] != upattern[j]
		match(ustr, upattern, i, j+2) // ignore
	}
	if ustr[i] == upattern[j] || upattern[j] == '.' && i < len(ustr) {
		return match(ustr, upattern, i+1, j+1)
	}

	return false
}
