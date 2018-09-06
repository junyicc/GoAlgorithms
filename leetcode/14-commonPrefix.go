package leetcode

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:
Input: ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Note:
All given inputs are in lowercase letters a-z.
*/
func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	commonPrefix := strs[0]
	strsLen := len(strs)
	for i := 1; i < strsLen; i++ {
		commonPrefix = longestCommonPrefixOfTwoStr(commonPrefix, strs[i])
	}
	return commonPrefix
}

func longestCommonPrefixOfTwoStr(str1, str2 string) string {
	if str1 == "" || str2 == "" {
		return ""
	}
	str1Len := len(str1)
	str2Len := len(str2)
	i, j := 0, 0
	for i < str1Len && j < str2Len {
		if str1[i] != str2[j] {
			break
		}
		i++
		j++
	}
	return str1[:i]
}

func longestCommonSubstring(str1, str2 string) int {
	if str1 == "" || str2 == "" {
		return 0
	}
	str1Len := len(str1)
	str2Len := len(str2)
	strMatrix := make([][]int, str1Len)
	for i := 0; i < str1Len; i++ {
		strMatrix[i] = make([]int, str2Len)
	}

	for i := 0; i < str1Len; i++ {
		for j := 0; j < str2Len; j++ {
			if str1[i] == str2[j] {
				if i > 0 && j > 0 {
					strMatrix[i][j] = strMatrix[i-1][j-1] + 1
				} else {
					strMatrix[i][j] = 1
				}
			}
		}
	}

	maxLen := 0
	for i := 0; i < str1Len; i++ {
		for j := 0; j < str2Len; j++ {
			if strMatrix[i][j] > maxLen {
				maxLen = strMatrix[i][j]
			}
		}
	}
	return maxLen
}
