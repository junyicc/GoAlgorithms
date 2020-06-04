package sort

// MaxLengthInsertSort of InsertSort, min length of quick sort
var MaxLengthInsertSort = 7

// QuickSort algorithm
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	lo, hi := 0, len(arr)
	quickSort(arr, lo, hi)
}

func quickSort(arr []int, lo, hi int) {
	for hi-lo > MaxLengthInsertSort {
		pivot := partition(arr, lo, hi)
		quickSort(arr, lo, pivot)
		lo = pivot + 1
	}
	if hi-lo > 1 {
		insertionSort(arr, lo, hi)
	}
}

func partition(arr []int, lo, hi int) int {
	mi := lo + (hi-lo)>>1
	medianOfThree(arr, lo, mi, hi-1)

	pivot := lo
	leftMark := lo + 1
	rightMark := hi - 1

	for {
		for leftMark <= rightMark && arr[pivot] <= arr[rightMark] {
			rightMark--
		}
		for leftMark <= rightMark && arr[leftMark] <= arr[pivot] {
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

// medianOfThree sets lo with median value
func medianOfThree(arr []int, lo, mi, hi int) {
	if arr[lo] > arr[hi] {
		arr[lo], arr[hi] = arr[hi], arr[lo]
	}
	if arr[mi] > arr[hi] {
		arr[mi], arr[hi] = arr[hi], arr[mi]
	}
	if arr[mi] > arr[lo] {
		arr[mi], arr[lo] = arr[lo], arr[mi]
	}
}
