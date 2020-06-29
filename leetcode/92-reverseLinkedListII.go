package leetcode

// Reverse a linked list from position m to n. Do it in one-pass.
//
// Note: 1 ≤ m ≤ n ≤ length of list.
//
// Example:
// Input: 1->2->3->4->5->NULL, m = 2, n = 4
// Output: 1->4->3->2->5->NULL

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}

	sentinel := &ListNode{}
	sentinel.Next = head

	// move to m
	var firstTail *ListNode
	node := head
	for i := 1; i < m && node != nil; i++ {
		firstTail = node
		node = node.Next
	}
	if node == nil {
		// m > list length
		return sentinel.Next
	}
	secondTail := node

	// reverse from m to n
	var pre *ListNode
	for i := m; i <= n && node != nil; i++ {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}
	secondHead := pre
	thirdHead := node

	// link three parts
	if firstTail == nil {
		secondTail.Next = thirdHead
		sentinel.Next = secondHead
	} else {
		firstTail.Next = secondHead
		secondTail.Next = thirdHead
	}
	return sentinel.Next
}
