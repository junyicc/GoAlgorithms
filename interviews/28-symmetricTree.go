package interviews

// 请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//     1
//    / \
//   2   2
//    \   \
//    3    3
//
//
//
// 示例 1：
// 输入：root = [1,2,2,3,4,4,3]
// 输出：true
//
// 示例 2：
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricCore(root.Left, root.Right)
}

func isSymmetricCore(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	if node1.Val != node2.Val {
		return false
	}

	return isSymmetricCore(node1.Left, node2.Right) && isSymmetricCore(node1.Right, node2.Left)
}
