package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// FindKthNodeInBST returns the kth node in the in-order sequence of BST
func FindKthNodeInBST(root *datastructure.TreeNode, k int) *datastructure.TreeNode {
	if root == nil || k < 1 {
		return nil
	}
	return findKthNodeInBST(root, &k)
}

func findKthNodeInBST(node *datastructure.TreeNode, k *int) *datastructure.TreeNode {
	var target *datastructure.TreeNode

	if node.LChild != nil {
		target = findKthNodeInBST(node.LChild, k)
	}

	if target == nil {
		if *k == 1 {
			target = node
		}
		*k--
	}

	if target == nil && node.RChild != nil {
		target = findKthNodeInBST(node.RChild, k)
	}

	return target
}
