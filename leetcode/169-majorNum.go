package leetcode

func majorityElement(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	times := 1
	num := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if times == 0 {
			num = nums[i]
			times++
		} else if nums[i] != num {
			times--
		} else if nums[i] == num {
			times++
		}
	}
	if times > 0 {
		return num
	}
	return 0
}
