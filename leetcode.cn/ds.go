package leetcodecn

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeList(vals []int) *ListNode {
	d := new(ListNode)
	n := d
	for _, val := range vals {
		cur := ListNode{
			Val: val,
		}
		n.Next = &cur
		n = n.Next
	}
	return d.Next
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
