package leetcodecn

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	d := new(ListNode)
	n := d
	n1 := list1
	n2 := list2
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
	return d.Next
}

// @lc code=end
