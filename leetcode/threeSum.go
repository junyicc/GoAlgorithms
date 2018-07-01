package leetcode

import (
	"math"
	"sort"
)

// ThreeSum Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
func ThreeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}

	var triplets [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] {
			continue
		}
		lo := i + 1
		hi := len(nums) - 1
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

// ThreeSumClosest Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.
func ThreeSumClosest(nums []int, target int) int {
	if nums == nil || len(nums) < 3 {
		return 0
	}
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	sort.Ints(nums)
	minDiff := math.MaxInt64
	minSum := 0
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] {
			continue
		}
		lo := i + 1
		hi := len(nums) - 1
		for lo < hi {
			b := nums[lo]
			c := nums[hi]
			s := a + b + c
			d := int(math.Abs(float64(s - target)))
			if d < minDiff {
				minDiff = d
				minSum = s
			}

			if s == target {
				return target
			} else if s > target {
				hi--
			} else {
				lo++
			}

		}
	}
	return minSum
}
