package leetcode

// Invert a binary tree.
//
// Example:
//
// Input:
// 4
// /   \
// 2     7
// / \   / \
// 1   3 6   9

// Output:
// 4
// /   \
// 7     2
// / \   / \
// 9   6 3   1

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}

	return root
}
