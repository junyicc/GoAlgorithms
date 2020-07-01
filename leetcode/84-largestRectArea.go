package leetcode

// Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.

// using mono-stack to record thr leftmost element index that less than or equal to element i
// the leftmost boundary of element i is the former element index when pushed into mono-stack
// the rightmost boundary of element i is the current element index when popped out
// calculate area when pop element
func largestRectangleArea(heights []int) int {
	if len(heights) < 1 {
		return 0
	}
	// add sentinel to handle tail boundary
	heights = append(heights, 0)

	n := len(heights)
	monoStack := []int{}
	maxRectArea := 0
	for i := 0; i < n; i++ {
		h := heights[i]
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] >= h {
			popIndex := monoStack[len(monoStack)-1]
			monoStack = monoStack[:len(monoStack)-1]

			var w int
			if len(monoStack) == 0 {
				// last one is the min height: w=i-(-1)-1
				w = i
			} else {
				// leftmost boundary
				nearestMinIndex := monoStack[len(monoStack)-1]
				w = i - nearestMinIndex - 1
			}

			area := w * heights[popIndex]
			if area > maxRectArea {
				maxRectArea = area
			}
		}
		monoStack = append(monoStack, i)
	}
	return maxRectArea
}

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
