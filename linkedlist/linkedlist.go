package linkedlist

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

//ReverseLinkedList reverse linkedlist in O(n)
func ReverseLinkedList(l *datastructure.LinkedList) {
	if l.Head() == nil || l.Head().Next == nil || l.Head().Next.Next == nil {
		return
	}
	head := l.Head()
	cur := head.Next
	var pre *datastructure.ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
	head.Next = pre
}

//HasCycle returns true if linkedlist has cycle
func HasCycle(l *datastructure.LinkedList) bool {
	if l.Head() == nil || l.Head().Next == nil || l.Head().Next.Next == nil {
		return false
	}
	head := l.Head()
	slow := head.Next
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

//MergeSortedList merge two sorted linkedlist
func MergeSortedList(l1, l2 *datastructure.LinkedList) *datastructure.LinkedList {
	if l1 == nil || l1.Head() == nil || l1.Head().Next == nil {
		return l2
	}
	if l2 == nil || l2.Head() == nil || l2.Head().Next == nil {
		return l1
	}
	n1 := l1.Head().Next
	n2 := l2.Head().Next
	l := datastructure.NewLinkedList()
	cur := l.Head()
	for n1 != nil && n2 != nil {
		if n1.Less(n2) {
			cur.Next = n1
			n1 = n1.Next
		} else {
			cur.Next = n2
			n2 = n2.Next
		}
		cur = cur.Next
	}

	if n1 != nil {
		cur.Next = n1
	} else if n2 != nil {
		cur.Next = n2
	}
	return l
}

//DeleteBottomN deletes the bottom n node
func DeleteBottomN(l *datastructure.LinkedList, n int) {
	if l == nil || l.Head() == nil || l.Head().Next == nil {
		return
	}
	if n < 1 || n > l.Len() {
		return
	}

	//find the bottom n node
	lo, hi := l.Head().Next, l.Head().Next
	for i := 0; i < n && hi != nil; i++ {
		hi = hi.Next
	}
	if hi == nil {
		return
	}
	for hi.Next != nil {
		lo = lo.Next
		hi = hi.Next
	}
	//delete node
	lo.Next = lo.Next.Next
}

//FindMiddleNode returns the middle node
func FindMiddleNode(l *datastructure.LinkedList) *datastructure.ListNode {
	if l == nil || l.Head() == nil || l.Head().Next == nil {
		return nil
	}
	slow, fast := l.Head().Next, l.Head().Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

//IsPalindrome returns true if the linkedlist is palindromic
func IsPalindrome(l *datastructure.LinkedList) bool {
	if l == nil || l.Head() == nil || l.Head().Next == nil {
		return false
	}
	//only one node
	firstNode := l.Head().Next
	if firstNode.Next == nil {
		return true
	}
	//fast go forward, slow go forward and reverse node
	slow, fast := firstNode, firstNode
	var pre *datastructure.ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		//reverse node
		next := slow.Next
		slow.Next = pre
		pre = slow
		slow = next
	}
	//linklist has been broken
	mid := slow

	isPalindrome := true
	var bw, fw *datastructure.ListNode
	bw = pre
	if fast == nil {
		//even length
		fw = mid
	} else {
		//odd length
		fw = mid.Next
	}
	for bw != nil && fw != nil {
		if !datastructure.Equal(bw.Data, fw.Data) {
			isPalindrome = false
			break
		}
		bw = bw.Next
		fw = fw.Next
	}
	if bw != nil || fw != nil {
		isPalindrome = false
	}

	//restore linkedlist
	node := pre
	pre = mid
	for node != nil {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}

	return isPalindrome
}

func PrintFromTail(h *datastructure.ListNode, f func(datastructure.Elem)) {
	if h == nil || h.Next == nil {
		return
	}

	stack := datastructure.NewDynamicStack()
	node := h.Next
	for node != nil {
		stack.Push(node.Data)
	}

	for !stack.IsEmpty() {
		data := stack.GetTop()
		f(data)
		stack.Pop()
	}
}

func TraverseFromTail(head *datastructure.ListNode, f func(node *datastructure.ListNode)) {
	if head == nil {
		return
	}

	TraverseFromTail(head.Next, f)

	f(head)
}
