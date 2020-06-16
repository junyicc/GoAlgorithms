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

func FindPathToTreeNode(root *datastructure.TreeNode, dst *datastructure.TreeNode, f func(stack *datastructure.DynamicStack)) {
	if root == nil || dst == nil {
		return
	}

	stack := datastructure.NewDynamicStack()
	findPathToTreeNode(root, dst, stack, f)
}

func findPathToTreeNode(node *datastructure.TreeNode, dst *datastructure.TreeNode, stack *datastructure.DynamicStack, f func(stack *datastructure.DynamicStack)) {
	stack.Push(node.Data)

	if datastructure.Equal(stack.GetTop(), dst.Data) {
		f(stack)
		return
	}

	if node.LChild != nil {
		findPathToTreeNode(node.LChild, dst, stack, f)
	}
	if node.RChild != nil {
		findPathToTreeNode(node.RChild, dst, stack, f)
	}

	stack.Pop()
}
