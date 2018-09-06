package leetcode

import (
	"strings"
)

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.
Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

Example 1:
Input: 3
Output: "III"

Example 2:
Input: 4
Output: "IV"

Example 3:
Input: 9
Output: "IX"

Example 4:
Input: 58
Output: "LVIII"
Explanation: C = 100, L = 50, XXX = 30 and III = 3.

Example 5:
Input: 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/

func intToRoman(num int) string {
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
}

func intToRoman2(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}
	romanMap := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
		4:    "IV",
		9:    "IX",
		40:   "XL",
		90:   "XC",
		400:  "CD",
		900:  "CM",
	}
	digit := 0
	stack := Stack{}
	for num > 0 {
		n := num % 10
		if digit > 0 {
			n = n * pow10(digit)
		}

		romanNum, ok := romanMap[n]
		if ok {
			stack.Push(romanNum)
		} else {
			if n > 1 && n < 5 {
				for i := 0; i < n; i++ {
					stack.Push(romanMap[1])
				}
			} else if n > 5 && n < 9 {
				for i := 0; i < n-5; i++ {
					stack.Push(romanMap[1])
				}
				stack.Push(romanMap[5])
			} else if n > 10 && n < 50 {
				for i := 0; i < n; i += 10 {
					stack.Push(romanMap[10])
				}
			} else if n > 50 && n < 90 {
				for i := 0; i < n-50; i += 10 {
					stack.Push(romanMap[10])
				}
				stack.Push(romanMap[50])
			} else if n > 100 && n < 500 {
				for i := 0; i < n; i += 100 {
					stack.Push(romanMap[100])
				}
			} else if n > 500 && n < 900 {
				for i := 0; i < n-500; i += 100 {
					stack.Push(romanMap[100])
				}
				stack.Push(romanMap[500])
			} else if n > 1000 && n <= 3000 {
				for i := 0; i < n; i += 1000 {
					stack.Push(romanMap[1000])
				}
			}
		}
		num /= 10
		digit++
	}
	var buf strings.Builder
	for !stack.IsEmpty() {
		c := (stack.Pop()).(string)
		buf.WriteString(c)
	}
	return buf.String()
}

func pow10(digit int) int {
	n := 1
	for i := 0; i < digit; i++ {
		n *= 10
	}
	return n
}
