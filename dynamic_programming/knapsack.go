package dynamic_programming

func KnapsackWithWeight(items []int, w int) int {
	n := len(items)
	if n < 1 {
		return 0
	}
	states := make([][]bool, n)
	for i := 0; i < n; i++ {
		states[i] = make([]bool, w+1)
	}

	// init state
	states[0][0] = true
	states[0][items[0]] = true

	for i := 1; i < n; i++ {
		// do no put item i into knapsack
		for j := 0; j < w+1; j++ {
			if states[i-1][j] {
				states[i][j] = states[i-1][j]
			}
		}
		// put item i into knapsack
		for j := 0; j < (w + 1 - items[i]); j++ {
			if states[i-1][j] {
				states[i][j+items[i]] = true
			}
		}
	}

	for j := w; j >= 0; j-- {
		if states[n-1][j] {
			return j
		}
	}
	return 0
}

func KnapsackWithValue(itemWeight []int, itemValue []int, w int) int {
	if len(itemWeight) < 1 || len(itemValue) < 1 || len(itemWeight) != len(itemValue) {
		return 0
	}
	n := len(itemWeight)

	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, w+1)
		for j := 0; j < w+1; j++ {
			states[i][j] = -1
		}
	}

	// init state
	states[0][0] = 0
	states[0][itemWeight[0]] = itemValue[0]

	for i := 1; i < n; i++ {
		// do no put item i into knapsack
		for j := 0; j < w+1; j++ {
			if states[i-1][j] > 0 {
				states[i][j] = states[i-1][j]
			}
		}

		// put item i into knapsack
		for j := 0; j < w+1-itemWeight[i]; j++ {
			if states[i-1][j] >= 0 {
				v := states[i-1][j] + itemValue[i]
				if v > states[i][j+itemWeight[i]] {
					states[i][j+itemWeight[i]] = v
				}
			}
		}
	}
	maxValue := states[n-1][w]
	for j := w - 1; j >= 1; j-- {
		if states[n-1][j] > maxValue {
			maxValue = states[n-1][j]
		}
	}
	return maxValue
}

func Knapsack(itemWeight []int, itemValue []int, w int) int {
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
