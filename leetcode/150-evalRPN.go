package leetcode

import (
	"container/list"
	"strconv"
)

// Evaluate the value of an arithmetic expression in Reverse Polish Notation.

// Valid operators are +, -, *, /. Each operand may be an integer or another expression.

// Note:
// Division between two integers should truncate toward zero.
// The given RPN expression is always valid. That means the expression would always evaluate to a result and there won't be any divide by zero operation.
//
// Example 1:
// Input: ["2", "1", "+", "3", "*"]
// Output: 9
// Explanation: ((2 + 1) * 3) = 9
//
// Example 2:
// Input: ["4", "13", "5", "/", "+"]
// Output: 6
// Explanation: (4 + (13 / 5)) = 6

func evalRPN(tokens []string) int {
	if len(tokens) < 1 {
		return 0
	}

	stack := list.New()
	for _, s := range tokens {
		switch s {
		case "+", "-", "*", "/":
			if stack.Len() < 2 {
				return 0
			}
			b, a := stack.Remove(stack.Back()).(int), stack.Remove(stack.Back()).(int)
			var res int
			switch s {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}
			stack.PushBack(res)
		default:
			// default number
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			stack.PushBack(n)
		}
	}
	return stack.Front().Value.(int)
}
