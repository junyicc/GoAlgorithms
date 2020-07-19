package leetcode

import "math"

// 你将获得 K 个鸡蛋，并可以使用一栋从 1 到 N  共有 N 层楼的建筑。
//
// 每个蛋的功能都是一样的，如果一个蛋碎了，你就不能再把它掉下去。
//
// 你知道存在楼层 F ，满足 0 <= F <= N 任何从高于 F 的楼层落下的鸡蛋都会碎，从 F 楼层或比它低的楼层落下的鸡蛋都不会破。
//
// 每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）。
//
// 你的目标是确切地知道 F 的值是多少。
//
// 无论 F 的初始值如何，在最坏情况下，你确定 F 的值的最小移动次数是多少？
//
// 示例 1：
// 输入：K = 1, N = 2
// 输出：2
// 解释：
// 鸡蛋从 1 楼掉落。如果它碎了，我们肯定知道 F = 0 。
// 否则，鸡蛋从 2 楼掉落。如果它碎了，我们肯定知道 F = 1 。
// 如果它没碎，那么我们肯定知道 F = 2 。
// 因此，在最坏的情况下我们需要移动 2 次以确定 F 是多少。
//
// 示例 2：
// 输入：K = 2, N = 6
// 输出：3
//
// 示例 3：
// 输入：K = 3, N = 14
// 输出：4

// 1 <= K <= 100
// 1 <= N <= 10000

type FloorEgg struct {
	N int
	K int
}

func superEggDrop(K int, N int) int {
	mem := make(map[FloorEgg]int)
	fe := FloorEgg{
		N: N,
		K: K,
	}
	return eggDropWithDichotomy(fe, mem)
}

func eggDrop(fe FloorEgg, mem map[FloorEgg]int) int {
	// base case
	if fe.N == 0 {
		return 0
	}
	if fe.K == 1 {
		return fe.N
	}

	if res, ok := mem[fe]; ok {
		return res
	}

	res := math.MaxInt32
	for i := 1; i <= fe.N; i++ {
		broken := eggDrop(FloorEgg{
			N: i - 1,
			K: fe.K - 1,
		}, mem)
		notBroken := eggDrop(FloorEgg{
			N: fe.N - i,
			K: fe.K,
		}, mem)
		res = min(res, max(broken, notBroken)+1)
	}
	floorEgg := FloorEgg{
		N: fe.N,
		K: fe.K,
	}
	mem[floorEgg] = res
	return res
}

func eggDropWithDichotomy(fe FloorEgg, mem map[FloorEgg]int) int {
	if fe.N == 0 {
		return 0
	}
	if fe.K == 1 {
		return fe.N
	}

	if res, ok := mem[fe]; ok {
		return res
	}

	res := math.MaxInt32
	lo, hi := 1, fe.N
	for lo <= hi {
		mi := lo + (hi-lo)>>1

		// broken
		broken := eggDropWithDichotomy(FloorEgg{
			N: mi - 1,
			K: fe.K - 1,
		}, mem)
		// not broken
		notBroken := eggDropWithDichotomy(FloorEgg{
			N: fe.N - mi,
			K: fe.K,
		}, mem)

		// res = min(max(broken，not_broken) + 1)
		if broken > notBroken {
			res = min(res, broken+1)
			hi = mi - 1
		} else {
			res = min(res, notBroken+1)
			lo = mi + 1
		}
	}

	mem[fe] = res
	return res
}
