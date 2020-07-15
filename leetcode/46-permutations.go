package leetcode

func permute(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	res := [][]int{}
	f := func(nums []int) {
		res = append(res, nums)
	}
	arrPermutations(nums, 0, len(nums), f)
	return res
}

func arrPermutations(nums []int, start, n int, f func([]int)) {
	if start == n-1 {
		pNums := make([]int, n)
		copy(pNums, nums)
		f(pNums)
		return
	}

	for i := start; i < n; i++ {
		nums[start], nums[i] = nums[i], nums[start]
		arrPermutations(nums, start+1, n, f)
		nums[start], nums[i] = nums[i], nums[start]
	}
}

func permuteWithBacktracking(nums []int) [][]int {
	if len(nums) < 1 {
		return nil
	}
	res := [][]int{}
	f := func(nums []int) {
		res = append(res, nums)
	}
	track := make([]int, 0, len(nums))
	backtrackPermute(nums, track, f)

	return res
}

func backtrackPermute(nums []int, track []int, f func([]int)) {
	if len(track) == len(nums) {
		clone := append(track[:0:0], track...)
		f(clone)
		return
	}

	for i := 0; i < len(nums); i++ {
		if inTrack(track, nums[i]) {
			continue
		}
		// make choice
		track = append(track, nums[i])

		backtrackPermute(nums, track, f)

		track = track[:len(track)-1]
	}
}

func inTrack(track []int, num int) bool {
	for _, n := range track {
		if n == num {
			return true
		}
	}
	return false
}
