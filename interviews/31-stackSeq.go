package interviews

import "github.com/CAIJUNYI/GoAlgorithms/datastructure"

// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。
// 例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。
//
// 示例 1：
// 输入：pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
// 输出：true
// 解释：我们可以按以下顺序执行：
// push(1), push(2), push(3), push(4), pop() -> 4,
// push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
//
// 示例 2：
// 输入：pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
// 输出：false
// 解释：1 不能在 2 之前弹出。

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}
	n := len(pushed)
	stack := make([]int, 0, n)

	var pushIdx, popIdx int
	for popIdx < n {
		curNum := popped[popIdx]
		if len(stack) == 0 || stack[len(stack)-1] != curNum {
			if pushIdx >= n {
				return false
			}
			stack = append(stack, pushed[pushIdx])
			pushIdx++
		} else {
			stack = stack[:len(stack)-1]
			popIdx++
		}
	}

	return popIdx == n && len(stack) == 0
}

// isPopOrder returns true if popOrder of stack is natural
func isPopOrder(pushOrder, popOrder []int) bool {
	if pushOrder == nil || popOrder == nil {
		return false
	}
	pushLen := len(pushOrder)
	popLen := len(popOrder)
	if pushLen < 1 || popLen < 1 || pushLen != popLen {
		return false
	}
	var pushIndex, popIndex int
	stack := datastructure.DynamicStack{}
	for popIndex < popLen {
		p := popOrder[popIndex]
		for stack.IsEmpty() || (!datastructure.Equal(stack.GetTop(), p) && pushIndex < popLen) {
			stack.Push(pushOrder[pushIndex])
			pushIndex++
		}
		if !datastructure.Equal(stack.GetTop(), p) {
			return false
		}
		stack.Pop()
		popIndex++
	}

	if stack.IsEmpty() && popIndex == popLen {
		return true
	}
	return false
}
