/*
Author : Liu Zhe
Date   : 2023/1/17
*/

/**********************************************************************************
https://leetcode.cn/problems/binary-tree-inorder-traversal/
**********************************************************************************/

package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// f(root) = [f(root.left), root, f(root.right)]
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var nums []int
	numsLeft := inorderTraversal(root.Left)
	nums = append(numsLeft, root.Val)
	numsRight := inorderTraversal(root.Right)
	nums = append(nums, numsRight...)
	return nums
}
