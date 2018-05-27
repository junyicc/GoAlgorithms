package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// TreeDepth return binary tree depth
func TreeDepth(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := TreeDepth(root.LChild)
	rDepth := TreeDepth(root.RChild)
	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

// IsBalanceTree returns true if the depth difference of subtrees is no more than 1
func IsBalanceTree(root *datastructure.TreeNode) bool {
	if root == nil {
		return false
	}
	var depth int
	return isBalance(root, &depth)
}

func isBalance(root *datastructure.TreeNode, depth *int) bool {
	if root == nil {
		*depth = 0
		return true
	}

	var left, right int
	if isBalance(root.LChild, &left) && isBalance(root.RChild, &right) {
		if left-right >= 0 && left-right <= 1 {
			*depth = left + 1
			return true
		} else if right-left >= 0 && right-left <= 1 {
			*depth = right + 1
			return true
		}
	}
	return false
}
