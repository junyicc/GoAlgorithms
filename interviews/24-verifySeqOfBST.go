package interviews

// VerifyPostSeqOfBST verify post-order sequence of BST
func VerifyPostSeqOfBST(arr []int, length int) bool {
	if arr == nil || length < 1 {
		return false
	}
	if length == 1 {
		return true
	}
	root := arr[length-1]
	var left int
	for ; left < length-1; left++ {
		if arr[left] > root {
			break
		}
	}

	right := left
	for ; right < length-1; right++ {
		if arr[right] < root {
			return false
		}
	}
	// verify left subtree
	leftOrder := true
	if left > 0 {
		leftOrder = VerifyPostSeqOfBST(arr[:left], left)
	}
	// verify right subtree
	rightOrder := true
	if left < length-1 {
		rightOrder = VerifyPostSeqOfBST(arr[left:length-1], length-1-left)
	}

	return leftOrder && rightOrder

}

// VerifyPreSeqOfBST verify pre-order sequence of BST
func VerifyPreSeqOfBST(arr []int, length int) bool {
	if arr == nil || length < 1 {
		return false
	}
	if length == 1 {
		return true
	}

	root := arr[0]
	left := 1
	for ; left < length; left++ {
		if arr[left] > root {
			break
		}
	}
	right := left
	for ; right < length; right++ {
		if arr[right] < root {
			return false
		}
	}

	// verify left subtree
	leftOrder := true
	if left > 1 {
		leftOrder = VerifyPreSeqOfBST(arr[1:left], left-1)
	}

	// verify right subtree
	rightOrder := true
	if left < length {
		rightOrder = VerifyPreSeqOfBST(arr[left:], length-left)
	}

	return leftOrder && rightOrder
}
