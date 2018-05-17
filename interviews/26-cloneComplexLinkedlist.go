package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// CloneComplexLinkedlist returns copy of linkedlist with complex node
func CloneComplexLinkedlist(head *datastructure.ComplexListNode) *datastructure.ComplexListNode {
	if head == nil {
		return nil
	}

	clone(head)
	connectSibling(head)
	copyHead := divide(head)
	return copyHead
}

func clone(node *datastructure.ComplexListNode) {
	for node != nil {
		copyNode := datastructure.ComplexListNode{}
		copyNode.Data = node.Data
		// insertion
		copyNode.Next = node.Next
		node.Next = &copyNode

		node = copyNode.Next
	}
}

func connectSibling(node *datastructure.ComplexListNode) {
	for node != nil {
		copyNode := node.Next
		if node.Sibling != nil {
			sibling := node.Sibling
			copyNode.Sibling = sibling.Next
		}
		node = copyNode.Next
	}
}

func divide(head *datastructure.ComplexListNode) *datastructure.ComplexListNode {
	node := head
	copyNode := head.Next
	copyHead := copyNode

	for copyNode.Next != nil {
		node.Next = copyNode.Next
		copyNode.Next = node.Next.Next

		node = node.Next
		copyNode = copyNode.Next
	}

	node.Next = nil

	return copyHead
}
