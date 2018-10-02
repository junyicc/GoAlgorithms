package leetcode

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	if n < 1 {
		return []string{}
	}
	res := []string{}
	for i := 1; i <= n; i++ {
		isFizz := i%3 == 0
		isBuzz := i%5 == 0
		if isFizz && isBuzz {
			res = append(res, "FizzBuzz")
		} else if isFizz {
			res = append(res, "Fizz")
		} else if isBuzz {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}
