package interviews

// MaxProfit actually returns the max diff between two numbers in the array
func MaxProfit(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	maxDiff := arr[1] - arr[0]
	min := arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i-1] < min {
			min = arr[i-1]
		}
		d := arr[i] - min
		if d > maxDiff {
			maxDiff = d
		}
	}
	if maxDiff > 0 {
		return maxDiff
	}
	return 0
}
