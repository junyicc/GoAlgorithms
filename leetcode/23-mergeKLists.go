package leetcode

import (
	"honnef.co/go/tools/simple"
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
