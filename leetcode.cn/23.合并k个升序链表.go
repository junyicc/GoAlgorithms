package leetcodecn

import (
	"container/heap"
)

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	lo, hi := 0, len(lists)
	return mergeListsRecur(lists, lo, hi)
}

func mergeListsRecur(lists []*ListNode, lo, hi int) *ListNode {
	if hi-lo < 2 {
		return lists[lo]
	}

	mi := (lo + hi) >> 1
	left := mergeListsRecur(lists, lo, mi)
	right := mergeListsRecur(lists, mi, hi)

	return mergeTwoLists(left, right)
}

// heap of ListNode
type ListNodeHeap struct {
	Data []*ListNode
	Cmp  func(i, j int) bool
}

func (h *ListNodeHeap) Len() int           { return len(h.Data) }
func (h *ListNodeHeap) Less(i, j int) bool { return h.Cmp(i, j) }
func (h *ListNodeHeap) Swap(i, j int)      { h.Data[i], h.Data[j] = h.Data[j], h.Data[i] }
func (h *ListNodeHeap) Push(x interface{}) { h.Data = append(h.Data, x.(*ListNode)) }
func (h *ListNodeHeap) Pop() interface{} {
	n := h.Len()
	old := h.Data[n-1]
	h.Data = h.Data[:n-1]
	return old
}

func mergeListsByHeap(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	// init heap
	h := ListNodeHeap{
		Data: make([]*ListNode, 0, len(lists)),
	}
	h.Cmp = func(i, j int) bool {
		return h.Data[i].Val < h.Data[j].Val
	}
	heap.Init(&h)

	d := new(ListNode)
	n := d
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		heap.Push(&h, lists[i])
	}

	for h.Len() > 0 {
		minNode := heap.Pop(&h).(*ListNode)
		if minNode.Next != nil {
			heap.Push(&h, minNode.Next)
		}

		n.Next = minNode
		n = n.Next
	}
	return d.Next
}

// @lc code=end
