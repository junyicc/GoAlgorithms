package interviews

import (
	"fmt"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// FindPathofTree finds the path that sum of nodes equal to an expected number
func FindPathofTree(root *datastructure.TreeNode, sum int) {
	if root == nil {
		return
	}
	var curSum int
	path := datastructure.DynamicStack{}
	findPath(root, sum, &path, curSum)
}

func findPath(node *datastructure.TreeNode, sum int, path *datastructure.DynamicStack, curSum int) {
	path.Push(*node)
	curSum += node.Key

	isLeaf := node.LChild == nil && node.RChild == nil
	if isLeaf && curSum == sum {
		fmt.Println(path)
	}

	if node.LChild != nil {
		findPath(node.LChild, sum, path, curSum)
	}
	if node.RChild != nil {
		findPath(node.RChild, sum, path, curSum)
	}
	// delete current node before back to parent node
	path.Pop()
}
