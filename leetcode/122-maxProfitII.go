package leetcode

import "math"

func maxProfitIIWithDp(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	n := len(prices)
	states := make([][2]int, n+1)

	// init state
	states[0][0] = 0
	states[0][1] = math.MinInt32
	for i := 1; i <= n; i++ {
		states[i][0] = max(states[i-1][0], states[i-1][1]+prices[i-1])
		states[i][1] = max(states[i-1][1], states[i-1][0]-prices[i-1])
	}
	return states[n][0]
}

// logic: d-a = (b-a) + (c-b) + (d-c)
func maxProfitII(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	profit := 0
	n := len(prices)
	for i := 1; i < n; i++ {
		profit += max(prices[i]-prices[i-1], 0)
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
