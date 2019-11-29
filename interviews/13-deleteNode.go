package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// DeleteNode deletes the toBeDel node in the linkedlist in O(1)
func DeleteNode(head **datastructure.ListNode, toBeDel *datastructure.ListNode) *datastructure.Elem {
	if head == nil || toBeDel == nil {
		return nil
	}
	e := toBeDel.Data
	if toBeDel.Next != nil {
		// O(1)
		next := toBeDel.Next
		toBeDel.Data = next.Data
		toBeDel.Next = next.Next
		next.Next = nil
		next = nil
	} else if *head == toBeDel {
		// only one node in linkedlist, delete head(tail) node
		*head = nil
	} else {
		// delete tail node
		node := *head
		// O(n)
		for node.Next != toBeDel {
			node = node.Next
		}
		node.Next = nil
	}
	return &e
}
