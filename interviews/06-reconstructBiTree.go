package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// ReConstructBiTree using pre-order and in-order sequence
func ReConstructBiTree(preorder, inorder []int) *datastructure.BST {
	if preorder == nil || inorder == nil || len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}

	root := construct(preorder, inorder)
	tree := datastructure.BST{}
	tree.Root = root
	return &tree
}

func construct(preorder, inorder []int) *datastructure.TreeNode {
	root := datastructure.TreeNode{}
	rootVal := preorder[0]
	root.Key = rootVal
	root.Data = rootVal

	var rootIndex int
	var leftNodeNum, rightNodeNum int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			rootIndex = i
			break
		}
		leftNodeNum++
	}
	rightNodeNum = len(inorder) - leftNodeNum - 1
	if leftNodeNum > 0 {
		root.LChild = construct(preorder[1:leftNodeNum+1], inorder[:rootIndex])
	}

	if rightNodeNum > 0 {
		root.RChild = construct(preorder[leftNodeNum+1:], inorder[rootIndex+1:])
	}

	return &root
}
