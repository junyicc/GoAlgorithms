package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// IsSymmertrical returns true if the left subtree and the right subtree are symmertrical
func IsSymmertrical(root *datastructure.TreeNode) bool {
	return isSymmertrical(root, root)
}

func isSymmertrical(n1, n2 *datastructure.TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 == nil || n2 == nil {
		return false
	}

	if !datastructure.Equal(n1.Data, n2.Data) {
		return false
	}

	return isSymmertrical(n1.LChild, n2.RChild) && isSymmertrical(n1.RChild, n2.LChild)
}

// IsSymmertrical2 returns true if the left subtree and the right subtree are symmertrical
func IsSymmertrical2(root *datastructure.TreeNode) bool {
	leftPreOrder := []datastructure.Elem{}
	visitLeft := func(e datastructure.Elem) {
		leftPreOrder = append(leftPreOrder, e)
	}
	traverseLeft(root, visitLeft)

	rightPreOrder := []datastructure.Elem{}
	visitRight := func(e datastructure.Elem) {
		rightPreOrder = append(rightPreOrder, e)
	}
	traverseRight(root, visitRight)

	for i, e := range leftPreOrder {
		if !datastructure.Equal(rightPreOrder[i], e) {
			return false
		}
	}
	return true
}

func traverseLeft(root *datastructure.TreeNode, f func(datastructure.Elem)) {
	if root == nil {
		f(nil)
	} else {
		f(root.Data)
	}
	if root != nil {
		traverseLeft(root.LChild, f)
		traverseLeft(root.RChild, f)
	}
}

func traverseRight(root *datastructure.TreeNode, f func(datastructure.Elem)) {
	if root == nil {
		f(nil)
	} else {
		f(root.Data)
	}
	if root != nil {
		traverseRight(root.RChild, f)
		traverseRight(root.LChild, f)
	}
}
