package leetcode

import "math"

/*
Say you have an array for which the ith element is the price of a given stock on day i.
If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:
Input: [7,1,5,3,6,4]
Output: 5

Example 2:
Input: [7,6,4,3,1]
Output: 0
*/

func maxProfitWithDp(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	n := len(prices)
	states := make([][2]int, n+1)
	states[0][0] = 0
	states[0][1] = math.MinInt32

	for i := 1; i <= n; i++ {
		states[i][0] = max(states[i-1][0], states[i-1][1]+prices[i-1])
		// states[i][1] = max(states[i-1][1], states[i-1][0]-prices[i-1])
		// at most once transaction
		states[i][1] = max(states[i-1][1], -prices[i-1])
	}
	return states[n][0]
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	n := len(prices)
	min := prices[0]
	profits := make([]int, n)
	for i := 1; i < n; i++ {
		if prices[i] < min {
			min = prices[i]
		}
		profits[i] = prices[i] - min
	}

	max := 0
	for _, p := range profits {
		if p > max {
			max = p
		}
	}
	return max
}

func maxProfitI(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}
	maxProfit, maxCur := 0, 0
	n := len(prices)
	for i := 1; i < n; i++ {
		maxCur += prices[i] - prices[i-1]
		maxCur = max(maxCur, 0)
		maxProfit = max(maxProfit, maxCur)
	}
	return maxProfit
}
