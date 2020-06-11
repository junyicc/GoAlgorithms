package math

import (
	"fmt"
	"strings"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

var digits = [16]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F'}

// ConvertBase converts num into the form in base
func ConvertBase(num, base int) (string, error) {
	if base > 16 || base < 2 {
		return "", fmt.Errorf("ConvertBase: base must be ge 2 and le 16")
	}
	stack := datastructure.DynamicStack{}
	for num > 0 {
		reaminder := num % base
		stack.Push(digits[reaminder])
		num /= base
	}
	var buf strings.Builder
	for !stack.IsEmpty() {
		digit := (stack.Pop()).(byte)
		buf.WriteByte(digit)
	}
	return buf.String(), nil
}
