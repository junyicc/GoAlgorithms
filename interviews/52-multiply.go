package interviews

// MultiplyArray returns an array
// where array[i] = arr[0]*arr[1]*...*arr[i-1]*arr[i+1]*...*arr[n-1]
func MultiplyArray(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return nil
	}
	length := len(arr)
	targetArr := make([]int, length)
	for i := 0; i < length; i++ {
		targetArr[i] = 1
	}

	tmp := 1
	for i := 1; i < length; i++ {
		tmp *= arr[i-1]
		targetArr[i] *= tmp
	}
	tmp = 1
	for i := length - 2; i >= 0; i-- {
		tmp *= arr[i+1]
		targetArr[i] *= tmp
	}
	return targetArr
}
