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
	reverseArray := make([]int, len(array), len(array))
	reverseArrayIndex := 0
	for index := len(array) - 1; index >= 0; index-- {
		reverseArray[reverseArrayIndex] = array[index]
		reverseArrayIndex++
	}
	return reverseArray
}
