package leetcode

// Newton's method (https://en.wikipedia.org/wiki/Integer_square_root#Using_only_integer_division)

func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
