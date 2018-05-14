package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// MergeLinkedlist merges h1 and h2 that are sorted linkedlist
func MergeLinkedlist(h1, h2 *datastructure.Node) *datastructure.Node {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	var n, n1, n2 *datastructure.Node
	// head of merge linkedlist
	if h1.Less(h2) {
		n = h1
		n1 = h1.Next
		n2 = h2
	} else {
		n = h2
		n2 = h2.Next
		n1 = h1
	}
	h := n

	for ; n1 != nil && n2 != nil; n = n.Next {
		if n1.Less(n2) {
			n.Next = n1
			n1 = n1.Next
		} else {
			n.Next = n2
			n2 = n2.Next
		}
	}

	if n1 != nil {
		n.Next = n1
	}

	if n2 != nil {
		n.Next = n2
	}

	return h
}

// MergeLinkedlist2 merge sorted linkedlist recursively
func MergeLinkedlist2(h1, h2 *datastructure.Node) *datastructure.Node {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	var h *datastructure.Node
	if h1.Less(h2) {
		h = h1
		h.Next = MergeLinkedlist2(h1.Next, h2)
	} else {
		h = h2
		h.Next = MergeLinkedlist2(h1, h2.Next)
	}
	return h
}
