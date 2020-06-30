package leetcode

// A linked list is given such that each node contains an additional random pointer which could point to any node in the list or null.
//
// Return a deep copy of the list.
//
// The Linked List is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:
// val: an integer representing Node.val
// random_index: the index of the node (range from 0 to n-1) where random pointer points to, or null if it does not point to any node.

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// clone listnodes
	node := head
	for node != nil {
		cloneNode := new(Node)
		cloneNode.Val = node.Val

		cloneNode.Next = node.Next
		node.Next = cloneNode

		node = node.Next.Next
	}

	// construct random pointer
	node = head
	for node != nil {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
		node = node.Next.Next
	}

	// detach two lists
	node = head
	cloneHead := head.Next
	cloneNode := cloneHead
	for cloneNode.Next != nil {
		node.Next = cloneNode.Next
		cloneNode.Next = node.Next.Next

		node = node.Next
		cloneNode = cloneNode.Next
	}
	node.Next = nil
	return cloneHead
}
