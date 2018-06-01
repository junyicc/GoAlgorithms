package interviews

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// GetLastCommonAncestorForBST returns the last common ancestor of two BST node
func GetLastCommonAncestorForBST(root, n1, n2 *datastructure.TreeNode) *datastructure.TreeNode {
	if root == nil || n1 == nil || n2 == nil {
		return nil
	}

	if (root.Key >= n1.Key && root.Key <= n2.Key) || (root.Key >= n2.Key && root.Key <= n1.Key) {
		return root
	} else if root.Key < n1.Key && root.Key < n2.Key {
		return GetLastCommonAncestorForBST(root.RChild, n1, n2)
	} else {
		return GetLastCommonAncestorForBST(root.LChild, n1, n2)
	}
}

// GetLastCommonAncestorForTree returns the last common ancestor of two TreeNode
// equal to get last common node of two list
func GetLastCommonAncestorForTree(root, n1, n2 *datastructure.TreeNode) *datastructure.TreeNode {
	if root == nil || n1 == nil || n2 == nil {
		return nil
	}
	var p1, p2 []*datastructure.TreeNode
	getNodePath(root, n1, &p1)
	getNodePath(root, n2, &p2)

	node := getLastCommonNode(p1, p2)
	return node
}

func getNodePath(root, node *datastructure.TreeNode, path *[]*datastructure.TreeNode) bool {
	found := false
	*path = append(*path, root)
	if root == node {
		return true
	}
	// children
	if !found && root.LChild != nil {
		found = getNodePath(root.LChild, node, path)
	}
	if !found && root.RChild != nil {
		found = getNodePath(root.RChild, node, path)
	}

	if !found {
		*path = (*path)[:len(*path)-1]
	}
	return found
}

func getLastCommonNode(l1, l2 []*datastructure.TreeNode) *datastructure.TreeNode {
	if l1 == nil || l2 == nil || len(l1) < 1 || len(l2) < 1 {
		return nil
	}
	var last *datastructure.TreeNode
	for i, j := 0, 0; i < len(l1) && j < len(l2); {
		if l1[i] != l2[j] {
			break
		} else {
			last = l1[i]
		}
		i++
		j++
	}
	return last
}
