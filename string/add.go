package string

func add(a, b string) string {
	if a == "" {
		return b
	}

	if b == "" {
		return a
	}

	n, m := len(a), len(b)
	var res []byte
	if n > m {
		res = make([]byte, n+1)
	} else {
		res = make([]byte, m+1)
	}

	var carry int
	k := len(res) - 1
	for i, j := n-1, m-1; i >= 0 || j >= 0; k-- {
		var n1, n2 int
		if i < 0 {
			n1 = 0
		} else {
			n1 = int(a[i] - '0')
			i--
		}

		if j < 0 {
			n2 = 0
		} else {
			n2 = int(b[j] - '0')
			j--
		}

		tmp := n1 + n2 + carry
		carry = tmp / 10
		res[k] = '0' + byte(tmp%10)
	}

	if carry > 0 {
		res[k] = '1'
	} else {
		res = res[1:]
	}

	return string(res)
}
