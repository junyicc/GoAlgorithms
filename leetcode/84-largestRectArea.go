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
	stackFromBack := make([]int, 0, n)
	nextFromBack := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		// pop
		for len(stackFromBack) > 0 && heights[stackFromBack[len(stackFromBack)-1]] >= heights[i] {
			stackFromBack = stackFromBack[:len(stackFromBack)-1]
		}

		if len(stackFromBack) == 0 {
			nextFromBack[i] = n
		} else {
			nextFromBack[i] = stackFromBack[len(stackFromBack)-1]
		}

		stackFromBack = append(stackFromBack, i)
	}

	stackFromFront := make([]int, 0, n)
	nextFromFront := make([]int, n)
	for i := 0; i < n; i++ {
		for len(stackFromFront) > 0 && heights[stackFromFront[len(stackFromFront)-1]] >= heights[i] {
			stackFromFront = stackFromFront[:len(stackFromFront)-1]
		}

		if len(stackFromFront) == 0 {
			nextFromFront[i] = -1
		} else {
			nextFromFront[i] = stackFromFront[len(stackFromFront)-1]
		}
		stackFromFront = append(stackFromFront, i)
	}

	maxArea := -1
	for i := 0; i < n; i++ {
		area := heights[i] * (nextFromBack[i] - nextFromFront[i] - 1)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
