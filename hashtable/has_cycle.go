/*
Author : Liu Zhe
Date   : 2023/1/17
*/

/**********************************************************************************
https://leetcode.cn/problems/linked-list-cycle/
**********************************************************************************/

package hashtable

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	mapping := map[*ListNode]struct{}{}
	for head != nil {
		_, ok := mapping[head]
		if ok {
			return true
		}
		mapping[head] = struct{}{}
		head = head.Next
	}
	return false
}
