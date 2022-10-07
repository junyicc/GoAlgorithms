package leetcodecn

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	d := &ListNode{Next: head}
	beforeLeftNode := indexList(d, left-1)
	leftNode := beforeLeftNode.Next

	rightNode := indexList(d, right)
	afterRightNode := rightNode.Next
	rightNode.Next = nil

	reverseList(leftNode)
	beforeLeftNode.Next = rightNode
	leftNode.Next = afterRightNode

	return d.Next
}

func indexList(head *ListNode, n int) *ListNode {
	for i := 0; i < n; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

// @lc code=end
