package leetcodecn

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	return isPalindromeByReverse(head)
}

// isPalindromeByReverse checks list isPalindrome by reversing half of the list
func isPalindromeByReverse(head *ListNode) bool {
	if head == nil {
		return false
	}

	// find mid
	mid := _findMid(head)

	// reverse half
	left := head
	right := _reverse(mid)
	last := right

	// check palindrome
	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	// restore list
	_reverse(last)

	return true
}

func _findMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func _reverse(head *ListNode) *ListNode {
	var pre, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

// isPalindromeByTraverse checks list isPalindrome recursively (O(n), O(n))
func isPalindromeByTraverse(head *ListNode) bool {
	if head == nil {
		return false
	}

	left := head
	right := head

	// traverse func
	var traverseList func(right *ListNode) bool
	traverseList = func(right *ListNode) bool {
		// base
		if right == nil {
			return true
		}
		// scale
		res := traverseList(right.Next)
		// change state: check palindrome
		res = res && (left.Val == right.Val)
		left = left.Next
		return res
	}
	return traverseList(right)
}

// @lc code=end
