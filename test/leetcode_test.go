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

func TestProducerConsumer(t *testing.T) {
	data := make(chan int, 5)
	leetcode.ProducerWG.Add(1)
	go leetcode.Producer("P1", data)
	leetcode.ConsumerWG.Add(1)
	go leetcode.Consumer("U1", data)
	leetcode.ProducerWG.Add(1)
	go leetcode.Producer("P2", data)
	leetcode.ConsumerWG.Add(1)
	go leetcode.Consumer("U2", data)

	leetcode.ProducerWG.Wait()
	close(data)
	leetcode.ConsumerWG.Wait()
}
