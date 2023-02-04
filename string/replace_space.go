/*
Author : Liu Zhe
Date   : 2023/2/4
*/

/**********************************************************************************
https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
**********************************************************************************/

package string

import "strings"

func replaceSpace(s string) string {
	sb := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			sb.Write([]byte{'%', '2', '0'})
		} else {
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}
