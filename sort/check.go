package sort

// Check whether the sequence is ascending
func Check(seq []interface{}, ascending bool) bool {
	if ascending {
		return isAscending(seq)
	}
	return isDescending(seq)
}

func isAscending(seq []interface{}) bool {
	inverse := false
	for i := 1; i < len(seq); i++ {
		pre, preOk := seq[i-1].(int)
		latter, latterOk := seq[i].(int)
		if preOk && latterOk {
			if pre > latter {
				inverse = true
				break
			}
		}
	}
	return !inverse
}

func isDescending(seq []interface{}) bool {
	inverse := false
	for i := 1; i < len(seq); i++ {
		pre, preOk := seq[i-1].(int)
		latter, latterOk := seq[i].(int)
		if preOk && latterOk {
			if pre < latter {
				inverse = true
				break
			}
		}
	}
	return !inverse
}
