package leetcode

// Given a circular array (the next element of the last element is the first element of the array), print the Next Greater Number for every element. The Next Greater Number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, output -1 for this number.
//
// Example 1:
// Input: [1,2,1]
// Output: [2,-1,2]
// Explanation: The first 1's next greater number is 2;
// The number 2 can't find next greater number;
// The second 1's next greater number needs to search circularly, which is also 2.

func nextGreaterElements(nums []int) []int {
	if len(nums) < 1 {
		return nil
	}
	n := len(nums)
	res := make([]int, n)
	stack := make([]int, 0, n)

	for i := 2*n - 1; i >= 0; i-- {
		// pop smaller element
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%n] {
			stack = stack[:len(stack)-1]
		}

		// record greater num
		if len(stack) == 0 {
			res[i%n] = -1
		} else {
			res[i%n] = stack[len(stack)-1]
		}

		// for next judgement
		stack = append(stack, nums[i%n])
	}
	return res
}
