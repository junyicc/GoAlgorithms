package math

// Permutation method
func Permutation(n, m int) int {
	if n < 0 || m < 0 {
		panic("Permutation: invalid n or m")
	}
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	if m == 0 {
		return 1
	} else if m == 1 {
		return n
	}

	return Factorial(n) / Factorial(n-m)
}

// StringPermutation permutates each char in string
func StringPermutation(str string) []string {
	if str == "" {
		return nil
	}
	result := []string{}
	f := func(s string) {
		result = append(result, s)
	}
	stringPermutation([]byte(str), 0, f)
	return result
}

func stringPermutation(str []byte, begin int, f func(string)) {
	n := len(str)
	if begin == n {
		f(string(str))
		return
	}

	for i := begin; i < n; i++ {
		str[begin], str[i] = str[i], str[begin]
		stringPermutation(str, begin+1, f)
		str[begin], str[i] = str[i], str[begin]
	}
}
