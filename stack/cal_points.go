/*
Author : Liu Zhe
Date   : 2023/1/14
*/

/**********************************************************************************
https://leetcode.cn/problems/baseball-game/
**********************************************************************************/

package stack

import "strconv"

func calPoints(operations []string) int {
	var stack []int
	var points int
	for index := range operations {
		switch operations[index] {
		case "C":
			points -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			continue
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		default:
			point, err := strconv.Atoi(operations[index])
			if err != nil {
				panic(err)
			}
			stack = append(stack, point)
		}
		points += stack[len(stack)-1]
	}
	return points
}
