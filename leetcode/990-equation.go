package leetcode

import "github.com/CAIJUNYI/GoAlgorithms/datastructure"

// Given an array equations of strings that represent relationships between variables, each string equations[i] has length 4 and takes one of two different forms: "a==b" or "a!=b".  Here, a and b are lowercase letters (not necessarily different) that represent one-letter variable names.
//
// Return true if and only if it is possible to assign integers to variable names so as to satisfy all the given equations.
//
// Example 1:
// Input: ["a==b","b!=a"]
// Output: false
// Explanation: If we assign say, a = 1 and b = 1, then the first equation is satisfied, but not the second.  There is no way to assign the variables to satisfy both equations.
//
// Example 2:
// Input: ["b==a","a==b"]
// Output: true
// Explanation: We could assign a = 1 and b = 1 to satisfy both equations.
//
// Example 3:
// Input: ["a==b","b==c","a==c"]
// Output: true
//
// Example 4:
// Input: ["a==b","b!=c","c==a"]
// Output: false

func equationsPossible(equations []string) bool {
	if len(equations) < 1 {
		return true
	}

	uf := datastructure.NewUnionFindInt(26)
	// build connection
	for _, equation := range equations {
		if equation[1] == '=' {
			uf.Union(int(equation[0]-'a'), int(equation[3]-'a'))
		}
	}
	// check connection
	for _, equation := range equations {
		if equation[1] == '!' {
			if uf.Connected(int(equation[0]-'a'), int(equation[3]-'a')) {
				return false
			}
		}
	}
	return true
}
