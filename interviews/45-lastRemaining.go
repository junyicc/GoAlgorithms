package interviews

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
