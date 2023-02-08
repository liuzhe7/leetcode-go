/*
Author : Liu Zhe
Date   : 2023/2/8
*/

/**********************************************************************************
https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
**********************************************************************************/

package array

func firstUniqChar(s string) byte {
	hashMap := make(map[byte]int, len(s))
	for index := range s {
		hashMap[s[index]]++
	}
	for index := range s {
		if hashMap[s[index]] == 1 {
			return s[index]
		}
	}
	return ' '
}
