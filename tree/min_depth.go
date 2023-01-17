/*
Author : Liu Zhe
Date   : 2023/1/17
*/

/**********************************************************************************
https://leetcode.cn/problems/minimum-depth-of-binary-tree/
**********************************************************************************/

package tree

// f(root)=min(f(root.left),f(root.right)) + 1
// root == nil
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 {
		return right + 1
	}
	if right == 0 {
		return left + 1
	}
	if left > right {
		return right + 1
	}
	return left + 1
}
