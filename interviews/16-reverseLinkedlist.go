package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// ReverseLinkedlist reverse origin linkedlist
func ReverseLinkedlist(head *datastructure.Node) *datastructure.Node {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, next, reverseHead *datastructure.Node
	curNode := head
	for curNode != nil {
		next = curNode.Next
		if next == nil {
			reverseHead = curNode
		}
		curNode.Next = pre
		pre = curNode
		curNode = next
	}
	return reverseHead
}

// ReverseLinkedlistWithHead reverse origin linkedlist
func ReverseLinkedlistWithHead(head *datastructure.Node) *datastructure.Node {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, next *datastructure.Node
	curNode := head.Next
	for curNode != nil {
		next = curNode.Next
		if next == nil {
			head.Next = curNode
		}
		curNode.Next = pre
		pre = curNode
		curNode = next
	}
	return head
}
