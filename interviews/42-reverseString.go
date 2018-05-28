package interviews

import (
	"bytes"
	"strings"
	"unicode"
)

// ReverseSentence reverses word sequence while keeps characters constant
func ReverseSentence(s string) string {
	if s == "" || len(s) < 2 {
		return s
	}
	// reverse whole sentence
	words := []rune(s)
	lo, hi := 0, len(words)-1
	reverseString(words, lo, hi)
	// reverse each word
	lo, hi = 0, 0
	for lo < len(words) {
		if unicode.IsSpace(words[lo]) {
			lo++
			hi++
		} else if hi == len(words) || unicode.IsSpace(words[hi]) {
			hi--
			reverseString(words, lo, hi)
			hi++
			lo = hi
		} else {
			hi++
		}
	}
	return string(words)
}

func reverseString(s []rune, lo, hi int) {
	if s == nil || len(s) < 2 {
		return
	}
	for lo < hi {
		s[lo], s[hi] = s[hi], s[lo]
		lo++
		hi--
	}
}

// ReverseSentence2 reverses word sequence while keeps characters constant
func ReverseSentence2(s string) string {
	if s == "" || len(s) < 2 {
		return s
	}
	reversedStr := reverseString2(s)
	reversedWords := strings.Fields(reversedStr)
	var b bytes.Buffer
	for i, w := range reversedWords {
		normalWord := reverseString2(w)
		b.WriteString(normalWord)
		if i < len(reversedWords)-1 {
			b.WriteString(" ")
		}
	}
	return b.String()
}

func reverseString2(s string) string {
	charArr := []rune(s)
	lo, hi := 0, len(charArr)-1
	for lo <= hi {
		charArr[lo], charArr[hi] = charArr[hi], charArr[lo]
		lo++
		hi--
	}
	return string(charArr)
}

// LeftRatateString rotates first k characters to the tail
func LeftRatateString(s string, k int) string {
	if s == "" || len(s) <= k || k == 0 {
		return s
	}
	words := []rune(s)
	lo, hi := 0, len(words)-1
	reverseString(words, lo, hi)
	reverseString(words, lo, hi-k)
	reverseString(words, hi-k+1, hi)

	return string(words)
}
