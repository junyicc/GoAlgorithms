package leetcode

// Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.
//
// Basically, the deletion can be divided into two stages:
//
// Search for a node to remove.
// If the node is found, delete the node.

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {

		// root.Val == key
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// have both left child and right child
			p := root
			node := root.Left
			for node.Right != nil {
				p = node
				node = node.Right
			}

			root.Val = node.Val

			if p == root {
				p.Left = node.Left
			} else {
				p.Right = node.Left
			}
			return root
		}
	}
	return root
}
