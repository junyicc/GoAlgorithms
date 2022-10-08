package leetcodecn

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// base case
	slow, fast := head, head
	for i := 0; i < k; i++ {
		if fast == nil {
			return slow
		}
		fast = fast.Next
	}

	// do reversion between [slow, fast)
	newHead := reverseNodesBetween(slow, fast)

	// sub problem
	slow.Next = reverseKGroup(fast, k)

	return newHead
}

// reverse nodes between [head, tail)
func reverseNodesBetween(head, tail *ListNode) *ListNode {
	pre := tail
	node := head
	for node != tail {
		next := node.Next
		node.Next = pre

		pre = node
		node = next
	}
	return pre
}

// @lc code=end
