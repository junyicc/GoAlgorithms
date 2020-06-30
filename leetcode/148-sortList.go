package leetcode

func sortList(head *ListNode) *ListNode {
	return mergesortList(head)
}

func mergesortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mi := findMidNode(head)
	secondHead := mi.Next
	mi.Next = nil

	left := mergesortList(head)
	right := mergesortList(secondHead)

	return mergeTwoLists(left, right)
}

// findMidNode returns the middle node of linkedlist
func findMidNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lo, hi := head, head.Next
	for hi != nil && hi.Next != nil {
		hi = hi.Next.Next
		lo = lo.Next
	}
	return lo
}
