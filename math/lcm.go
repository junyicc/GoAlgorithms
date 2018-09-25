package math

// LCM returns the lcm of a, b
func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}
