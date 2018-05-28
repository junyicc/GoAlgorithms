package interviews

import (
	"fmt"
	"math"
)

var diceDim = 6

// PrintDiceProbability prints the probabilities of the sum of n dices that are ge n && le n*diceDim
func PrintDiceProbability(n int) {
	if n < 1 {
		return
	}
	length := diceDim*n + 1
	p1 := make([]int, length)
	p2 := make([]int, length)
	p := [2][]int{}
	p[0] = p1
	p[1] = p2
	// the first dice
	flag := 0
	for i := 1; i <= diceDim; i++ {
		p[flag][i] = 1
	}
	for i := 2; i <= n; i++ { // i is the dice index
		for j := 0; j < i; j++ {
			p[1-flag][j] = 0
		}

		for j := i; j <= diceDim*i; j++ {
			p[1-flag][j] = 0
			for k := 1; k <= j && k <= diceDim; k++ {
				p[1-flag][j] += p[flag][j-k]
			}
		}
		flag = 1 - flag
	}
	total := math.Pow(float64(diceDim), float64(n))
	for i := n; i <= diceDim*n; i++ {
		ratio := float64(p[flag][i]) / total
		fmt.Printf("%d: %e\n", i, ratio)
	}
}
