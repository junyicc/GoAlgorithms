package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) < 1 {
		return nil
	}
	return sortedArrToBBST(nums, 0, len(nums)-1)
}

func sortedArrToBBST(nums []int, lo, hi int) *TreeNode {
	var root TreeNode
	mi := (lo + hi) >> 1
	root.Val = nums[mi]

	if lo == hi {
		return &root
	}

	if mi > lo {
		root.Left = sortedArrToBBST(nums, lo, mi-1)
	}
	if mi < hi {
		root.Right = sortedArrToBBST(nums, mi+1, hi)
	}
	return &root
}
