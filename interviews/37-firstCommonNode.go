package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// FindFirstCommonNode returns the first common node of two linkedlist
func FindFirstCommonNode(h1, h2 *datastructure.Node) *datastructure.Node {
	if h1 == nil || h2 == nil {
		return nil
	}
	n1, n2 := h1, h2
	l1 := getLinkedlistLength(n1)
	l2 := getLinkedlistLength(n2)
	if l1 > l2 {
		step := l1 - l2
		for i := 0; i < step; i++ {
			n1 = n1.Next
		}
	} else {
		step := l2 - l1
		for i := 0; i < step; i++ {
			n2 = n2.Next
		}
	}
	for n1 != nil && n2 != nil && n1 != n2 {
		n1 = n1.Next
		n2 = n2.Next
	}
	return n1
}

func getLinkedlistLength(head *datastructure.Node) int {
	if head == nil {
		return 0
	}
	length := 0
	node := head
	for node != nil {
		length++
		node = node.Next
	}
	return length
}
