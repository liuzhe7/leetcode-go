/*
Author : Liu Zhe
Date   : 2023/1/15
*/

/**********************************************************************************
https://leetcode.cn/problems/remove-linked-list-elements/
**********************************************************************************/

package recursion

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if nil == head {
		return head
	}

	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
