package leetcode

import "math"

// Say you have an array for which the ith element is the price of a given stock on day i.
//
// Design an algorithm to find the maximum profit. You may complete at most two transactions.
//
// Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

// Example 1:
// Input: [3,3,5,0,0,3,1,4]
// Output: 6
// Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
//              Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
//
// Example 2:
// Input: [1,2,3,4,5]
// Output: 4
// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
//              Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
//              engaging multiple transactions at the same time. You must sell before buying again.
//
// Example 3:
// Input: [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transaction is done, i.e. max profit = 0.

// base case：
// dp[-1][k][0] = dp[i][0][0] = 0
// dp[-1][k][1] = dp[i][0][1] = -infinity
//
// 状态转移方程：
// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i]) // sell share the same k transaction
// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i]) // buy start a new transaction

func maxProfitIII(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxK := 2
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
