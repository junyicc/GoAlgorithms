package leetcode

func singleNumber(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}

	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}
