/*
Author : Liu Zhe
Date   : 2023/1/13
*/

/**********************************************************************************
https://leetcode.cn/problems/remove-element/
**********************************************************************************/

package array

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] == val {
			nums[left] = nums[right]
			right--
		} else {
			left++
		}
	}
	return left
}
