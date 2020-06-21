package leetcode

import "container/list"

// Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.
//
// Note: A leaf is a node with no children.
//
// Example:
//
// Given the below binary tree and sum = 22,
//
// 5
// / \
// 4   8
// /   / \
// 11  13  4
// /  \      \
// 7    2      1
// return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	stack := list.New()
	return hasPathInSubTree(root, sum, stack, 0)
}

func hasPathInSubTree(node *TreeNode, sum int, stack *list.List, curSum int) bool {
	if node == nil {
		return false
	}

	curSum += node.Val
	stack.PushBack(node)

	isLeaf := node.Left == nil && node.Right == nil
	if curSum == sum && isLeaf {
		return true
	}

	// find in left or right subtree
	var left, right bool
	if node.Left != nil {
		left = hasPathInSubTree(node.Left, sum, stack, curSum)
	}
	if node.Right != nil {
		right = hasPathInSubTree(node.Right, sum, stack, curSum)
	}

	stack.Remove(stack.Back())

	return left || right
}
