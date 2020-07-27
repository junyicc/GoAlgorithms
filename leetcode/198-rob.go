package leetcode

/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:
Input: [1,2,3,1]
Output: 4

Example 2:
Input: [2,7,9,3,1]
Output: 12
*/

func rob(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}
	houseNum := len(nums)
	if houseNum == 1 {
		return nums[0]
	}
	vals := make([]int, houseNum)
	copy(vals, nums)
	for i := 0; i < houseNum; i++ {
		if i-2 >= 0 {
			if vals[i]+vals[i-2] > vals[i-1] {
				vals[i] = vals[i] + vals[i-2]
			} else {
				vals[i] = vals[i-1]
			}
		} else {
			if i-1 >= 0 && vals[i-1] > vals[i] {
				vals[i] = vals[i-1]
			}
		}
	}
	return vals[houseNum-1]
}

func rob2(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}
	houseNum := len(nums)
	rob, notRob := 0, 0
	for i := 0; i < houseNum; i++ {
		curRob := notRob + nums[i] // rob current house
		notRob = max(notRob, rob)  // max value of rob (i-1) house and notrob
		rob = curRob
	}
	return max(rob, notRob)
}

func rob3(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	n := len(nums)
	states := make([]int, n+2)
	for i := 2; i < n+2; i++ {
		states[i] = max(states[i-2]+nums[i-2], states[i-1])
	}
	return states[n+1]
}
