/*
Author : Liu Zhe
Date   : 2023/1/18
*/

/**********************************************************************************
https://leetcode.cn/problems/symmetric-tree/
**********************************************************************************/

package tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// f(p,q) = f(p.left,q.right) && f(p.right,q.left) && p.val==q.val
// p.Right == nil && q.Left == nil;return true
func isMirror(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	return isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left) && p.Val == q.Val
}
