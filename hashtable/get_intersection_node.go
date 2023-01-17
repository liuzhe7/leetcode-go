/*
Author : Liu Zhe
Date   : 2023/1/17
*/

/**********************************************************************************
https://leetcode.cn/problems/intersection-of-two-linked-lists/
**********************************************************************************/

package hashtable

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mapping := map[*ListNode]struct{}{}
	for headA != nil {
		mapping[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		_, ok := mapping[headB]
		if ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
