package leetcode

func missingNumber(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] == i || nums[i] == n {
			i++
		} else {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	return n
}

func missingNumberII(nums []int) (ret int) {
	n := len(nums)
	if n == 0 {
		return
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}

	return (n+1)*n/2 - sum
}
