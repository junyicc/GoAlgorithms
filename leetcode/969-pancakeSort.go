package leetcode

// 给定数组 A，我们可以对其进行煎饼翻转：我们选择一些正整数 k <= A.length，然后反转 A 的前 k 个元素的顺序。我们要执行零次或多次煎饼翻转（按顺序一次接一次地进行）以完成对数组 A 的排序。
//
// 返回能使 A 排序的煎饼翻转操作所对应的 k 值序列。任何将数组排序且翻转次数在 10 * A.length 范围内的有效答案都将被判断为正确。
//
// 示例 1：
// 输入：[3,2,4,1]
// 输出：[4,2,4,3]
// 解释：
// 我们执行 4 次煎饼翻转，k 值分别为 4，2，4，和 3。
// 初始状态 A = [3, 2, 4, 1]
// 第一次翻转后 (k=4): A = [1, 4, 2, 3]
// 第二次翻转后 (k=2): A = [4, 1, 2, 3]
// 第三次翻转后 (k=4): A = [3, 2, 1, 4]
// 第四次翻转后 (k=3): A = [1, 2, 3, 4]，此时已完成排序。
//
// 示例 2：
// 输入：[1,2,3]
// 输出：[]
// 解释：
// 输入已经排序，因此不需要翻转任何内容。
// 请注意，其他可能的答案，如[3，3]，也将被接受。

func pancakeSort(A []int) []int {
	if len(A) < 2 {
		return A
	}
	var res []int
	reversePancake(A, len(A), &res)
	return res
}

func reversePancake(A []int, n int, res *[]int) {
	if n == 1 {
		return
	}
	maxIdx := 0
	for i := 1; i < n; i++ {
		if A[i] > A[maxIdx] {
			maxIdx = i
		}
	}
	if maxIdx < n-1 {
		// reverse first maxIdx
		lo := 0
		hi := maxIdx
		for lo < hi {
			A[lo], A[hi] = A[hi], A[lo]
			lo++
			hi--
		}
		*res = append(*res, maxIdx+1)
		// reverse n
		lo, hi = 0, n-1
		for lo < hi {
			A[lo], A[hi] = A[hi], A[lo]
			lo++
			hi--
		}
		*res = append(*res, n)
	}

	reversePancake(A, n-1, res)
}
