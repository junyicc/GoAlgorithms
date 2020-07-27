package leetcode

// The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.
//
// Determine the maximum amount of money the thief can rob tonight without alerting the police.
//
// Example 1:
// Input: [3,2,3,null,3,null,1]
//
// 3
// / \
// 2   3
// \   \
// 3   1
//
// Output: 7
// Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
//
// Example 2:
// Input: [3,4,5,1,3,null,1]
//  3
// / \
// 4   5
// / \   \
// 1   3   1
//
// Output: 9
// Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.

func robOnTree(root *TreeNode) int {
	mem := make(map[*TreeNode]int)
	return robOnTreeCore(root, mem)
}

func robOnTreeCore(root *TreeNode, mem map[*TreeNode]int) int {
	if root == nil {
		return 0
	}

	if val, ok := mem[root]; ok {
		return val
	}

	// rob
	robIt := root.Val
	// rob next next level
	if root.Left != nil {
		robIt += robOnTreeCore(root.Left.Left, mem) + robOnTreeCore(root.Left.Right, mem)
	}
	if root.Right != nil {
		robIt += robOnTreeCore(root.Right.Left, mem) + robOnTreeCore(root.Right.Right, mem)
	}

	// not rob
	notRob := 0
	// rob next level
	if root.Left != nil {
		notRob += robOnTreeCore(root.Left, mem)
	}
	if root.Right != nil {
		notRob += robOnTreeCore(root.Right, mem)
	}

	mem[root] = max(robIt, notRob)

	return mem[root]
}
