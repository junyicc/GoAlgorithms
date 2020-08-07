package interviews

import "github.com/CAIJUNYI/GoAlgorithms/datastructure"

// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
//
// 示例:
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := head
	var pre, next *ListNode
	for node != nil {
		next = node.Next
		node.Next = pre

		pre = node
		node = next
	}
	return pre
}

// reverseLinkedlistWithHead reverse origin linkedlist
func reverseLinkedlistWithHead(head *datastructure.ListNode) *datastructure.ListNode {
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
