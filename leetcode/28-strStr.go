package leetcode

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	n, m := len(haystack), len(needle)
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == m {
				return i
			}
			if i+j == n {
				return -1
			}
			if haystack[i+j] != needle[j] {
				break
			}
		}
	}
}
