/*
Author : Liu Zhe
Date   : 2023/1/14
*/

/**********************************************************************************
https://leetcode.cn/problems/backspace-string-compare/
**********************************************************************************/

package stack

import "reflect"

func backspaceCompare(s string, t string) bool {
	sOne := &StackRune{}
	sTwo := &StackRune{}
	for _, strOne := range s {
		switch strOne {
		case '#':
			sOne.Pop()
		default:
			sOne.Insert(strOne)
		}
	}
	for _, strTwo := range t {
		switch strTwo {
		case '#':
			sTwo.Pop()
		default:
			sTwo.Insert(strTwo)
		}
	}
	return reflect.DeepEqual(sTwo, sOne)
}
