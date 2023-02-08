/*
Author : Liu Zhe
Date   : 2023/2/8
*/

/**********************************************************************************
https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
**********************************************************************************/

package array

func minArray(numbers []int) int {
	min, max, mid := 0, len(numbers)-1, 0
	for max > min {
		mid = min + (max-min)>>1
		if numbers[mid] > numbers[max] {
			min = mid + 1
		} else if numbers[mid] < numbers[max] {
			max = mid
		} else if numbers[mid] == numbers[max] {
			max--
		}
	}
	return numbers[min]
}
