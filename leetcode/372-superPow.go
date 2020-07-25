package leetcode

// Your task is to calculate a^b mod 1337 where a is a positive integer and b is an extremely large positive integer given in the form of an array.
//
// Example 1:
// Input: a = 2, b = [3]
// Output: 8
//
// Example 2:
// Input: a = 2, b = [1,0]
// Output: 1024

func superPow(a int, b []int) int {
	if len(b) < 1 {
		return 1
	}

	base := 1337

	// (a*b)%k = (a%k)(b%k)%k
	mypow := func(a, k int) int {
		// k is between [0, 10]
		res := 1

		// may overflow
		// for i := 0; i < k; i++ {
		// 	tmp := a % base
		// 	res *= tmp
		// }

		tmp := a % base
		for i := 0; i < k; i++ {
			res *= tmp
			res %= base
		}

		return res
	}

	// a^1234 = a^4 * a^123^10
	last := b[len(b)-1]
	b = b[:len(b)-1]

	// part1 = ((a^last)%base) % base
	part1 := mypow(a, last)
	part2 := mypow(superPow(a, b), 10)

	return part1 * part2 % base
}

func superPow2(a int, b []int) int {
	if len(b) < 1 {
		return 1
	}

	base := 1337

	// a^1234 = a^4 * a^123^10
	last := b[len(b)-1]
	b = b[:len(b)-1]

	// part1 = ((a^last)%base) % base
	part1 := mypow(a, last, base)
	part2 := mypow(superPow(a, b), 10, base)

	return part1 * part2 % base
}

func mypow(a, k int, base int) int {
	if k == 0 {
		return 1
	}
	if k == 1 {
		return a % base
	}

	// (a*b)%k = (a%k)(b%k)%k
	res := mypow(a, k>>1, base)
	res *= res
	// k is odd
	if k&1 == 1 {
		res *= a % base
	}

	return res % base
}
