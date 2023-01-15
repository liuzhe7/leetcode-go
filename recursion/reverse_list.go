/*
Author : Liu Zhe
Date   : 2023/1/15
*/

/**********************************************************************************
https://leetcode.cn/problems/reverse-linked-list/submissions/
**********************************************************************************/

package recursion

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// f(head) =
// f(head.next)
// head.next.next = head
// head.next = nil

// head == nil; head.Next == nil
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
