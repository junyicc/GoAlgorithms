package leetcode

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

func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}
	n := len(prices)
	dp := make([]int, n)
	minPrice := prices[0]
	for i := 1; i < n; i++ {
		dp[i] = prices[i] - minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	maxDiff := dp[0]
	for i := 1; i < n; i++ {
		if dp[i] > maxDiff {
			maxDiff = dp[i]
		}
	}
	return maxDiff
}
