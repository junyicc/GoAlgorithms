package leetcode

import "math"

// Say you have an array for which the ith element is the price of a given stock on day i.
//
// Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:
//
// You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
// After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)
// Example:
//
// Input: [1,2,3,0,2]
// Output: 3
// Explanation: transactions = [buy, sell, cooldown, buy, sell]

func maxProfitWithCooldown(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	n := len(prices)
	states := make([][2]int, n+2)

	// init base cases
	states[0][1] = math.MinInt32
	states[1][1] = math.MinInt32

	for i := 2; i < n+2; i++ {
		states[i][0] = max(states[i-1][0], states[i-1][1]+prices[i-2])
		states[i][1] = max(states[i-1][1], states[i-2][0]-prices[i-2])
	}

	return states[n+1][0]
}
