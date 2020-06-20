package leetcode

// Newton's method (https://en.wikipedia.org/wiki/Integer_square_root#Using_only_integer_division)

func mySqrtWithNewton(x int) int {
	if x <= 0 {
		return 0
	}
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}

// sqrt(x) <= x/2
func mySqrtWithBinary(x int) int {
	if x <= 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	lo, hi := 0, x>>1
	var ans int
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if mi*mi <= x {
			ans = mi
			lo = mi + 1
		} else {
			hi = mi - 1
		}
	}
	return ans
}
