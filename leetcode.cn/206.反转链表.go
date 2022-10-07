package leetcodecn

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	n := head
	var pre *ListNode
	for n != nil {
		next := n.Next
		n.Next = pre

		pre = n
		n = next
	}
	return pre
}

func reverseListRecur(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseListRecur(head.Next)
	head.Next.Next = head
	head.Next = nil

	return last
}

// @lc code=end
