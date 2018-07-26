package math

// StrIncrement self-increment
func StrIncrement(nums []byte) bool {
	isOverflow := false
	carry := 0
	n := len(nums)
	for i := n - 1; i >= 0; i-- {
		sum := int(nums[i]-'0') + carry
		if i == n-1 {
			sum++
		}
		if sum >= 10 {
			if i == 0 {
				isOverflow = true
			} else {
				carry = 1
				sum -= 10
				nums[i] = byte(sum) + '0'
			}
		} else {
			carry = 0
			nums[i] = byte(sum) + '0'
			break
		}
	}
	return isOverflow
}

// BinaryAdd adds nums without using + - * /
func BinaryAdd(a, b int) int {
	sum := a
	carry := b
	for carry != 0 {
		sum = a ^ b
		carry = (a & b) << 1
		a = sum
		b = carry
	}
	return sum
}
