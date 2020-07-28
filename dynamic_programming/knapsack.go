package dynamic_programming

func KnapsackWithWeight(items []int, w int) int {
	if len(items) < 1 {
		return 0
	}

	n := len(items)
	states := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		states[i] = make([]int, w+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j-items[i-1] < 0 {
				// cannot put into knapsack
				states[i][j] = states[i-1][j]
			} else {
				states[i][j] = max(states[i-1][j], states[i-1][j-items[i-1]]+items[i-1])
			}
		}
	}
	return states[n][w]
}

func KnapsackWithValue(itemWeight []int, itemValue []int, w int) int {
	if len(itemWeight) < 1 || len(itemValue) < 1 || w < 1 {
		return 0
	}

	n := len(itemWeight)
	states := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		states[i] = make([]int, w+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j-itemWeight[i-1] < 0 {
				// cannot put item into knapsack
				states[i][j] = states[i-1][j]
			} else {
				// do not put item into knapsack
				// put item into knapsack
				// choose the max value
				states[i][j] = max(states[i-1][j], states[i-1][j-itemWeight[i-1]]+itemValue[i-1])
			}
		}
	}
	return states[n][w]
}
