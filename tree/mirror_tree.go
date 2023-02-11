/*
Author : Liu Zhe
Date   : 2023/2/11
*/

/**********************************************************************************
https://leetcode.cn/problems/er-cha-shu-de-jing-xiang-lcof/
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
// f(root)= root.left=f(root.right); root.right=f(root.left)
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	tmp := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(tmp)
	return root
}
