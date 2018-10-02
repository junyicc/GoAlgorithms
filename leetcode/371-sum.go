package leetcode

func getSum(a int, b int) int {
	sum := a
	carrier := 0
	for b != 0 {
		sum = a ^ b
		carrier = (a & b) << 1
		a = sum
		b = carrier
	}
	return sum
}
