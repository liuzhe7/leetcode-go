/*
Author : Liu Zhe
Date   : 2023/1/27
*/

/**********************************************************************************
https://leetcode.cn/problems/longest-common-prefix/submissions/
**********************************************************************************/

package string

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j])-1 < i {
				return strs[0][:i]
			}
			if strs[j][i] != strs[j-1][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
