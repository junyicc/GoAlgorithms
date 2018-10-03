package leetcode

func containsDuplicate(nums []int) bool {
	n := len(nums)
	if nums == nil || n < 1 {
		return false
	}
	numsMap := make(map[int]bool)
	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			return true
		}
		numsMap[num] = true
	}
	return false
}
