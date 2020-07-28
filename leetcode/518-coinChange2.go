package leetcode

// You are given coins of different denominations and a total amount of money. Write a function to compute the number of combinations that make up that amount. You may assume that you have infinite number of each kind of coin.
//
// Example 1:
// Input: amount = 5, coins = [1, 2, 5]
// Output: 4
// Explanation: there are four ways to make up the amount:
// 5=5
// 5=2+2+1
// 5=2+1+1+1
// 5=1+1+1+1+1
//
// Example 2:
// Input: amount = 3, coins = [2]
// Output: 0
// Explanation: the amount of 3 cannot be made up just with coins of 2.
//
// Example 3:
// Input: amount = 10, coins = [10]
// Output: 1

func change(amount int, coins []int) int {
	// complete knapsack problem:
	// 	- amount weight knapsack
	// 	- coins is the weight of each item, assuming that we have infinity items

	n := len(coins)
	// states[i][j]: with weight j knapsack and the first i items, how many choices we have
	states := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		states[i] = make([]int, amount+1)
	}

	// init base cases
	for i := 0; i <= n; i++ {
		states[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] < 0 {
				// cannot put into knapsack
				states[i][j] = states[i-1][j]
			} else {
				// add choices of both
				// - do not put item i into knapsack
				// - put item i into knapsack
				states[i][j] = states[i-1][j] + states[i][j-coins[i-1]]
			}
		}
	}
	return states[n][amount]
}
