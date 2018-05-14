package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// HasSubtree return true if r2 is a subtree of r1
func HasSubtree(r1, r2 *datastructure.TreeNode) bool {
	if r1 == nil || r2 == nil {
		return false
	}
	has := false
	if r1.Key == r2.Key {
		has = isSubtree(r1, r2)
	}
	if !has {
		has = HasSubtree(r1.LChild, r2)
	}
	if !has {
		has = HasSubtree(r1.RChild, r2)
	}

	return has
}

// isSubtree returns true r2 is a subtree of r1
func isSubtree(r1, r2 *datastructure.TreeNode) bool {
	// r2 reaches leaf
	if r2 == nil {
		return true
	}
	// r1 reaches leaf but r2 does not
	if r1 == nil {
		return false
	}

	if r1.Key != r2.Key {
		return false
	}
	// check recursively
	return isSubtree(r1.LChild, r2.RChild) && isSubtree(r1.RChild, r2.RChild)
}
