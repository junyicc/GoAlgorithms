package leetcode

func isValid(s string) bool {
	if s == "" {
		return true
	}
	stack := Stack{}
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			stack.Push(c)
		} else if c == '}' {
			if !stack.IsEmpty() && (stack.Top()).(rune) == '{' {
				stack.Pop()
			} else {
				return false
			}
		} else if c == ']' {
			if !stack.IsEmpty() && (stack.Top()).(rune) == '[' {
				stack.Pop()
			} else {
				return false
			}
		} else if c == ')' {
			if !stack.IsEmpty() && (stack.Top()).(rune) == '(' {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	if stack.IsEmpty() {
		return true
	}
	return false
}
