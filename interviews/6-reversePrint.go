package interviews

import (
	"fmt"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// PrintLinkedlistReversely prints linkedlist reversely
func PrintLinkedlistReversely(list *datastructure.LinkedList) {
	if list == nil || list.FirstNode() == nil {
		return
	}

	stack := datastructure.NewDynamicStack()
	node := list.FirstNode()
	for node != nil {
		stack.Push(node.Data)
		node = node.Next
	}

	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := make([]int, 0)
	reverseListNode(head, &res)
	return res
}

func reverseListNode(node *ListNode, res *[]int) {
	if node == nil {
		return
	}

	reverseListNode(node.Next, res)

	*res = append(*res, node.Val)
}
