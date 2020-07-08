package leetcode

// Given a linked list, swap every two adjacent nodes and return its head.
//
// You may not modify the values in the list's nodes, only nodes itself may be changed.
//
// Example:
// Given 1->2->3->4, you should return the list as 2->1->4->3.

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	sentinel := &ListNode{}
	sentinel.Next = head
	pre := sentinel
	node := head
	for node != nil && node.Next != nil {
		next := node.Next.Next
		// change
		pre.Next = node.Next
		node.Next.Next = node
		node.Next = next

		// shift
		pre = node
		node = node.Next
	}
	return sentinel.Next
}
