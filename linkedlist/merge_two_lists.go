/*
Author : Liu Zhe
Date   : 2023/1/13
*/

/**********************************************************************************
https://leetcode.cn/problems/merge-two-sorted-lists/
**********************************************************************************/

package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	nodeZero := new(ListNode)
	dynamicPointer := nodeZero
	for {
		if list1 == nil {
			dynamicPointer.Next = list2
			break
		}
		if list2 == nil {
			dynamicPointer.Next = list1
			break
		}
		if list1.Val <= list2.Val {
			dynamicPointer.Next = list1
			list1 = list1.Next
		} else {
			dynamicPointer.Next = list2
			list2 = list2.Next
		}
		dynamicPointer = dynamicPointer.Next
	}
	return nodeZero.Next
}
