package interviews

func findRepeatNumber(nums []int) int {
	cnt := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if c, ok := cnt[nums[i]]; ok && c > 0 {
			return nums[i]
		}
		cnt[nums[i]]++
	}

	return -1
}
