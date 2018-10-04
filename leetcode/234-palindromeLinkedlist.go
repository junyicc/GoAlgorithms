package leetcode

func isPalindromeLinkedList(head *ListNode) bool {
	if head == nil || (head != nil && head.Next == nil) {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	slow = reverseList(slow)
	fast = head

	for slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true
}
