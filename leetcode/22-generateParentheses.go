package leetcode

func generateParenthesis(n int) []string {
	if n < 1 {
		return []string{""}
	}
	var res []string
	addParentheses(&res, "", 0, 0, n)
	return res
}

func addParentheses(res *[]string, s string, left, right int, n int) {
	if len(s) == n*2 {
		*res = append(*res, s)
		return
	}

	if left < n {
		addParentheses(res, s+"(", left+1, right, n)
	}
	if right < left {
		addParentheses(res, s+")", left, right+1, n)
	}
}
