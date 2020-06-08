package dynamic_programming

func KnapsackWithWeight(items []int, w int) int {
	n := len(items)
	if n < 1 {
		return 0
	}
	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, w+1)
	}

	// init state
	states[0][0] = 1
	states[0][items[0]] = 1

	for i := 1; i < n; i++ {
		// do no put item i into knapsack
		for j := 0; j < w+1; j++ {
			if states[i-1][j] == 1 {
				states[i][j] = states[i-1][j]
			}
		}
		// put item i into knapsack
		for j := 0; j < (w + 1 - items[i]); j++ {
			if states[i-1][j] == 1 {
				states[i][j+items[i]] = 1
			}
		}
	}

	for j := w; j >= 0; j-- {
		if states[n-1][j] == 1 {
			return j
		}
	}
	return 0
}
