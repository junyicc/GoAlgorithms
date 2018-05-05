package search

import (
	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// FibSearch uses Fibonacci sequence to search element e in the array
// n = len(left_part) + mi + len(right_part)
// n = Fib(k)-1
// Fib(k)-1 = Fib(k-1)-1 + 1 + Fib(k-2) -1
func FibSearch(arr []int, e int) (int, bool) {
	if arr == nil || len(arr) < 1 {
		return -1, false
	}

	lo, hi := 0, len(arr)
	var k int
	var f datastructure.Fib
	for lo < hi {
		var fib = f.New()
		for hi-lo < fib() {
			k++
		}
		mi := lo + f.Get(k-1) - 1
		if e < arr[mi] {
			hi = mi
		} else if arr[mi] < e {
			lo = mi + 1
		} else {
			return mi, true
		}
	}
	return -1, false
}
