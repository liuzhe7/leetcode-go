/*
Author : Liu Zhe
Date   : 2023/1/13
*/

/*
*********************************************************************************
https://leetcode.cn/problems/remove-duplicates-from-sorted-array/
*********************************************************************************
*/

package array

func removeDuplicates(nums []int) int {
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow-1] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	nums = nums[:slow]
	return slow
}
