package leetcode

import (
	"math"
	"strings"
)

/*
Implement atoi which converts a string to an integer.
The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.
The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.
If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.
If no valid conversion could be performed, a zero value is returned.

Note:
Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.

Example 1:
Input: "42"
Output: 42

Example 2:
Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
			 Then take as many numerical digits as possible, which gets 42.

Example 3:
Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.

Example 4:
Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical
             digit or a +/- sign. Therefore no valid conversion could be performed.

Example 5:
Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−231) is returned.
*/
func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}

	s := []byte(str)
	if str[0] == '+' || str[0] == '-' {
		s = s[1:]
	}
	start := 0
	for _, c := range s {
		if c == '0' {
			start++
		} else {
			break
		}
	}
	s = s[start:]

	num := 0
	numLen := 0
	for i, c := range s {
		if c < '0' || c > '9' {
			break
		}
		if i > 9 {
			if str[0] == '-' {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		n := int(c - '0')
		num = num*10 + n
		numLen++
	}

	if str[0] == '-' {
		num = 0 - num
		if num < math.MinInt32 {
			return math.MinInt32
		}
	}
	if num > math.MaxInt32 {
		return math.MaxInt32
	}

	return num
}
