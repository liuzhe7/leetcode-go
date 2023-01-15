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

// f(n,val) =
// list = f(n.next,val)
// if n.val != val {
// n.next = list
// list = n
// }
// return list

// n == nil
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	list := removeElements(head.Next, val)
	if head.Val != val {
		head.Next = list
		list = head
	}
	return list
}
