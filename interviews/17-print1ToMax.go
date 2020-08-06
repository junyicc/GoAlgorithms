package interviews

import (
	"fmt"
	"math"
	"strings"
)

// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
//
// 示例 1:
// 输入: n = 1
// 输出: [1,2,3,4,5,6,7,8,9]

func printNumbers(n int) []int {
	if n < 1 {
		return nil
	}
	cnt := int(math.Pow10(n))
	res := make([]int, cnt-1)
	for i := 1; i < cnt; i++ {
		res[i-1] = i
	}
	return res
}

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
			nTakeover = 0
			break // stop continuing increasing
		}
	}
	return isOverflow
}

// IToA converts number array to string
func IToA(number []byte) string {
	if len(number) == 0 {
		return ""
	}

	var b strings.Builder
	n := len(number)
	isBeginning0 := true // ignore beginning 0
	for i := 0; i < n; i++ {
		if isBeginning0 && number[i] != '0' {
			isBeginning0 = false // print following 0
		}
		if !isBeginning0 {
			b.WriteByte(number[i])
		}
	}
	return b.String()
}

func printNumber(number []byte) {
	s := IToA(number)
	fmt.Printf("%s\t", s)
}
