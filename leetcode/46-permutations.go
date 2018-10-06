package leetcode

func permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	res := [][]int{}
	f := func(nums []int) {
		res = append(res, nums)
	}
	arrPermutations(nums, 0, len(nums), f)
	return res
}

func arrPermutations(nums []int, start, n int, f func([]int)) {
	if start == n-1 {
		pNums := make([]int, n)
		copy(pNums, nums)
		f(pNums)
		return
	}

	for i := start; i < n; i++ {
		nums[start], nums[i] = nums[i], nums[start]
		arrPermutations(nums, start+1, n, f)
		nums[start], nums[i] = nums[i], nums[start]
	}
}
