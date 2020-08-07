package interviews

import "github.com/CAIJUNYI/GoAlgorithms/datastructure"

// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
//
// 示例1：
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	sentinel := new(ListNode)
	node := sentinel

	n1, n2 := l1, l2
	for n1 != nil && n2 != nil {
		if n1.Val < n2.Val {
			node.Next = n1
			n1 = n1.Next
		} else {
			node.Next = n2
			n2 = n2.Next
		}
		node = node.Next
	}

	if n1 != nil {
		node.Next = n1
	}
	if n2 != nil {
		node.Next = n2
	}

	return sentinel.Next
}

// mergeLinkedlist merge sorted linkedlist recursively
func mergeLinkedlist(h1, h2 *datastructure.ListNode) *datastructure.ListNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	var h *datastructure.ListNode
	if h1.Less(h2) {
		h = h1
		h.Next = mergeLinkedlist(h1.Next, h2)
	} else {
		h = h2
		h.Next = mergeLinkedlist(h1, h2.Next)
	}
	return h
}
