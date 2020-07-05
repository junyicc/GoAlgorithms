package interviews

// CountingK returns the number of k in the sorted array
func CountingK(arr []int, k int) int {
	if len(arr) < 1 {
		return 0
	}

	firstIndex := getFirstK(arr, k)
	lastIndex := getLastK(arr, k)
	cnt := 0
	if firstIndex >= 0 && lastIndex >= 0 {
		cnt = lastIndex - firstIndex + 1
	}
	return cnt
}

func getFirstK(arr []int, k int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if arr[mi] < k {
			lo = mi + 1
		} else {
			if mi == 0 || arr[mi-1] < k {
				return mi
			} else {
				hi = mi - 1
			}
		}
	}
	return -1
}
func getLastK(arr []int, k int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if arr[mi] > k {
			hi = mi - 1
		} else {
			if mi == len(arr)-1 || arr[mi+1] > k {
				return mi
			} else {
				lo = mi + 1
			}
		}
	}
	return -1
}

// 一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
//
// 示例 1:
// 输入: [0,1,3]
// 输出: 2
//
// 示例 2:
// 输入: [0,1,2,3,4,5,6,7,9]
// 输出: 8
func missingNumber(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if nums[mi] == mi {
			lo = mi + 1
		} else {
			if mi == 0 || nums[mi-1] == mi-1 {
				return mi
			} else {
				hi = mi - 1
			}
		}
	}
	if lo > hi {
		return len(nums)
	}
	return -1
}
