package leetcode

// Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
//
// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
//
// Example:
//
// Given this linked list: 1->2->3->4->5
//
// For k = 2, you should return: 2->1->4->3->5
//
// For k = 3, you should return: 3->2->1->4->5

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k < 2 {
		return head
	}

	tail := head
	cnt := 0
	for tail != nil && cnt < k {
		tail = tail.Next
		cnt++
	}
	// not enough k
	if cnt != k {
		return head
	}

	newHead := reverseListBetween(head, tail)

	head.Next = reverseKGroup(tail, k)

	return newHead
}

func reverseListBetween(head *ListNode, tail *ListNode) *ListNode {
	var pre, next *ListNode
	for head != tail {
		next = head.Next
		head.Next = pre

		pre = head
		head = next
	}
	return pre
}
