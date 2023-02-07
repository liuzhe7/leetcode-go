/*
Author : Liu Zhe
Date   : 2023/2/7
*/

/**********************************************************************************
https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
**********************************************************************************/

package array

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	var min, max, mid = 0, len(nums) - 1, 0
	for max >= min {
		mid = min + (max-min)>>1
		if nums[mid] == target {
			break
		}
		if nums[mid] > target {
			max = mid - 1
		}
		if nums[mid] < target {
			min = mid + 1
		}
	}
	if nums[mid] != target {
		return 0
	}
	for mid >= 0 {
		if nums[mid] == target {
			mid--
			continue
		}
		if nums[mid] != target {
			break
		}
	}
	res := 0
	for {
		mid++
		if mid > len(nums)-1 {
			return res
		}
		if nums[mid] == target {
			res++
		}
		if nums[mid] != target {
			return res
		}
	}
}
