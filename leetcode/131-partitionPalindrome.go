package leetcode

// Given a string s, partition s such that every substring of the partition is a palindrome.
//
// Return all possible palindrome partitioning of s.
//
// Example:
//
// Input: "aab"
// Output:
// [
// ["aa","b"],
// ["a","a","b"]
// ]

// 画递归树
// 解题思路
// 1、每一个结点表示剩余没有扫描到的字符串，产生分支是截取了剩余字符串的前缀；
//
// 2、产生前缀字符串的时候，判断前缀字符串是否是回文。
//
// 如果前缀字符串是回文，则可以产生分支和结点；
// 如果前缀字符串不是回文，则不产生分支和结点，这一步是剪枝操作。
// 3、在叶子结点是空字符串的时候结算，此时从根结点到叶子结点的路径，就是结果集里的一个结果，使用深度优先遍历，记录下所有可能的结果。
//
// 采用一个路径变量 path 搜索，path 全局使用一个（注意结算的时候，需要生成一个拷贝），因此在递归执行方法结束以后需要回溯，即将递归之前添加进来的元素拿出去；
// path 的操作只在列表的末端，因此合适的数据结构是栈。

func partitionPalindrome(s string) [][]string {
	if s == "" {
		return nil
	}

	n := len(s)
	var res [][]string
	var stack []string
	backtracking(s, 0, n, stack, &res)
	return res
}

func backtracking(s string, lo, hi int, stack []string, res *[][]string) {
	if lo == hi {
		g := make([]string, len(stack))
		g = append(stack[:0:0], stack...)
		*res = append(*res, g)
		return
	}

	for i := lo; i < hi; i++ {
		// current s[lo:i+1] is not a palindrome
		if !checkPalindrome(s, lo, i) {
			continue
		}

		stack = append(stack, s[lo:i+1])

		backtracking(s, i+1, hi, stack, res)
		// pop
		stack = stack[:len(stack)-1]
	}
}

func checkPalindrome(s string, lo, hi int) bool {
	for lo < hi {
		if s[lo] != s[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}
