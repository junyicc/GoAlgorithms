package leetcode

/*
On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).
Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

Note:
cost will have a length in the range [2, 1000].
Every cost[i] will be an integer in the range [0, 999].
*/

func minCostClimbingStairs(cost []int) int {
	stairLen := len(cost)
	a, b := cost[0], cost[1]
	var c int
	for i := 2; i < stairLen; i++ {
		c = cost[i] + min(a, b)
		a = b
		b = c
	}
	return min(a, b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
