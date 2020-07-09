package leetcode

// You are given two arrays (without duplicates) nums1 and nums2 where nums1’s elements are subset of nums2. Find all the next greater numbers for nums1's elements in the corresponding places of nums2.
//
// The Next Greater Number of a number x in nums1 is the first greater number to its right in nums2. If it does not exist, output -1 for this number.

// Example 1:
// Input: nums1 = [4,1,2], nums2 = [1,3,4,2].
// Output: [-1,3,-1]
// Explanation:
// For number 4 in the first array, you cannot find the next greater number for it in the second array, so output -1.
// For number 1 in the first array, the next greater number for it in the second array is 3.
// For number 2 in the first array, there is no next greater number for it in the second array, so output -1.

// Example 2:
// Input: nums1 = [2,4], nums2 = [1,2,3,4].
// Output: [3,-1]
// Explanation:
// For number 2 in the first array, the next greater number for it in the second array is 3.
// For number 4 in the first array, there is no next greater number for it in the second array, so output -1.

// Using mono-stack to find next greater element in O(n)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) < 1 || len(nums2) < 1 {
		return nil
	}

	n1, n2 := len(nums1), len(nums2)
	stack := make([]int, 0, n2)
	next := make([]int, n2)
	for i := n2 - 1; i >= 0; i-- {
		// pop smaller element
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		// record next greater element
		if len(stack) == 0 {
			next[i] = -1
		} else {
			next[i] = stack[len(stack)-1]
		}
		// for next judgement
		stack = append(stack, nums2[i])
	}

	res := make([]int, 0, n1)
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if nums1[i] == nums2[j] {
				res = append(res, next[j])
			}
		}
	}
	return res
}
