package leetcode

// Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
//
// You should preserve the original relative order of the nodes in each of the two partitions.
//
// Example:
// Input: head = 1->4->3->2->5->2, x = 3
// Output: 1->2->2->4->3->5

func partitionLinkedList(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	// init
	smallHead := &ListNode{}
	smallNode := smallHead

	bigHead := &ListNode{}
	bigNode := bigHead

	node := head
	for node != nil {
		if node.Val >= x {
			bigNode.Next = node
			bigNode = bigNode.Next
		} else {
			smallNode.Next = node
			smallNode = smallNode.Next
		}

		node = node.Next
	}
	if smallNode == nil {
		return bigHead
	}
	if bigNode == nil {
		return smallHead
	}

	smallNode.Next = bigHead.Next
	bigNode.Next = nil

	return smallHead.Next
}
