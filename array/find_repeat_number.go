/*
Author : Liu Zhe
Date   : 2023/2/6
*/

/**********************************************************************************
https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
**********************************************************************************/

package array

func findRepeatNumber(nums []int) int {
	var numberTable = make(map[int]bool, len(nums))
	for index := range nums {
		_, ok := numberTable[nums[index]]
		if ok {
			return nums[index]
		}
		numberTable[nums[index]] = true
	}
	return -1
}
