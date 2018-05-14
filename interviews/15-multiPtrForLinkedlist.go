package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// FindKthToTail return the node that Kth to tail of linkedlist
// k starts at 1
func FindKthToTail(head *datastructure.Node, k uint) *datastructure.Node {
	if head == nil || k == 0 {
		return nil
	}
	behid, ahead := head, head

	for i := uint(0); i < k-1; i++ {
		if ahead.Next == nil {
			return nil
		}
		ahead = ahead.Next
	}

	for ahead.Next != nil {
		ahead = ahead.Next
		behid = behid.Next
	}

	return behid
}

// FindMidNode returns the middle node of linkedlist
func FindMidNode(head *datastructure.Node) *datastructure.Node {
	if head == nil {
		return nil
	}

	ahead, behind := head, head
	for ahead.Next != nil && ahead.Next.Next != nil {
		ahead = ahead.Next.Next
		behind = behind.Next
	}
	return behind
}

// IsLoop returns if the linkedlist is loop
func IsLoop(head *datastructure.Node) bool {
	if head == nil {
		return false
	}

	ahead, behind := head, head

	for ahead.Next != nil && ahead.Next.Next != nil {
		ahead = ahead.Next.Next
		behind = behind.Next
		if ahead == behind {
			return true
		}
	}
	return false
}
