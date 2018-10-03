package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var l ListNode
	if l1.Val < l2.Val {
		l.Val = l1.Val
		l.Next = mergeTwoLists(l1.Next, l2)
	} else {
		l.Val = l2.Val
		l.Next = mergeTwoLists(l1, l2.Next)
	}
	return &l
}
