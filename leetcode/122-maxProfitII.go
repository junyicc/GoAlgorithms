package leetcode

import "math"

// Say you have an array prices for which the ith element is the price of a given stock on day i.
//
// Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).
//
// Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
//
// Example 1:
// Input: [7,1,5,3,6,4]
// Output: 7
// Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
//              Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
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
