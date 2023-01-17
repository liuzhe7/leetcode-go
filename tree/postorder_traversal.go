/*
Author : Liu Zhe
Date   : 2023/1/17
*/

/**********************************************************************************
https://leetcode.cn/problems/binary-tree-postorder-traversal/
**********************************************************************************/

package tree

// f(root) = [f(root.left),f(root.right),root]
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	numsLeft := postorderTraversal(root.Left)
	numsRight := postorderTraversal(root.Right)
	var nums []int
	nums = append(nums, numsLeft...)
	nums = append(nums, numsRight...)
	nums = append(nums, root.Val)
	return nums
}
