package string

func sortString(s string) string {
	if len(s) < 2 {
		return s
	}
	bs := []byte(s)
	sortByteSlice(bs)
	return string(bs)
}

func sortByteSlice(bs []byte) {
	asc := [256]int{}
	// counting
	for _, c := range bs {
		asc[c]++
	}
	idx := 0
	for i := 0; i < 256; i++ {
		if asc[i] > 0 {
			cnt := asc[i]
			for j := 0; j < cnt; j++ {
				bs[idx] = byte(i)
				idx++
			}
		}
	}
}
