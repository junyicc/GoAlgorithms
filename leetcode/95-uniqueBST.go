package leetcode

// Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.
//
// Example:
//
// Input: 3
// Output:
// [
//   [1,null,3,2],
//   [3,2,null,1],
//   [3,1,null,null,2],
//   [2,1,3],
//   [1,null,2,null,3]
// ]
// Explanation:
// The above output corresponds to the 5 unique BST's shown below:
//
// 1         3     3      2      1
// \       /     /      / \      \
// 3     2     1      1   3      2
// /     /       \                 \
// 2     1         2                 3

func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return nil
	}
	return genUniqueBST(1, n)
}

func genUniqueBST(lo, hi int) []*TreeNode {
	trees := make([]*TreeNode, 0)
	if lo > hi {
		trees = append(trees, nil)
		return trees
	}
	// for every i, [1, i) as left subtree, (i, n] as right subtree
	for i := lo; i <= hi; i++ {
		lefts := genUniqueBST(lo, i-1)
		rights := genUniqueBST(i+1, hi)

		// merge tree
		for _, lChild := range lefts {
			for _, rChild := range rights {
				node := TreeNode{
					Val:   i,
					Left:  lChild,
					Right: rChild,
				}
				trees = append(trees, &node)
			}
		}
	}
	return trees
}
