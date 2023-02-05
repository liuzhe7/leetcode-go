/*
Author : Liu Zhe
Date   : 2023/1/18
*/

/**********************************************************************************
https://leetcode.cn/problems/same-tree/
**********************************************************************************/

package tree

// f(p,q) = f(p.left,q.left) && f(p.right,q.right) && p.val==q.val
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) && p.Val == q.Val
}
