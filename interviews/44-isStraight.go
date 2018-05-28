package interviews

import (
	"sort"
)

// IsStraight returns true if the array is straight of squeezer where 1 for A, 2~10, 11 for Jack, 12 for Queen, 13 for King, and 0 for Joker
func IsStraight(arr []int) bool {
	if arr == nil || len(arr) < 5 {
		return false
	}

	// sort the array
	intSlice := sort.IntSlice(arr)
	sort.Sort(intSlice)
	// counting the Jokers
	jokerCnt := 0
	for _, e := range intSlice {
		if e == 0 {
			jokerCnt++
		}
	}
	// counting the gaps
	gapCnt := 0
	for i := 1; i < intSlice.Len(); i++ {
		if intSlice[i] == intSlice[i-1] && intSlice[i] != 0 {
			return false
		} else if intSlice[i]-intSlice[i-1] > 1 && intSlice[i-1] != 0 {
			gapCnt += (intSlice[i] - intSlice[i-1] - 1)
		}
	}

	if jokerCnt >= gapCnt {
		return true
	}
	return false
}
