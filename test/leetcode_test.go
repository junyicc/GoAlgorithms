package test

import (
	"fmt"
	"testing"

	"github.com/CAIJUNYI/GoAlgorithms/leetcode"
)

func TestThreeSum(t *testing.T) {
	result := leetcode.ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(result)

	result = leetcode.ThreeSum([]int{0, 0, 0})
	fmt.Println(result)
}

func TestThreeSumClosest(t *testing.T) {
	if result := leetcode.ThreeSumClosest([]int{-1, 2, 1, 4}, 1); result != 2 {
		t.Errorf("expected 2 and got %d", result)
	}
}

func TestFourSum(t *testing.T) {
	result := leetcode.FourSum([]int{0, 0, 0, 0}, 0)
	fmt.Println(result)
}
