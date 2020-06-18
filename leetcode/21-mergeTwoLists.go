package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	l := new(ListNode)
	n, n1, n2 := l, l1, l2
	for n1 != nil && n2 != nil {
		if n1.Val < n2.Val {
			n.Next = n1
			n1 = n1.Next
		} else {
			n.Next = n2
			n2 = n2.Next
		}
		n = n.Next
	}

	if n1 != nil {
		n.Next = n1
	}
	if n2 != nil {
		n.Next = n2
	}
	return l.Next
}

func mergeTwoListsRecur(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var l ListNode
	if l1.Val < l2.Val {
		l.Val = l1.Val
		l.Next = mergeTwoListsRecur(l1.Next, l2)
	} else {
		l.Val = l2.Val
		l.Next = mergeTwoListsRecur(l1, l2.Next)
	}
	return &l
}
