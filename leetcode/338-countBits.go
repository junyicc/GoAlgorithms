package leetcode

/*
Given a non negative integer number num. For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.

Example 1:
Input: 2
Output: [0,1,1]

Example 2:
Input: 5
Output: [0,1,1,2,1,2]
*/

func countBits(num int) []int {
	if num < 0 {
		return []int{}
	}

	bits := make([]int, num+1)
	bits[0] = 0
	power2 := []int{}
	lastPower2 := len(power2) - 1
	for i := 1; i <= num; i++ {
		cnt := 0
		if isPowerOf2(i) {
			cnt = 1
			power2 = append(power2, i)
			lastPower2++
		} else {
			cnt = bits[i-power2[lastPower2]] + 1
		}
		bits[i] = cnt
	}
	return bits
}

func isPowerOf2(n int) bool {
	if n > 0 {
		return n&(n-1) == 0
	}
	return false
}

func countBits2(num int) []int {
	if num < 0 {
		return []int{}
	}
	bits := make([]int, num+1)
	lastPower2 := 1
	for i := 1; i <= num; i++ {
		if i&(i-1) == 0 {
			lastPower2 = i
			bits[i] = 1
		} else {
			bits[i] = 1 + bits[i%lastPower2]
		}
	}
	return bits
}
