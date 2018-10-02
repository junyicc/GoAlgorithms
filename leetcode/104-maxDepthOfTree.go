package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	var depth int
	if leftDepth > rightDepth {
		depth = leftDepth + 1
	} else {
		depth = rightDepth + 1
	}
	return depth
}
