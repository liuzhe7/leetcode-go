/*
Author : Liu Zhe
Date   : 2023/1/17
*/

/**********************************************************************************
https://leetcode.cn/problems/search-insert-position/
**********************************************************************************/

package halfintervalsearch

func searchInsert(nums []int, target int) int {
	max := len(nums) - 1
	min := 0
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
	return max + 1
}
