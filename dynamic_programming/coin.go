package dynamic_programming

import "math"

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
