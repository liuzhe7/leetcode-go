/*
Author : Liu Zhe
Date   : 2023/1/15
*/

/**********************************************************************************
https://leetcode.cn/problems/fibonacci-number/
**********************************************************************************/

package recursion

// f(n) = f(n-1) + f(n-2)

// f(0) return 0
// f(1) return 1
var record = map[int]int{}

func fib(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	score, ok := record[n]
	if ok {
		return score
	}
	record[n] = fib(n-1) + fib(n-2)
	return record[n]
}
