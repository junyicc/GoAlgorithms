package interviews

// 实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。
//
// 示例1:
// 输入: 2.00000, 10
// 输出: 1024.00000
//
// 示例2:
// 输入: 2.10000, 3
// 输出: 9.26100
//
// 示例3:
// 输入: 2.00000, -2
// 输出: 0.25000
// 解释: 2-2 = 1/22 = 1/4 = 0.25

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if x == 1 || n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	var negative bool
	if n < 0 {
		negative = true
		n = -n
	}

	res := myPow(x, n>>1)
	res *= res
	if n&1 == 1 {
		res *= x
	}

	if negative {
		res = 1.0 / res
	}
	return res
}
