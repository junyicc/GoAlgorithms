package leetcode

import (
	"container/list"
	"math"
)

// Given a binary tree, determine if it is a valid binary search tree (BST).
//
// Assume a BST is defined as follows:
//
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := list.New()
	node := root
	lastVal := math.MinInt64
	for {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}

		if stack.Len() == 0 {
			break
		}

		node = stack.Remove(stack.Back()).(*TreeNode)
		if node.Val <= lastVal {
			return false
		}
		lastVal = node.Val
		node = node.Right
	}

	return true
}
