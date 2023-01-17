/*
Author : Liu Zhe
Date   : 2023/1/17
*/

/**********************************************************************************
https://leetcode.cn/problems/binary-tree-preorder-traversal/submissions/
**********************************************************************************/

package tree

// f(root) = [root,f(root.left),f(root.right)]
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var nums []int
	nums = append(nums, root.Val)
	numsLeft := preorderTraversal(root.Left)
	numsRight := preorderTraversal(root.Right)
	nums = append(nums, numsLeft...)
	nums = append(nums, numsRight...)
	return nums
}
