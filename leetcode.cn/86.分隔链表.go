package leetcodecn

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	d1, d2 := new(ListNode), new(ListNode)
	n1, n2 := d1, d2
	n := head
	for n != nil {
		if n.Val < x {
			n1.Next = n
			n1 = n1.Next
		} else {
			n2.Next = n
			n2 = n2.Next
		}
		tmp := n.Next
		n.Next = nil
		n = tmp
	}
	n1.Next = d2.Next
	return d1.Next
}

// @lc code=end
