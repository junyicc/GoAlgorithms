package leetcodecn

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nA, nB := headA, headB
	for nA != nB {
		// next
		if nA == nil {
			nA = headB
		} else {
			nA = nA.Next
		}
		if nB == nil {
			nB = headA
		} else {
			nB = nB.Next
		}
	}
	return nA
}

// @lc code=end
