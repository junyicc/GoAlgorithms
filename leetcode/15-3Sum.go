package leetcode

import (
	"sort"
)

// ThreeSum Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
func ThreeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}

	var triplets [][]int
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] {
			continue
		}
		if a+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if a+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		lo := i + 1
		hi := n - 1
		for lo < hi {
			b := nums[lo]
			c := nums[hi]
			s := a + b + c
			if s == 0 {
				triplets = append(triplets, []int{a, b, c})

				for lo < hi && nums[lo] == b {
					lo++
				}
				for lo < hi && nums[hi] == c {
					hi--
				}
			} else if s > 0 {
				hi--
			} else {
				lo++
			}
		}
	}
	return triplets
}
