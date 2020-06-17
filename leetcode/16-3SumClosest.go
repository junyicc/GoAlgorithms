package leetcode

import (
	"math"
	"sort"
)

// ThreeSumClosest Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers.
// You may assume that each input would have exactly one solution.
func ThreeSumClosest(nums []int, target int) int {
	if len(nums) <= 3 {
		return sumOfSlice(nums)
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
				for lo < hi && nums[hi] == c {
					hi--
				}
			} else {
				for lo < hi && nums[lo] == b {
					lo++
				}
			}

		}
	}
	return minSum
}

func sumOfSlice(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}
