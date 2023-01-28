/*
Author : Liu Zhe
Date   : 2023/1/28
*/

/**********************************************************************************
https://leetcode.cn/problems/running-sum-of-1d-array/submissions/
**********************************************************************************/

package array

func runningSum(nums []int) []int {
	for index := 1; index < len(nums); index++ {
		nums[index] = nums[index] + nums[index-1]
	}
	return nums
}
