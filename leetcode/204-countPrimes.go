package leetcode

/*
Count the number of prime numbers less than a non-negative number, n.

Example:
Input: 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
*/

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	notPrimes := make([]bool, n)
	cnt := 0
	for i := 2; i < n; i++ {
		if notPrimes[i] {
			continue
		}
		cnt++
		for j := i * i; j < n; j += i {
			notPrimes[j] = true
		}
	}
	return cnt
}
