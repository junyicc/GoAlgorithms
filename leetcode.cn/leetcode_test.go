package leetcodecn

import "testing"

func TestIsPalindromeByReverse(t *testing.T) {
	head := makeList([]int{1, 2, 2, 1})
	isPalindromeByReverse(head)
}
