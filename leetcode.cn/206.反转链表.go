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

// @lc code=end
