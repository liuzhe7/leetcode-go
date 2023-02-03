/*
Author : Liu Zhe
Date   : 2023/2/3
*/

/**********************************************************************************
https://leetcode.cn/problems/binary-search/
**********************************************************************************/

package halfintervalsearch

func search(nums []int, target int) int {
	min := 0
	max := len(nums) - 1

	for min <= max {
		mid := min + (max-min)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			max = mid - 1
		}
		if nums[mid] < target {
			min = mid + 1
		}
	}
	return -1
}
