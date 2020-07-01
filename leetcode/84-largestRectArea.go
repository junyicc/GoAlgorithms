package leetcode

func largestRectangleAreaSimple(heights []int) int {
	if len(heights) < 1 {
		return 0
	}
	n := len(heights)
	var maxRectArea int
	for i := 0; i < n; i++ {
		h := heights[i]

		// check left side
		leftCnt := 0
		for left := i - 1; left >= 0 && heights[left] >= h; left-- {
			leftCnt++
		}
		// check right side
		rightCnt := 0
		for right := i + 1; right < n && heights[right] >= h; right++ {
			rightCnt++
		}
		area := (leftCnt + rightCnt + 1) * h
		if area > maxRectArea {
			maxRectArea = area
		}
	}

	return maxRectArea
}
