/*
Author : Liu Zhe
Date   : 2023/1/13
*/

/**********************************************************************************
https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
**********************************************************************************/

package linkedlist

func reversePrint(head *ListNode) []int {
	var array []int
	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}
	for left, right := 0, len(array)-1; left < right; {
		array[left], array[right] = array[right], array[left]
		left++
		right--
	}
	return array
}
