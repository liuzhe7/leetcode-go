/*
Author : Liu Zhe
Date   : 2023/2/5
*/

/**********************************************************************************
https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
**********************************************************************************/

package string

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
