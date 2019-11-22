package interviews

// ReorderOddEven reorders the array as odd part + even part
func ReorderOddEven(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	lo, hi := 0, len(arr)-1
	for {
		for lo < hi && IsOdd(arr[lo]) {
			lo++
		}
		for lo < hi && !IsOdd(arr[hi]) {
			hi--
		}
		if lo < hi {
			arr[lo], arr[hi] = arr[hi], arr[lo]
		} else {
			break
		}
	}
	return arr
}

// ReorderOddEven2 extensible version
func ReorderOddEven2(arr []int) []int {
	return Reorder(arr, IsOdd)
}

// Reorder array according to function f
func Reorder(arr []int, f func(int) bool) []int {
	if len(arr) == 0 {
		return nil
	}

	lo, hi := 0, len(arr)-1
	for {
		for lo < hi && f(arr[lo]) {
			lo++
		}
		for lo < hi && !f(arr[hi]) {
			hi--
		}
		if lo < hi {
			arr[lo], arr[hi] = arr[hi], arr[lo]
		} else {
			break
		}
	}
	return arr
}
