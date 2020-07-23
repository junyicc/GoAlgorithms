package leetcode

import (
	"math/rand"
	"time"
)

// Given a singly linked list, return a random node's value from the linked list. Each node must have the same probability of being chosen.
//
// Follow up:
// What if the linked list is extremely large and its length is unknown to you? Could you solve this efficiently without using extra space?
//
// Example:
//
// // Init a singly linked list [1,2,3].
// ListNode head = new ListNode(1);
// head.next = new ListNode(2);
// head.next.next = new ListNode(3);
// Solution solution = new Solution(head);
//
// // getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
// solution.getRandom();

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type RandomSolution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func RandomConstructor(head *ListNode) *RandomSolution {
	rand.Seed(time.Now().Unix())
	return &RandomSolution{head: head}
}

/** Returns a random node's value. */
func (s *RandomSolution) GetRandom() int {
	var res int
	p := s.head
	i := 0
	for p != nil {
		i++
		// [0, i)
		if rand.Intn(i) == 0 {
			res = p.Val
		}
		p = p.Next
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
