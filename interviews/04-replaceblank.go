package interviews

// ReplaceBlank blank with %20
func ReplaceBlank(str string) string {
	if str == "" || len(str) < 1 {
		return ""
	}
	strArr := []rune(str)
	oLen := len(str)
	numOfOld := 0
	for _, c := range str {
		if c == ' ' {
			numOfOld++
		}
	}
	dLen := oLen + 2*numOfOld
	for i := 0; i < 2*numOfOld; i++ {
		strArr = append(strArr, ' ')
	}

	for i, j := oLen-1, dLen-1; i >= 0 && j > i; i-- {
		if strArr[i] != ' ' {
			strArr[j] = strArr[i]
			j--
		} else {
			strArr[j] = '0'
			strArr[j-1] = '2'
			strArr[j-2] = '%'
			j -= 3
		}
	}
	return string(strArr)
}
