/*
Author : Liu Zhe
Date   : 2023/1/14
*/

/**********************************************************************************
https://leetcode.cn/problems/valid-parentheses/
**********************************************************************************/

package stack

func isValid(s string) bool {
	stack := &StackRune{}
	for _, char := range s {
		switch char {
		case '(':
			stack.Insert(')')
		case '[':
			stack.Insert(']')
		case '{':
			stack.Insert('}')
		default:
			c, err := stack.Pop()
			if err != nil {
				return false
			}
			if c != char {
				return false
			}
		}
	}
	return len(stack.list) == 0
}
