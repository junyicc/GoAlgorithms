package leetcodecn

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	d := &ListNode{Next: head}
	preNode := findNthFromEnd(d, n+1)
	if preNode == nil {
		return head
	}
	theNode := preNode.Next
	preNode.Next = theNode.Next
	theNode = nil

	return d.Next
}

func findNthFromEnd(head *ListNode, n int) *ListNode {
	fast := &ListNode{Next: head}
	for i := 0; i < n; i++ {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next
	}

	slow := &ListNode{Next: head}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// @lc code=end
