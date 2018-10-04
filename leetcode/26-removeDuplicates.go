package leetcode

func removeDuplicates(nums []int) int {
	n := len(nums)
	if nums == nil || n < 1 {
		return 0
	}
	if n < 2 {
		return n
	}
	i, j := 0, 1
	for j < n {
		if nums[j] != nums[i] {
			nums[i+1] = nums[j]
			i++
		}
		j++
	}
	nums = nums[:i+1]
	return i + 1
}
