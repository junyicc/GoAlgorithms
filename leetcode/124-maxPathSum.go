package leetcode

import "math"

// Given a non-empty binary tree, find the maximum path sum.

// For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.

// The path is not the path from root to leaf
// The path can start from any node in any direction

// The max path sum of one node = node.Val + path sum of left child + path sum of right child

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum := math.MinInt64
	maxPathSumOfOneNode(root, &maxSum)
	return maxSum
}

func maxPathSumOfOneNode(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}
	var curMaxPathSum int

	left := maxPathSumOfOneNode(node.Left, maxSum)
	right := maxPathSumOfOneNode(node.Right, maxSum)
	leftMax := max(left, 0)
	rightMax := max(right, 0)

	curMaxPathSum = node.Val + leftMax + rightMax
	if curMaxPathSum > *maxSum {
		*maxSum = curMaxPathSum
	}

	return node.Val + max(leftMax, rightMax)
}
