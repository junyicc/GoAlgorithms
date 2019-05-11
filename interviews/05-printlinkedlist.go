package interviews

import (
	"fmt"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// PrintLinkedlistReversely prints linkedlist reversely
func PrintLinkedlistReversely(list *datastructure.LinkedList) {
	if list == nil || list.Head() == nil {
		return
	}

	stack := datastructure.NewDynamicStack()
	node := list.Head()
	for node != nil {
		stack.Push(node.Data)
		node = node.Next
	}

	for !stack.IsEmpty() {
		fmt.Println(*stack.Pop())
	}
}
