package interviews

import (
	"strconv"
)

// NumberOf1 returns number of 1 from 1 to n
func NumberOf1(num int) (int, bool) {
	if num <= 0 {
		return 0, false
	}
	s := strconv.Itoa(num)
	return numberOf1([]byte(s))
}

func numberOf1(num []byte) (int, bool) {
	if num == nil || len(num) < 1 || num[0] < '0' || num[0] > '9' {
		return 0, false
	}
	first := int(num[0] - '0')
	length := len(num)
	if length == 1 && first == 0 {
		return 0, true
	}
	if length == 1 && first > 0 {
		return 1, true
	}
	// number at first digit
	numAtFirstDigit := 0
	if first > 1 {
		numAtFirstDigit = powerBase10(uint(length - 1))
	} else if first == 1 {
		numFirstDigit, err := strconv.Atoi(string(num[1:]))
		if err != nil {
			return 0, false
		}
		numAtFirstDigit = numFirstDigit + 1
	}
	// number at other digits: permutation
	numAtOtherDigits := first * (length - 1) * powerBase10(uint(length-2))
	// number of 1 at lower digits
	numOfRecur, ok := numberOf1(num[1:])
	if !ok {
		return 0, false
	}
	return numAtFirstDigit + numAtOtherDigits + numOfRecur, true
}

func powerBase10(n uint) int {
	result := 1
	for i := uint(0); i < n; i++ {
		result *= 10
	}
	return result
}
