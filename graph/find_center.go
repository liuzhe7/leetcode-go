/*
Author : Liu Zhe
Date   : 2023/1/31
*/

/**********************************************************************************
https://leetcode.cn/problems/find-center-of-star-graph/
**********************************************************************************/

package graph

func findCenter(edges [][]int) int {
	n := len(edges) + 1
	way := make([]int, n+1)
	for index := range edges {
		way[edges[index][0]]++
		way[edges[index][1]]++
	}
	for i := 1; i <= n; i++ {
		if way[i] == n-1 {
			return i
		}
	}
	return -1
}
