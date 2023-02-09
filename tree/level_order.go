/*
Author : Liu Zhe
Date   : 2023/2/9
*/

/**********************************************************************************
https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var (
		res   [][]int
		layer int
	)
	for len(queue) > 0 {
		res = append(res, []int{})
		queueLength := len(queue)
		for index := 0; index < queueLength; index++ {
			res[layer] = append(res[layer], queue[index].Val)
			if queue[index].Left != nil {
				queue = append(queue, queue[index].Left)
			}
			if queue[index].Right != nil {
				queue = append(queue, queue[index].Right)
			}
		}
		layer++
		queue = queue[queueLength:]
	}
	return res
}
