package interviews

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var occurrence [256]int

// PrintFirstAppearingOnce prints first appearing once char
func PrintFirstAppearingOnce() {
	// init occurrence
	for i := 0; i < len(occurrence); i++ {
		occurrence[i] = -1
	}
	inputReader := bufio.NewReader(os.Stdin)
	allLength := 0
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// process input
		for i, ch := range []byte(input) {
			if occurrence[ch] == -1 {
				if allLength != 0 {
					occurrence[ch] = i + allLength - 1
				} else {
					occurrence[ch] = i
				}
			} else if occurrence[ch] >= 0 {
				occurrence[ch] = -2
			}
		}
		allLength += len(input)
		// statistic
		minIndex := math.MaxInt64
		var charOnce byte
		for i := 0; i < 256; i++ {
			if occurrence[i] >= 0 && occurrence[i] < minIndex {
				charOnce = byte(i)
				minIndex = occurrence[i]
			}
		}
		if minIndex == math.MaxInt64 {
			fmt.Println("no char appearing once")
		} else {
			fmt.Printf("'%c' appears at %d\n", charOnce, minIndex)
		}
	}
}
