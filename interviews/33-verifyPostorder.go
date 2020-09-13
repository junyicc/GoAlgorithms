package interviews

func verifyPostorder(postorder []int) bool {
	if len(postorder) < 2 {
		return true
	}

	root := postorder[len(postorder)-1]
	var left int
	for left < len(postorder)-1 && postorder[left] <= root {
		left++
	}

	// leftOrder and rightOrder is true, when leftSubtree or rightSubtree does not exist
	leftOrder, rightOrder := true, true

	// check valid right subtree
	right := left
	for ; right < len(postorder)-1; right++ {
		if postorder[right] <= root {
			return false
		}
	}

	if left > 0 {
		leftOrder = verifyPostorder(postorder[:left])
	}

	if right > left {
		rightOrder = verifyPostorder(postorder[left : len(postorder)-1])
	}

	return leftOrder && rightOrder
}
