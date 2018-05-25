package interviews

import (
	"bytes"
	"sort"
	"strconv"
	"strings"

	"github.com/CAIJUNYI/GoAlgorithms/datastructure"
)

// FindMinPermutation returns the min permutation of array items
func FindMinPermutation(arr []int) string {
	intHeap := datastructure.IntHeap{
		Data: arr,
	}
	intHeap.Cmp = func(i, j int) bool {
		a := strconv.Itoa(intHeap.Data[i])
		b := strconv.Itoa(intHeap.Data[j])
		result := strings.Compare(a+b, b+a)
		if result < 0 {
			return true
		}
		return false
	}

	sort.Sort(intHeap)
	var b bytes.Buffer
	for _, e := range intHeap.Data {
		b.WriteString(strconv.Itoa(e))
	}
	return b.String()
}
