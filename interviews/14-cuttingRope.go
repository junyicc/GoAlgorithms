package interviews

// 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
//
// 示例 1：
// 输入: 2
// 输出: 1
// 解释: 2 = 1 + 1, 1 × 1 = 1

// 示例 2:
// 输入: 10
// 输出: 36
// 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
//
// 提示：
// 2 <= n <= 58

// f(n) = max(f(i) * f(n-i))  1<=i<=n-1
func cuttingRope(n int) int {

	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	states := make([]int, n+1)
	// base cases
	states[0] = 0
	states[1] = 1
	states[2] = 2
	states[3] = 3

	for i := 4; i <= n; i++ {
		var maxProduct int
		for j := 1; j <= i>>1; j++ {
			product := states[j] * states[i-j]
			if product > maxProduct {
				maxProduct = product
			}
		}
		states[i] = maxProduct
	}
	return states[n]
}
