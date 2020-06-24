package leetcode

import "math"

// You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.
//
// Example 1:
// Input: coins = [1, 2, 5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1
//
// Example 2:
// Input: coins = [2], amount = 3
// Output: -1

// f(i) = min(f(i-coin_j)) + 1  where coin_j is the jth coin of the coins
// calculate all possible min counts of coin_j for the number from coin_j to amount

func coinChange(coins []int, amount int) int {
	if len(coins) < 1 || amount < 1 {
		return 0
	}

	count := make([]int, amount+1)
	// init state
	for i := 0; i < amount+1; i++ {
		count[i] = math.MaxInt32
	}
	count[0] = 0
	for _, coin := range coins {
		for i := coin; i < amount+1; i++ {
			count[i] = min(count[i], count[i-coin]+1)
		}
	}
	if count[amount] == math.MaxInt32 {
		return -1
	}
	return count[amount]
}
