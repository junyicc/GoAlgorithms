package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// FindEntryOfLoop returns the entry node of loop linkedlist
func FindEntryOfLoop(head *datastructure.ListNode) *datastructure.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			break
		}
	}

	slow = head
	for fast != nil && fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	if fast == nil {
		return nil
	}
	return slow
}
