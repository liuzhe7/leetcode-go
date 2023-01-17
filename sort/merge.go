/*
Author : Liu Zhe
Date   : 2023/1/17
*/

/**********************************************************************************
https://leetcode.cn/problems/merge-sorted-array/
**********************************************************************************/

package sort

import "sort"

func merge(nums1 []int, m int, nums2 []int, _ int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}
