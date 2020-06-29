package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	node := head
	for node.Next != nil {
		if node.Next.Val == node.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return head
}
