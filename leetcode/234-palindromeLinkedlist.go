package leetcode

func isPalindromeLinkedList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	secondHead := slow.Next
	slow.Next = nil

	fromTail := reverseList(secondHead)
	fromHead := head
	for fromHead != nil && fromTail != nil {
		if fromHead.Val != fromTail.Val {
			return false
		}
		fromHead = fromHead.Next
		fromTail = fromTail.Next
	}
	return true
}
