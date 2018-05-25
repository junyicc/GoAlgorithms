package interviews

// FindMaxSumOfSubArray returns the max sum of sub array
func FindMaxSumOfSubArray(arr []int) (int, bool) {
	if arr == nil || len(arr) < 1 {
		return 0, false
	}

	maxSum, curSum := ^int(^uint(0)>>1), 0
	for _, e := range arr {
		if curSum <= 0 {
			curSum = e
		} else {
			curSum += e
		}
		if maxSum < curSum {
			maxSum = curSum
		}
	}
	return maxSum, true
}

// FindSubArrayWithMaxSum returns the sub array with max sum
func FindSubArrayWithMaxSum(arr []int) []int {
	if arr == nil || len(arr) < 1 {
		return nil
	}
	sub := []int{}
	begin := 0
	maxSum, curSum := ^int(^uint(0)>>1), 0
	for i, e := range arr {
		if curSum <= 0 {
			begin = i
			curSum = e
			sub = []int{e}
		} else {
			curSum += e
		}
		if maxSum < curSum {
			maxSum = curSum
			sub = arr[begin : i+1]
		}
	}
	return sub
}
