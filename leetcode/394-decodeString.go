package leetcode

import (
	"strconv"
)

// Given an encoded string, return its decoded string.
//
// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.
//
// You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.
//
// Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].
//
//
// Example 1:
// Input: s = "3[a]2[bc]"
// Output: "aaabcbc"
//
// Example 2:
// Input: s = "3[a2[c]]"
// Output: "accaccacc"
//
// Example 3:
// Input: s = "2[abc]3[cd]ef"
// Output: "abcabccdcdcdef"
//
// Example 4:
// Input: s = "abc3[cd]xyz"
// Output: "abccdcdcdxyz"

func decodeString(s string) string {
	if s == "" {
		return ""
	}

	var stack []rune
	rs := []rune(s)

	for i := 0; i < len(rs); i++ {
		if rs[i] == ']' {
			// process result
			var tmp []rune
			for stack[len(stack)-1] != '[' {
				char := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				tmp = append(tmp, char)
			}
			// pop '['
			stack = stack[:len(stack)-1]

			// get num
			index := 1
			for index <= len(stack) && stack[len(stack)-index] >= '0' && stack[len(stack)-index] <= '9' {
				index++
			}
			num, _ := strconv.Atoi(string(stack[len(stack)-index+1:]))
			stack = stack[:len(stack)-index+1]

			for i := 0; i < num; i++ {
				for j := len(tmp) - 1; j >= 0; j-- {
					stack = append(stack, tmp[j])
				}
			}
		} else {
			// push back string
			stack = append(stack, rs[i])
		}
	}

	return string(stack)
}
