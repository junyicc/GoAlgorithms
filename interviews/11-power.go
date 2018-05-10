package interviews

// Power returns base^exponent
func Power(base float64, exponent int) (float64, bool) {
	// invalid input
	if Equal(base, 0) && exponent < 0 {
		return 0, false
	}
	if Equal(base, 0) {
		return 0, true
	}

	absExponent := uint(exponent)
	if exponent < 0 {
		absExponent = (uint)(-exponent)
	}

	result := power(base, absExponent)
	if exponent < 0 {
		result = 1.0 / result
	}

	return result, true
}

func power(base float64, exponent uint) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}

	result := power(base, exponent>>1)
	result *= result
	if IsOdd(exponent) {
		result *= base
	}

	return result
}

// Equal returns true if two float number are equal
func Equal(f1, f2 float64) bool {
	if f1-f2 > -0.0000001 && f1-f2 < 0.0000001 {
		return true
	}
	return false
}
