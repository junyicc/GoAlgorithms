package interviews

import (
	"strings"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// ReverseStr returns reversed string
func ReverseStr(s string) string {
	if s == "" || len(s) <= 1 {
		return s
	}
	res := []byte(s)
	lo, hi := 0, len(s)-1
	for lo < hi {
		res[lo], res[hi] = res[hi], res[lo]
		lo++
		hi--
	}
	return string(res)
}

// ReverseStrByBit returns reversed string by bit
func ReverseStrByBit(s string) string {
	if s == "" || len(s) <= 1 {
		return s
	}
	res := []byte(s)
	lo, hi := 0, len(s)-1
	for lo < hi {
		res[lo] = res[lo] ^ res[hi]
		res[hi] = res[lo] ^ res[hi]
		res[lo] = res[lo] ^ res[hi]
		lo++
		hi--
	}
	return string(res)
}

// ReverseStrByStack returns reversed string using stack
func ReverseStrByStack(s string) string {
	if s == "" || len(s) <= 1 {
		return s
	}
	stack := datastructure.DynamicStack{}
	buf := strings.Builder{}
	for _, c := range []byte(s) {
		stack.Push(c)
	}
	for !stack.IsEmpty() {
		c := (stack.Pop()).(byte)
		buf.WriteByte(c)
	}
	return buf.String()
}
