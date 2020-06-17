package leetcode

func majorityElementII(nums []int) []int {
	if len(nums) < 1 {
		return nil
	}
	res := make([]int, 0, 2)

	cand1, cand2 := nums[0], nums[0]
	c1, c2 := 0, 0
	for _, num := range nums {
		if num == cand1 {
			c1++
			continue
		}
		if num == cand2 {
			c2++
			continue
		}
		if c1 == 0 {
			cand1 = num
			c1++
			continue
		}
		if c2 == 0 {
			cand2 = num
			c2++
			continue
		}
		c1--
		c2--
	}
	// counting
	c1, c2 = 0, 0
	for _, num := range nums {
		if num == cand1 {
			c1++
			continue
		}
		if num == cand2 {
			c2++
			continue
		}
	}
	// check
	if c1 > len(nums)/3 {
		res = append(res, cand1)
	}
	if c2 > len(nums)/3 {
		res = append(res, cand2)
	}
	return res
}
