package leetcode

import "sort"

// FourSum Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
func FourSum(nums []int, target int) [][]int {
	if nums == nil || len(nums) < 4 {
		return nil
	}
	n := len(nums)
	sort.Ints(nums)
	var quadruplets [][]int

	for i := 0; i < n-3; i++ {
		a := nums[i]
		if i > 0 && a == nums[i-1] {
			continue
		}
		if a+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if a+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		// three sum
		for j := i + 1; j < n-2; j++ {
			b := nums[j]
			if j > i+1 && b == nums[j-1] {
				continue
			}
			if a+b+nums[j+1]+nums[j+2] > target {
				break
			}
			if a+b+nums[n-2]+nums[n-1] < target {
				continue
			}
			lo, hi := j+1, n-1
			for lo < hi {
				c, d := nums[lo], nums[hi]
				if a+b+c+d == target {
					quadruplets = append(quadruplets, []int{a, b, c, d})
					for lo < hi && c == nums[lo] {
						lo++
					}

					for lo < hi && d == nums[hi] {
						hi--
					}
				} else if a+b+c+d > target {
					hi--
				} else {
					lo++
				}

			}
		}
	}
	return quadruplets
}
