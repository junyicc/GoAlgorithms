package leetcode

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	carry := 0
	var l, n *ListNode
	start := true
	n1, n2 := l1, l2
	for n1 != nil && n2 != nil {
		sum := n1.Val + n2.Val + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		node := ListNode{
			Val: sum,
		}
		if start {
			l = &node
			n = l
			start = false
		} else {
			n.Next = &node
			n = n.Next
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	for n1 != nil {
		sum := n1.Val + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		node := ListNode{
			Val: sum,
		}
		n.Next = &node
		n = n.Next
		n1 = n1.Next
	}
	for n2 != nil {
		sum := n2.Val + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}
		node := ListNode{
			Val: sum,
		}
		n.Next = &node
		n = n.Next
		n2 = n2.Next
	}
	if carry > 0 {
		node := ListNode{
			Val: carry,
		}
		n.Next = &node
		n = n.Next
	}
	return l
}
