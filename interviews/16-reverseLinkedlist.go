package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// ReverseLinkedlist reverse origin linkedlist
func ReverseLinkedlist(head *datastructure.Node) *datastructure.Node {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, next *datastructure.Node
	node := head
	for node.Next != nil {
		next = node.Next
		node.Next = pre
		pre = node
		node = next
	}
	// tail
	node.Next = pre
	return node
}
