/*
Author : Liu Zhe
Date   : 2023/1/17
*/

/**********************************************************************************
https://leetcode.cn/problems/sqrtx/
**********************************************************************************/

package halfintervalsearch

func mySqrt(x int) int {
	min := 0
	max := x
	for min <= max {
		mid := min + (max-min)>>1
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			min = mid + 1
		}
		if mid*mid > x {
			max = mid - 1
		}
	}
	return min - 1
}
