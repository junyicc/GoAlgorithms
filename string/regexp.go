package string

// regexp token: . *

func isMatch(s string, pattern string) bool {
	return matchCore([]byte(s), 0, []byte(pattern), 0)
}

func matchCore(s []byte, sIdx int, pattern []byte, pIdx int) bool {
	// both s and pattern end
	if sIdx >= len(s) && pIdx >= len(pattern) {
		return true
	}
	// s ends, but pattern does not end
	if sIdx >= len(s) {
		// find if there is pairs of *
		// the remaining pattern char number must be even
		if (len(pattern)-pIdx)&1 != 0 {
			return false
		}
		for i := pIdx + 1; i < len(pattern); i += 2 {
			if pattern[i] != '*' {
				return false
			}
		}
		return true
	}

	// s has not ended, but pattern ends
	if pIdx >= len(pattern) {
		return false
	}
	// if the next char of pattern is *
	if pIdx+1 < len(pattern) && pattern[pIdx+1] == '*' {
		if s[sIdx] == pattern[pIdx] || pattern[pIdx] == '.' {
			// current char equal
			return matchCore(s, sIdx+1, pattern, pIdx+2) ||
				matchCore(s, sIdx+1, pattern, pIdx) ||
				matchCore(s, sIdx, pattern, pIdx+2)
		} else {
			// current char not equal
			// ignore current pattern char
			return matchCore(s, sIdx, pattern, pIdx+2)
		}
	} else {
		// the next char of pattern is not *
		if s[sIdx] == pattern[pIdx] || pattern[pIdx] == '.' {
			return matchCore(s, sIdx+1, pattern, pIdx+1)
		}
	}
	return false
}
