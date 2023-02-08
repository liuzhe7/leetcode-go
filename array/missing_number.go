/*
Author : Liu Zhe
Date   : 2023/2/8
*/

/**********************************************************************************
https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/
**********************************************************************************/

package array

func missingNumber(nums []int) int {
	if nums[0] == 1 {
		return 0
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 != nums[i+1] {
			return nums[i] + 1
		}
	}
	return nums[len(nums)-1] + 1
}
