package string

import "unicode/utf8"

func BF(main, pattern string) int {
	if len(main) < 1 || len(pattern) < 1 || len(main) < len(pattern) {
		return -1
	}
	n := utf8.RuneCountInString(main)
	m := utf8.RuneCountInString(pattern)
	mainRunes := []rune(main)
	patternRunes := []rune(pattern)
	for i := 0; i <= n-m; i++ {
		sub := mainRunes[i : i+m]
		if string(sub) == string(patternRunes) {
			return i
		}
	}
	return -1
}
