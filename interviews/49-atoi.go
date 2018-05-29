package interviews

import (
	"math"
)

// Atoi converts string s to int
func Atoi(s string) (int, bool) {
	if s == "" || len(s) < 1 {
		return 0, false
	}

	num := 0
	negative := false
	index := 0
	// symbol
	if s[0] == '-' {
		negative = true
		index++
	} else if s[0] == '+' {
		index++
	}
	// only one symbol without number
	if index >= len(s) {
		return 0, false
	}
	// atoi core
	for ; index < len(s); index++ {
		// not number
		if s[index] < '0' || s[index] > '9' {
			return 0, false
		}
		flag := 1
		if negative {
			flag = -1
		}
		curNum := int(s[index] - '0')
		num = num*10 + flag*curNum
		// check boundary
		if !negative && (num > math.MaxInt64 || num < 0) {
			return 0, false
		}
		if negative && (num < math.MinInt64 || num > 0) {
			return 0, false
		}
	}
	return num, true
}
