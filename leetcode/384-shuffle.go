package leetcode

import "math/rand"

// Solution for shuffle
type Solution struct {
	nums    []int
	shuffle []int
}

// ShuffleConstructor Constructor
func ShuffleConstructor(nums []int) Solution {
	return Solution{
		nums:    nums,
		shuffle: append([]int{}, nums...),
	}
}

// Reset Resets the array to its original configuration and return it.
func (s *Solution) Reset() []int {
	return s.nums
}

// Shuffle returns a random shuffling of the array.
func (s *Solution) Shuffle() []int {
	n := len(s.nums)
	rand.Shuffle(n, func(i, j int) {
		s.shuffle[i], s.shuffle[j] = s.shuffle[j], s.shuffle[i]
	})
	// simple shuffle
	// for i := 0; i < n; i++ {
	// 	j := rand.Intn(n)
	// 	s.shuffle[i], s.shuffle[j] = s.shuffle[j], s.shuffle[i]
	// }
	return s.shuffle
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
