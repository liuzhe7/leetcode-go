/*
Author : Liu Zhe
Date   : 2023/2/2
*/

/**********************************************************************************
https://leetcode.cn/problems/length-of-last-word/
**********************************************************************************/

package string

func lengthOfLastWord(s string) int {
	var res int
	for index := len(s) - 1; index >= 0; index-- {
		if s[index] != ' ' {
			res++
		}
		if res != 0 && s[index] == ' ' {
			return res
		}
		if res != 0 && index == 0 {
			return res
		}
	}
	return len(s)
}
