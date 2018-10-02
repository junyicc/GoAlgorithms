package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, next *ListNode
	node := head
	for node != nil {
		next = node.Next
		node.Next = pre

		pre = node
		node = next
	}
	return pre
}
