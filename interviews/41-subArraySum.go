package interviews

// TwoNumWithSum returns two numbers whose sum is equal to s
func TwoNumWithSum(arr []int, s int) (int, int, bool) {
	if arr == nil || len(arr) < 2 {
		return 0, 0, false
	}
	found := false
	lo, hi := 0, len(arr)-1
	for lo < hi {
		if arr[lo]+arr[hi] == s {
			found = true
			break
		} else if arr[lo]+arr[hi] < s {
			lo++
		} else {
			hi--
		}
	}
	return arr[lo], arr[hi], found
}

// ContinuousSeqWithSum visits continuous positive slices starting at lo and ending at hi, whose sum are equal to s
func ContinuousSeqWithSum(s int, f func(int, int)) bool {
	if s < 1 || s == 2 {
		return false
	}
	if s == 1 {
		f(1, 1)
		return true
	}

	found := false
	lo, hi := 1, 2
	curSum := lo + hi
	mi := (s + 1) >> 1
	for lo < mi {
		if curSum == s {
			f(lo, hi)
			found = true
		}
		for lo < mi && curSum > s {
			curSum -= lo
			lo++
			if curSum == s {
				f(lo, hi)
				found = true
			}
		}
		// curSum < s
		hi++
		curSum += hi
	}
	return found
}
