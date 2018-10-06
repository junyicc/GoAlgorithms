package leetcode

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := Stack{}
	node := root
	for {
		goAlongLeft(&stack, node)
		if stack.IsEmpty() {
			break
		}
		node = (stack.Pop()).(*TreeNode)
		res = append(res, node.Val)
		node = node.Right
	}
	return res
}

func goAlongLeft(stack *Stack, node *TreeNode) {
	for node != nil {
		stack.Push(node)
		node = node.Left
	}
}
