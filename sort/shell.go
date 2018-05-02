package sort

// ShellSort algorithm
func ShellSort(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}
	var i, j, tmp int
	d := len(arr)
	for d > 1 {
		// step sequence
		d = d/3 + 1
		for i = d; i < len(arr); i++ {
			// insertsort
			if arr[i] < arr[i-d] {
				tmp = arr[i]
				for j = i - d; j >= 0 && tmp < arr[j]; j -= d {
					arr[j+d] = arr[j]
				}
				arr[j+d] = tmp
			}
		}
	}
	return arr
}
