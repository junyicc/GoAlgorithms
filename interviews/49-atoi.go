package interviews

import (
	"fmt"
	"math"
	"strconv"
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

func atoi(s string) (int, error) {
	if s == "" || len(s) < 1 {
		return 0, fmt.Errorf("atoi: invalid input")
	}

	sLen := len(s)
	if 0 < sLen && sLen < 19 {
		s0 := s[0]
		if s0 == '-' || s0 == '+' {
			s = s[1:]
			if len(s) < 1 {
				// only '+' or '-'
				return 0, fmt.Errorf("atoi: invalid input")
			}
		}
		// convert
		n := 0
		for _, c := range s {
			c -= '0'
			// invalid character
			if c < 0 || c > 9 {
				return 0, fmt.Errorf("atoi: invalid character in string number")
			}
			n = n*10 + int(c)
		}
		if s0 == '-' {
			n = 0 - n
		}
		return n, nil
	}

	n, err := strconv.ParseInt(s, 10, 0)
	return int(n), err
}
