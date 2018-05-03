package sort

// MaxLengthInsertSort of InsertSort, min length of quick sort
var MaxLengthInsertSort = 7

// QuickSort algorithm
func QuickSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	if len(arr) > MaxLengthInsertSort {
		lo, hi := 0, len(arr)-1
		quickSort(arr, lo, hi)
		return arr
	}

	return InsertSort(arr)

}

func quickSort(arr []int, lo, hi int) {
	var pivot int
	for lo < hi {
		pivot = partition(arr, lo, hi)
		quickSort(arr, lo, pivot-1)
		lo = pivot + 1
	}
}

func partition(arr []int, lo, hi int) int {
	pivot := getPivot(arr, lo, hi)
	leftMark := lo + 1
	rightMark := hi
	for {
		for leftMark <= rightMark && pivot <= arr[rightMark] {
			rightMark--
		}
		for leftMark <= rightMark && arr[leftMark] <= pivot {
			leftMark++
		}
		if rightMark < leftMark {
			break
		} else {
			arr[leftMark], arr[rightMark] = arr[rightMark], arr[leftMark]
		}
	}
	arr[lo], arr[rightMark] = arr[rightMark], arr[lo]
	return rightMark
}

// getPivot returns median of three
func getPivot(arr []int, lo, hi int) int {
	mi := (lo + hi) >> 1

	if arr[lo] > arr[hi] {
		arr[lo], arr[hi] = arr[hi], arr[lo]
	}
	if arr[mi] > arr[hi] {
		arr[mi], arr[hi] = arr[hi], arr[mi]
	}
	if arr[mi] > arr[lo] {
		arr[mi], arr[lo] = arr[lo], arr[mi]
	}
	return arr[lo]
}
