/*
Author : Liu Zhe
Date   : 2023/1/14
*/

/**********************************************************************************
https://leetcode.cn/problems/valid-parentheses/
**********************************************************************************/

package stack

func isValid(s string) bool {
	var stack []rune
	for _, char := range s {
		switch char {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			if len(stack) == 0 {
				return false
			}
			if char != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
