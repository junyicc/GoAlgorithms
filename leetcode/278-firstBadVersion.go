package leetcode

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

// Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

// You are given an API bool isBadVersion(version) which will return whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.

func firstBadVersion(n int) int {
	if n < 1 {
		return -1
	}
	lo, hi := 1, n
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if !isBadVersion(mi) {
			lo = mi + 1
		} else {
			if mi == 1 || !isBadVersion(mi-1) {
				return mi
			} else {
				hi = mi - 1
			}
		}
	}
	if lo > hi {
		return n
	}
	return -1
}

func isBadVersion(n int) bool {
	bad := 4
	if n == bad {
		return true
	}
	return false
}
