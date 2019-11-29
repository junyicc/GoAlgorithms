package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// ReverseLinkedlist reverse origin linkedlist
func ReverseLinkedlist(head *datastructure.ListNode) *datastructure.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, next, reverseHead *datastructure.ListNode
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
func ReverseLinkedlistWithHead(head *datastructure.ListNode) *datastructure.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, next *datastructure.ListNode
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
