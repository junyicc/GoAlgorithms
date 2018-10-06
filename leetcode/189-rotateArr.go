package leetcode

func rotate(nums []int, k int) {
	n := len(nums)
	if nums == nil || n < 1 || k < 1 || k == n {
		return
	}
	k = k % n
	reverseArr(nums, 0, n-k-1)
	reverseArr(nums, n-k, n-1)
	reverseArr(nums, 0, n-1)
}

func reverseArr(nums []int, lo, hi int) {
	for lo < hi {
		nums[lo] = nums[lo] ^ nums[hi]
		nums[hi] = nums[lo] ^ nums[hi]
		nums[lo] = nums[lo] ^ nums[hi]
		lo++
		hi--
	}
}
