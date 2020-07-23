package math

import (
	"math/rand"
	"time"
)

// 分析洗牌算法正确性的准则：产生的结果必须有 n! 种可能，否则就是错误的。

func shuffle(nums []int) {
	rand.Seed(time.Now().UnixNano())
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		// [0, i)
		// possibility: n*(n-1)*(n-2)*(n-3)*...*2*1
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}
}
