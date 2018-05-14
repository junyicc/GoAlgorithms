package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// BiTreeMirror returns the mirror of a binary tree
func BiTreeMirror(node *datastructure.TreeNode) {
	if node == nil {
		return
	}

	if node.LChild != nil || node.RChild != nil {
		node.LChild, node.RChild = node.RChild, node.LChild

		if node.LChild != nil {
			BiTreeMirror(node.LChild)
		}

		if node.RChild != nil {
			BiTreeMirror(node.RChild)
		}
	}
}
