/*
Author : Liu Zhe
Date   : 2023/2/13
*/

/**********************************************************************************
https://leetcode.cn/problems/maximum-depth-of-binary-tree/
**********************************************************************************/

package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// f(r) = max(f(r.right),f(r.left)) + 1
// f(nil) = 0
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	rightMax := maxDepth(root.Right)
	leftMax := maxDepth(root.Left)
	if rightMax > leftMax {
		return rightMax + 1
	}
	return leftMax + 1
}
