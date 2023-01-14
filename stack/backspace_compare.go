/*
Author : Liu Zhe
Date   : 2023/1/14
*/

/**********************************************************************************
https://leetcode.cn/problems/backspace-string-compare/
**********************************************************************************/

package stack

import "reflect"

func pop(stack []rune) {
	if len(stack) == 0 {
		return
	}
	stack = stack[:len(stack)-1]
}

func backspaceCompare(s string, t string) bool {
	var sOne, sTwo []rune
	for _, strOne := range s {
		switch strOne {
		case '#':
			if len(sOne) == 0 {
				continue
			}
			sOne = sOne[:len(sOne)-1]
		default:
			sOne = append(sOne, strOne)
		}
	}
	for _, strTwo := range t {
		switch strTwo {
		case '#':
			if len(sTwo) == 0 {
				continue
			}
			sTwo = sTwo[:len(sTwo)-1]
		default:
			sTwo = append(sTwo, strTwo)
		}
	}
	return reflect.DeepEqual(sTwo, sOne)
}
