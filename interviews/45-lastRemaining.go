package interviews

import "container/ring"

// LastRemaining returns the last remaining number in the loop of 0, 1, ..., n-1 when deleting the mth number each time
// f(n, m) =  0,			  n=1
//         =  [f(n-1,m)+m]%n, n>1
func LastRemaining(n, m int) int {
	if n < 1 || m < 1 {
		return -1
	}
	var last int
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}

// LastRemaining2 simulates deleting the mth number each time from the loop of 0, 1, ..., n-1, and returns the last remaining number
func LastRemaining2(n, m int) (int, bool) {
	if n < 1 || m < 1 || n < m {
		return 0, false
	}

	numRing := ring.New(n)
	for i := 0; i < n; i++ {
		numRing.Value = i
		numRing = numRing.Next()
	}

	for numRing.Len() > 1 {
		for i := 1; i < m-1; i++ {
			numRing = numRing.Next()
		}

		numRing.Unlink(1)
		numRing = numRing.Next()
	}
	return numRing.Value.(int), true
}
