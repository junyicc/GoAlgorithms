package sort

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(values []int, left, right int) []int {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}

	}

	values[p] = temp

	if p-left > 1 {
		// quickSort(values, left, p-1)
		go quickSort(values, left, p-1)
	}
	if right-p > 1 {
		// quickSort(values, p+1, right)
		go quickSort(values, p+1, right)
	}
	return values
}

// QuickSort algorithm
func QuickSort(values []int) []int {
	return quickSort(values, 0, len(values)-1)
}

// TestQuickSort tests QuickSort algorithm
func TestQuickSort() {
	ceshi := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		ceshi[i] = rand.Intn(100)
	}
	startTime := time.Now()
	sortedCeshi := QuickSort(ceshi)
	endTime := time.Since(startTime)
	fmt.Println("after sort", ceshi)
	fmt.Println("count time ", endTime)
	fmt.Println(sortedCeshi)
}
