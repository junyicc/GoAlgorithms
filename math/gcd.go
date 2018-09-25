package math

// GCD returns the gcd of a, b
func GCD(a, b int) int {
	if a < b {
		return GCD(b, a)
	}
	if b == 0 {
		return a
	}

	aEven := IsEven(a)
	bEven := IsEven(b)

	if aEven && bEven {
		return 2 * GCD(a>>1, b>>1)
	} else if aEven && !bEven {
		return GCD(a>>1, b)
	} else if !aEven && bEven {
		return GCD(a, b>>1)
	} else {
		return GCD(b, a-b)
	}
}

// GCD1 returns the gcd of a, b
func GCD1(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
