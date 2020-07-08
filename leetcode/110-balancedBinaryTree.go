package leetcode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	depth := 0
	return balancedTreeWithDepth(root, &depth)
}

func balancedTreeWithDepth(node *TreeNode, depth *int) bool {
	if node == nil {
		*depth = 0
		return true
	}

	var lDepth, rDepth int
	lBalanced := balancedTreeWithDepth(node.Left, &lDepth)
	rBalanced := balancedTreeWithDepth(node.Right, &rDepth)
	if lBalanced && rBalanced {
		if lDepth-rDepth >= 0 && lDepth-rDepth <= 1 {
			*depth = lDepth + 1
			return true
		} else if lDepth-rDepth >= -1 && lDepth-rDepth <= 0 {
			*depth = rDepth + 1
			return true
		}
	}
	return false
}
