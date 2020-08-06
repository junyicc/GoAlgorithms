package interviews

// 请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数。例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
//
// 示例 1：
// 输入：00000000000000000000000000001011
// 输出：3
// 解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。

func hammingWeight(num uint32) int {
	var cnt int
	for num != 0 {
		cnt++
		num = num & (num - 1)
	}
	return cnt
}

// IsPowerOf2 return true if n is the power of 2
func IsPowerOf2(n int) bool {
	if n > 0 {
		return n&(n-1) == 0
	}
	return false
}

// NumOfChangingBits returns number of bits that are needed to be changed from m to n
func NumOfChangingBits(m, n int) int {
	diff := m ^ n
	return hammingWeight(uint32(diff))
}

// IsOdd return true if n is an odd number
func IsOdd(n int) bool {
	if n <= 0 {
		return false
	}
	return uint(n)&1 == 1
}
