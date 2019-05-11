package datastructure

import "fmt"

// FibRecur returns fib number using recursive method
func FibRecur(n int) int {
	if n < 2 {
		return n
	}
	return FibRecur(n-1) + FibRecur(n-2)
}

// Fib struct
type Fib struct{}

// Generate Fibnacci sequence number each time
func (f Fib) Generate() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

// Get fibonacci number
func (f Fib) Get(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// GetSeq returns fibonacci sequence
func (f Fib) GetSeq(n int) []int {
	fib := f.Generate()
	seq := []int{0}
	for i := 0; i < n; i++ {
		seq = append(seq, fib())
	}
	return seq
}

// TestFib tests Fib algorithm
func TestFib() {
	n := 5
	// fmt.Printf("fibonacci number %d: %d\n", n, GetFibNum(n))
	// fmt.Printf("fibonacci sequence %d: %d\n", n, GetFibSeq(n))
	f := Fib{}
	f1 := f.Generate()
	for i := 0; i < n; i++ {
		fmt.Println(f1())
	}
	fmt.Println()
	f2 := f.Generate()
	for i := 0; i < n; i++ {
		fmt.Println(f2())
	}

}
