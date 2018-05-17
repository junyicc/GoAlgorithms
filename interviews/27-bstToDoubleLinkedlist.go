package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// BSTToDoubleLinkedlist converts bst to a sorted double linkedlist
func BSTToDoubleLinkedlist(root *datastructure.TreeNode) *datastructure.TreeNode {
	if root == nil || (root.LChild == nil && root.RChild == nil) {
		return root
	}
	var last *datastructure.TreeNode
	convert(root, &last)
	head := last
	for head != nil && head.LChild != nil {
		head = head.LChild
	}
	return head
}

func convert(node *datastructure.TreeNode, last **datastructure.TreeNode) {
	if node == nil {
		return
	}

	if node.LChild != nil {
		convert(node.LChild, last)
	}
	// connect precursor
	node.LChild = *last

	// connect successor
	if (*last) != nil {
		(*last).RChild = node
	}
	// change last
	*last = node

	if node.RChild != nil {
		convert(node.RChild, last)
	}
}
