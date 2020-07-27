package leetcode

import "math"

// Your are given an array of integers prices, for which the i-th element is the price of a given stock on day i; and a non-negative integer fee representing a transaction fee.
//
// You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction. You may not buy more than 1 share of a stock at a time (ie. you must sell the stock share before you buy again.)
//
// Return the maximum profit you can make.
//
// Example 1:
// Input: prices = [1, 3, 2, 8, 4, 9], fee = 2
// Output: 8
// Explanation: The maximum profit can be achieved by:
// Buying at prices[0] = 1
// Selling at prices[3] = 8
// Buying at prices[4] = 4
// Selling at prices[5] = 9
// The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.

func maxProfitWithFee(prices []int, fee int) int {
	if len(prices) < 1 {
		return 0
	}

	n := len(prices)
	states := make([][2]int, n+1)
	// init base cases
	states[0][0] = 0
	states[0][1] = math.MinInt32

	for i := 1; i <= n; i++ {
		states[i][0] = max(states[i-1][0], states[i-1][1]+prices[i-1])
		states[i][1] = max(states[i-1][1], states[i-1][0]-prices[i-1]-fee)
	}
	return states[n][0]
}
