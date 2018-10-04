package leetcode

func plusOne(digits []int) []int {
	if digits == nil || len(digits) < 1 {
		return []int{}
	}
	n := len(digits)
	carry := 0
	for i := n - 1; i >= 0; i-- {
		sum := digits[i] + carry
		if i == n-1 {
			sum = digits[i] + 1
		}
		if sum > 9 {
			sum -= 10
			digits[i] = sum
			carry = 1
		} else {
			digits[i] = sum
			carry = 0
			break
		}
	}

	var nums []int
	if carry != 0 {
		// overflow
		nums = make([]int, n+1)
		nums[0] = 1
		copy(nums[1:], digits)
	} else {
		nums = make([]int, n)
		copy(nums, digits)
	}
	return nums
}
