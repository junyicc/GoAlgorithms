package interviews

// 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
//
// 例如:
// 给定二叉树: [3,9,20,null,null,15,7],

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			// visit node
			res = append(res, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}

// level order in multiline
func levelOrderII(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		tmp := make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			// visit node
			tmp = append(tmp, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, tmp)
	}

	return res
}

// level order in Z
func levelOrderIII(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var stack [2][]*TreeNode
	stack[0] = append(stack[0], root)
	flag := 0
	for len(stack[0]) > 0 || len(stack[1]) > 0 {
		tmp := make([]int, 0)

		size := len(stack[flag])
		for i := 0; i < size; i++ {
			node := stack[flag][len(stack[flag])-1]
			stack[flag] = stack[flag][:len(stack[flag])-1]

			// visit node
			tmp = append(tmp, node.Val)

			if flag == 0 {
				if node.Left != nil {
					stack[1-flag] = append(stack[1-flag], node.Left)
				}
				if node.Right != nil {
					stack[1-flag] = append(stack[1-flag], node.Right)
				}
			} else {
				if node.Right != nil {
					stack[1-flag] = append(stack[1-flag], node.Right)
				}
				if node.Left != nil {
					stack[1-flag] = append(stack[1-flag], node.Left)
				}
			}
		}
		res = append(res, tmp)
		flag = 1 - flag
	}

	return res
}
