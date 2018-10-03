package leetcode

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	slow, fast := n, n
	for {
		slow = nextNum(slow)
		fast = nextNum(fast)
		fast = nextNum(fast)
		println(slow)
		println(fast)
		println("---")
		if slow == fast {
			break
		}
	}
	if slow == 1 {
		return true
	}
	return false
}

func nextNum(n int) int {
	newNum := 0
	for n != 0 {
		digit := n % 10
		newNum += digit * digit
		n /= 10
	}
	return newNum
}
