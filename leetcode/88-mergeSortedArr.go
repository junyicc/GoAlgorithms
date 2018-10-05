package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	if nums1 == nil || len(nums1) < 1 || nums2 == nil || n < 1 {
		return
	}

	last := len(nums1) - 1
	for i, j := m-1, n-1; j >= 0; {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[last] = nums1[i]
			i--
		} else {
			nums1[last] = nums2[j]
			j--
		}
		last--
	}
}
