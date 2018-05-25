package interviews

// GetUglyNum returns the nth ugly number
// whose prime factors only include 2, 3, 5, and the first ugly number is 1
func GetUglyNum(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	if n == 1 {
		return 1, true
	}
	uglyNums := []int{1}
	ugly2, ugly3, ugly5 := 0, 0, 0

	for next := 1; next < n; next++ {
		min := minOfThree(uglyNums[ugly2]*2, uglyNums[ugly3]*3, uglyNums[ugly5]*5)
		uglyNums = append(uglyNums, min)

		for uglyNums[ugly2]*2 <= uglyNums[next] {
			ugly2++
		}
		for uglyNums[ugly3]*3 <= uglyNums[next] {
			ugly3++
		}
		for uglyNums[ugly5]*5 <= uglyNums[next] {
			ugly5++
		}
	}
	return uglyNums[n-1], true
}

func minOfThree(n1, n2, n3 int) int {
	min := n1
	if n2 < n1 {
		min = n2
	}
	if n3 < min {
		min = n3
	}
	return min
}

// IsUglyNum returns true if n's prime factors only include 2, 3, 5
func IsUglyNum(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	for n%2 == 0 {
		n /= 2
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%5 == 0 {
		n /= 5
	}

	if n == 1 {
		return true
	}
	return false
}

// GetSuperUglyNum returns the nth ugly number
// whose prime factors only include primes, and the first ugly number is 1
func GetSuperUglyNum(n int, primes []int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	if n == 1 {
		return 1, true
	}
	length := len(primes)
	uglyNums := []int{1}
	uglyFactors := make([]int, length)
	for next := 1; next < n; next++ {
		tmpUgly := make([]int, length)
		for i := 0; i < length; i++ {
			tmpUgly[i] = uglyNums[uglyFactors[i]] * primes[i]
		}
		min := minOfSlice(tmpUgly)
		uglyNums = append(uglyNums, min)
		for j := 0; j < length; j++ {
			if tmpUgly[j] <= uglyNums[next] {
				uglyFactors[j]++
			}
		}
	}
	return uglyNums[n-1], true
}

func minOfSlice(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
