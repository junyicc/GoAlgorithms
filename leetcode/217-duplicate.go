package leetcode

func containsDuplicate(nums []int) bool {
	if nums == nil || len(nums) < 2 {
		return false
	}
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

func containsDuplicateII(nums []int) bool {
	if nums == nil || len(nums) < 2 {
		return false
	}
	return partition(nums, 0, len(nums)-1)
}

func partition(nums []int, lo, hi int) bool {
	if hi <= lo {
		return false
	}
	val := nums[(lo+hi)>>1]
	left, right := lo, hi
	for left <= right {
		for left <= right && nums[left] < val {
			left++
		}
		for left <= right && nums[right] > val {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return right >= lo && left <= hi && nums[left] == nums[right] ||
		partition(nums, lo, right) || partition(nums, left, hi)
}
