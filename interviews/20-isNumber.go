package interviews

import "strings"

// 请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。
// 例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"0123"都表示数值，
// 但"12e"、"1a3.14"、"1.2.3"、"+-5"、"-1E-16"及"12e+5.4"都不是。

// 考虑：
// - 空格
// - 正负号
// - 小数点
// - e E
// - 数字

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}

	var numeric bool
	n := len(s)
	var idx int
	firstChar := s[0]
	// only one char
	if n == 1 {
		return isAsciiNumber(firstChar)
	}

	if firstChar == '+' || firstChar == '-' {
		idx++
	}

	// scan numeric part
	for idx < n && isAsciiNumber(s[idx]) {
		numeric = true
		idx++
	}

	// check float
	// .123
	// 123.
	// 12.3
	if idx < n && s[idx] == '.' {
		idx++
		for idx < n && isAsciiNumber(s[idx]) {
			numeric = true
			idx++
		}
	}

	// check science number
	if idx > 0 && idx < n-1 && numeric && (s[idx] == 'e' || s[idx] == 'E') {
		idx++
		if s[idx] == '+' || s[idx] == '-' {
			idx++
		}
		// 12e+ is not a number
		if idx == n {
			return false
		}

		for idx < n && isAsciiNumber(s[idx]) {
			numeric = true
			idx++
		}
	}

	return numeric && idx == n
}

func isAsciiNumber(b byte) bool {
	return b >= '0' && b <= '9'
}
