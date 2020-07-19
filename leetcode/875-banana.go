package leetcode

// Koko loves to eat bananas.  There are N piles of bananas, the i-th pile has piles[i] bananas.  The guards have gone and will come back in H hours.
//
// Koko can decide her bananas-per-hour eating speed of K.  Each hour, she chooses some pile of bananas, and eats K bananas from that pile.  If the pile has less than K bananas, she eats all of them instead, and won't eat any more bananas during this hour.
//
// Koko likes to eat slowly, but still wants to finish eating all the bananas before the guards come back.
//
// Return the minimum integer K such that she can eat all the bananas within H hours.
//
// Example 1:
// Input: piles = [3,6,7,11], H = 8
// Output: 4
//
// Example 2:
// Input: piles = [30,11,23,4,20], H = 5
// Output: 30
//
// Example 3:
// Input: piles = [30,11,23,4,20], H = 6
// Output: 23

func minEatingSpeed(piles []int, H int) int {
	if len(piles) < 1 || H < 1 {
		return 0
	}

	lo, hi := 1, maxOfSlice(piles)
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if !canFinish(piles, H, mi) {
			lo = mi + 1
		} else {
			if mi == 1 || !canFinish(piles, H, mi-1) {
				return mi
			} else {
				hi = mi - 1
			}
		}
	}
	return 0
}

func canFinish(piles []int, H int, speed int) bool {
	times := 0
	for _, p := range piles {
		times += timesOf(p, speed)
	}
	return times <= H
}

func timesOf(n int, speed int) int {
	if n%speed > 0 {
		return n/speed + 1
	}
	return n / speed
}

func maxOfSlice(nums []int) int {
	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	return maxVal
}
