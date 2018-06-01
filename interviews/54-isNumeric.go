package interviews

// IsNumeric returns true if str is a numeric string
// that matches the pattern [sign]integral-digits[.[fractional-digits]][e|E[sign]exponential-digits]
func IsNumeric(str string) bool {
	if str == "" {
		return false
	}

	strArr := []rune(str)

	i := 0
	if strArr[i] == '+' || strArr[i] == '-' {
		i++
	}
	if i == len(strArr) {
		return false
	}
	numeric := true
	scanDigits(strArr, &i)
	if i < len(strArr) {
		// for float
		if strArr[i] == '.' {
			i++
			scanDigits(strArr, &i)
			if i < len(strArr) && (strArr[i] == 'e' || strArr[i] == 'E') {
				i++
				if i < len(strArr) && strArr[i] == '+' || strArr[i] == '-' {
					i++
				}
				scanDigits(strArr, &i)
			}
		} else if strArr[i] == 'e' || strArr[i] == 'E' {
			i++
			if i < len(strArr) && strArr[i] == '+' || strArr[i] == '-' {
				i++
			}
			scanDigits(strArr, &i)
		} else {
			numeric = false
		}
	}
	numeric = numeric && i == len(strArr)
	return numeric
}

func scanDigits(strArr []rune, i *int) {
	for *i < len(strArr) && strArr[*i] >= '0' && strArr[*i] <= '9' {
		*i++
	}
}
