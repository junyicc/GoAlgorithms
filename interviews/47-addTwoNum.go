package interviews

// AddTwoNum without using + - * /
func AddTwoNum(n1, n2 int) int {
	var sum, carry int
	for n2 != 0 {
		sum = n1 ^ n2
		carry = (n1 & n2) << 1
		n1 = sum
		n2 = carry
	}
	return n1
}
