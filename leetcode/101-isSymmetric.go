package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(lNode, rNode *TreeNode) bool {
	if lNode == nil && rNode == nil {
		return true
	}
	if lNode == nil || rNode == nil {
		return false
	}
	if lNode.Val != rNode.Val {
		return false
	}
	return isSymmetricTree(lNode.Left, rNode.Right) && isSymmetricTree(lNode.Right, rNode.Left)
}
