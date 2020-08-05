package interviews

import "github.com/CAIJUNYI/GoAlgorithms/datastructure"

// treeNode struct
type treeNode struct {
	Data   datastructure.Elem
	LChild *treeNode
	RChild *treeNode
	Parent *treeNode
}

// Succeesor returns the next node in the in-order sequence
func Succeesor(node *treeNode) *treeNode {
	if node == nil {
		return nil
	}
	var successor *treeNode
	if node.RChild != nil {
		child := node.RChild
		for child.LChild != nil {
			child = child.LChild
		}
		successor = child
	} else if node.Parent != nil {
		if node.Parent.LChild == node {
			successor = node.Parent
		} else {
			child := node
			ancestor := node.Parent
			for ancestor.LChild != child {
				child = ancestor
				ancestor = ancestor.Parent
			}
			successor = ancestor
		}
	}
	return successor
}
