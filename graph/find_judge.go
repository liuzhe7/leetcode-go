/*
Author : Liu Zhe
Date   : 2023/1/31
*/

/**********************************************************************************
https://leetcode.cn/problems/find-the-town-judge/submissions/
**********************************************************************************/

package graph

func findJudge(n int, trust [][]int) int {
	inDegrees := make([]int, n+1)
	outDegrees := make([]int, n+1)
	for index := range trust {
		inDegrees[trust[index][1]]++
		outDegrees[trust[index][0]]++
	}
	for i := 1; i <= n; i++ {
		if inDegrees[i] == n-1 && outDegrees[i] == 0 {
			return i
		}
	}
	return -1
}
