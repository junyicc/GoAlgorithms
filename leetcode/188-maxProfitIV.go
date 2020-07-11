package leetcode

import "math"

func maxProfitIV(maxK int, prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	n := len(prices)
	states := make([][][2]int, n+1)
	for i := 0; i <= n; i++ {
		states[i] = make([][2]int, maxK+1)
	}

	// init states
	for k := 0; k <= maxK; k++ {
		// no transaction
		states[0][k][0] = 0
		states[0][k][1] = math.MinInt32
	}

	for i := 1; i <= n; i++ {
		for k := 1; k <= maxK; k++ {
			states[i][k][0] = max(states[i-1][k][0], states[i-1][k][1]+prices[i-1])
			states[i][k][1] = max(states[i-1][k][1], states[i-1][k-1][0]-prices[i-1])
		}
	}
	return states[n][maxK][0]
}
