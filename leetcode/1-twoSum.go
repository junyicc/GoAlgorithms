package leetcode

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
func twoSum(nums []int, target int) []int {
	numIndex := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		d := target - nums[i]
		if _, ok := numIndex[d]; ok {
			return []int{numIndex[d], i}
		}
		numIndex[nums[i]] = i
	}
	return nil
}
