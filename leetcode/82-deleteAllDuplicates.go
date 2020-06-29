package leetcode

// Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.
//
// Return the linked list sorted as well.

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	sentinel := &ListNode{}
	sentinel.Next = head
	node := sentinel
	for node.Next != nil && node.Next.Next != nil {
		if node.Next.Val == node.Next.Next.Val {
			rmVal := node.Next.Val
			for node.Next != nil && node.Next.Val == rmVal {
				node.Next = node.Next.Next
			}
		} else {
			node = node.Next
		}
	}

	return sentinel.Next
}
