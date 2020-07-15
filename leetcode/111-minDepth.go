package leetcode

// Given a binary tree, find its minimum depth.
//
// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
//
// Note: A leaf is a node with no children.
//
// Example:
//
// Given binary tree [3,9,20,null,null,15,7],
//
// 3
// / \
// 9  20
// /  \
// 15   7
// return its minimum depth = 2.

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := make([]*TreeNode, 0)
	visited := make(map[*TreeNode]bool)
	q = append(q, root)
	visited[root] = true
	depth := 1

	for len(q) > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			// dequeue
			node := q[0]
			q = q[1:]

			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				if !visited[node.Left] {
					q = append(q, node.Left)
					visited[node.Left] = true
				}
			}
			if node.Right != nil {
				if !visited[node.Right] {
					q = append(q, node.Right)
					visited[node.Right] = true
				}
			}
		}
		depth++
	}
	return depth
}
