package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// FindEntryOfLoop returns the entry node of loop linkedlist
func FindEntryOfLoop(head *datastructure.ListNode) *datastructure.ListNode {
	mi := findMeetingNode(head)
	if mi == nil {
		return nil
	}
	node := mi.Next
	cnt := 1
	for node != nil && node != mi {
		node = node.Next
		cnt++
	}
	behind, ahead := head, head
	i := 0
	for ; i < cnt && ahead != nil; i++ {
		ahead = ahead.Next
	}
	if i == cnt {
		for behind != ahead {
			behind = behind.Next
			ahead = ahead.Next
		}
		return ahead
	}
	return nil
}

func findMeetingNode(head *datastructure.ListNode) *datastructure.ListNode {
	if head == nil {
		return nil
	}
	var lo, hi *datastructure.ListNode
	lo = head
	if lo.Next == nil {
		return nil
	}
	hi = lo.Next.Next

	for lo != nil && hi != nil {
		if lo == hi {
			return lo
		}
		lo = lo.Next
		hi = hi.Next
		if hi != nil {
			hi = hi.Next
		}
	}
	return nil
}
