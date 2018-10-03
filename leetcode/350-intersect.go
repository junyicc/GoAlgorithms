package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	if nums1 == nil || len(nums1) < 1 || nums2 == nil || len(nums2) < 1 {
		return nil
	}
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		return intersect(nums2, nums1)
	}
	numCnt := make(map[int]int)
	for _, num := range nums2 {
		numCnt[num]++
	}
	res := []int{}
	for _, num := range nums1 {
		if c, ok := numCnt[num]; ok && c > 0 {
			res = append(res, num)
			numCnt[num]--
		}
	}
	return res
}
