package leetcode

import "math"

// There are n bulbs that are initially off. You first turn on all the bulbs. Then, you turn off every second bulb. On the third round, you toggle every third bulb (turning on if it's off or turning off if it's on). For the i-th round, you toggle every i bulb. For the n-th round, you only toggle the last bulb. Find how many bulbs are on after n rounds.
//
// Example:
//
// Input: 3
// Output: 1
// Explanation:
// At first, the three bulbs are [off, off, off].
// After first round, the three bulbs are [on, on, on].
// After second round, the three bulbs are [on, off, on].
// After third round, the three bulbs are [on, off, off].
//
// So you should return 1, because there is only one bulb is on.

// the remaining on bulb must be switch odd times
// only those with odd factors can be switch odd times
// like: 16: [1, 2, 4, 8, 16], because 4 repeat
// so the remaining on bulb must be 1(1*1), 4(2*2), 9(3*3), 16(4*4)

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
