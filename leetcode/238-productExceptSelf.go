package leetcode

func productExceptSelf(nums []int) []int {
	if nums == nil || len(nums) < 1 {
		return nil
	}
	n := len(nums)
	res := make([]int, n)
	tmp := 1
	for i := 0; i < n; i++ {
		res[i] = tmp
		tmp *= nums[i]
	}
	tmp = 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= tmp
		tmp *= nums[i]
	}
	return res
}
