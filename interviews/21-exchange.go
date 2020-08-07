package interviews

// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
//
// 示例：
//
// 输入：nums = [1,2,3,4]
// 输出：[1,3,2,4]
// 注：[3,1,2,4] 也是正确的答案之一。

func exchange(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	exchangeCore(nums, IsOdd)

	return nums
}

func exchangeCore(nums []int, f func(int) bool) {
	lo, hi := 0, len(nums)-1
	for {
		for lo <= hi && f(nums[lo]) {
			lo++
		}
		for lo <= hi && !f(nums[hi]) {
			hi--
		}

		if lo > hi {
			break
		}

		nums[lo], nums[hi] = nums[hi], nums[lo]
	}
}
