package leetcode

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
