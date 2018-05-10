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
