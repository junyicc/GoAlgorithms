package leetcode

import (
	"math"
)

// Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
//
// Example:
//
// Input:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// Output: 1->1->2->3->4->4->5->6

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}

	return mergeLists(lists, 0, len(lists))
}

func mergeLists(lists []*ListNode, lo, hi int) *ListNode {
	if hi-lo < 2 {
		return lists[lo]
	}

	mi := lo + (hi-lo)>>1
	left := mergeLists(lists, lo, mi)
	right := mergeLists(lists, mi, hi)

	return mergeTwoLists(left, right)
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	head := new(ListNode)
	node := head
	for {
		minIndex := getMinIndexOfLists(lists)
		if minIndex == -1 {
			break
		}
		node.Next = lists[minIndex]
		lists[minIndex] = lists[minIndex].Next
		node = node.Next
	}
	return head.Next
}

func getMinIndexOfLists(lists []*ListNode) int {
	min, minIndex := math.MaxInt64, -1
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil && lists[i].Val < min {
			min = lists[i].Val
			minIndex = i
		}
	}
	return minIndex
}
