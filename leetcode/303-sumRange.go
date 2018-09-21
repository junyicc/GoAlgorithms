package leetcode

/*
Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.
*/

// NumArray struct
type NumArray struct {
	nums []int
}

// Constructor returns NumArray object
func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))
	copy(sums, nums)
	l := len(sums)
	for i := 1; i < l; i++ {
		sums[i] += sums[i-1]
	}
	return NumArray{
		nums: sums,
	}
}

// SumRange returns the sum of numbers between i and j
func (na *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return na.nums[j]
	}
	return na.nums[j] - na.nums[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
