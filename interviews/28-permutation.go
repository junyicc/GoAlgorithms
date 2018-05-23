package interviews

// Permutation string
func Permutation(str *string, f func(string)) {
	if str == nil || *str == "" {
		return
	}
	permutation(([]byte)(*str), 0, f)
}

func permutation(str []byte, begin int, f func(string)) {
	if begin == len(str) {
		f(string(str))
	} else {
		for i := begin; i < len(str); i++ {
			// swap
			str[begin], str[i] = str[i], str[begin]
			permutation(str, begin+1, f)
			// reset
			str[begin], str[i] = str[i], str[begin]
		}
	}
}
