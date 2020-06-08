package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// IsPopOrder returns true if popOrder of stack is natural
func IsPopOrder(pushOrder, popOrder []int) bool {
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

// IsPopOrder2 returns true if popOrder of stack is natural
func IsPopOrder2(pushOrder, popOrder []int) bool {
	if pushOrder == nil || popOrder == nil || len(pushOrder) < 1 || len(popOrder) < 1 {
		return false
	}
	pushQ := datastructure.DynamicQueue{}
	for _, e := range pushOrder {
		pushQ.Enqueue(e)
	}
	popQ := datastructure.DynamicQueue{}
	for _, e := range popOrder {
		popQ.Enqueue(e)
	}
	stack := datastructure.DynamicStack{}
	for !popQ.IsEmpty() {
		p := popQ.Dequeue()
		for stack.IsEmpty() || (!datastructure.Equal(stack.GetTop(), p) && !pushQ.IsEmpty()) {
			stack.Push(pushQ.Dequeue())
		}
		if !datastructure.Equal(stack.GetTop(), p) {
			return false
		}
		stack.Pop()
	}
	if stack.IsEmpty() && popQ.IsEmpty() {
		return true
	}
	return false
}
