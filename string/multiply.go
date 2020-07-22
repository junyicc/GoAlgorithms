package string

func multiply(a, b string) string {
	m, n := len(a), len(b)
	if m < 1 || n < 1 {
		return "0"
	}

	res := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(a[i]-'0') * int(b[j]-'0')
			// result should locate at i+j, i+j+1
			p1, p2 := i+j, i+j+1

			sum := mul + res[p2]
			res[p2] = sum % 10
			res[p1] += sum / 10
		}
	}

	// remove leading 0
	var i int
	for ; i < (m + n); i++ {
		if res[i] != 0 {
			break
		}
	}

	bs := make([]byte, 0, (m+n)-i)
	for ; i < (m + n); i++ {
		bs = append(bs, byte(res[i]+'0'))
	}
	if len(bs) > 1 {
		return string(bs)
	}
	return "0"
}
