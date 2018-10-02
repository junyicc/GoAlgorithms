package leetcode

func moveZeroes(nums []int) {
	if nums == nil || len(nums) < 1 {
		return
	}
	idx := 0
	for i, n := range nums {
		if n != 0 {
			nums[idx], nums[i] = nums[i], nums[idx]
			idx++
		}
	}
}
