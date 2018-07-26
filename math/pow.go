package math

import (
	"fmt"
)

// Pow method
func Pow(base float64, exponent int) (float64, error) {
	if equal(base, 0.0) && exponent < 0 {
		return 0.0, fmt.Errorf("Pow: divided by 0")
	} else if equal(base, 0.0) && exponent >= 0 {
		return 0.0, nil
	}

	absExponent := exponent
	if exponent < 0 {
		absExponent = -exponent
	}
	result := pow(base, absExponent)
	if exponent < 0 {
		result = 1 / result
	}
	return result, nil
}

func pow(base float64, exponent int) float64 {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	result := pow(base, exponent>>1)
	result *= result
	if exponent&1 == 1 {
		result *= base
	}
	return result
}

func equal(a, b float64) bool {
	if a-b < 0.00000001 || b-a < 0.00000001 {
		return true
	}
	return false
}
