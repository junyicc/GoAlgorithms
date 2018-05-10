package interviews

import (
	"bytes"
	"fmt"
)

// Print1ToMaxOfNDigits prints 1 to max of n digits
func Print1ToMaxOfNDigits(n int) {
	if n <= 0 {
		return
	}

	number := make([]byte, n)
	for i := 0; i < n; i++ {
		number[i] = '0'
	}

	for !increment(number) {
		printNumber(number)
	}
}

func increment(number []byte) bool {
	isOverflow := false
	var nTakeover byte
	n := len(number)
	for i := n - 1; i >= 0; i-- {
		nsum := number[i] - '0' + nTakeover
		// add one
		if i == n-1 {
			nsum++
		}

		if nsum > 9 {
			if i == 0 {
				isOverflow = true
			} else {
				nsum -= 10
				nTakeover = 1
				number[i] = nsum + '0'
			}
		} else {
			number[i] = nsum + '0'
			break // stop continuing increasing
		}
	}
	return isOverflow
}

// IToA converts number array to string
func IToA(number []byte) string {
	if number == nil || len(number) == 0 {
		return ""
	}

	var b bytes.Buffer
	n := len(number)
	isBeginning0 := true // ignore beginning 0
	for i := 0; i < n; i++ {
		if isBeginning0 && number[i] != '0' {
			isBeginning0 = false // print following 0
		}
		if !isBeginning0 {
			b.WriteString(fmt.Sprintf("%c", number[i]))
		}
	}
	return b.String()
}

func printNumber(number []byte) {
	s := IToA(number)
	fmt.Printf("%s\t", s)
}
