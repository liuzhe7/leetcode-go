/*
Author : Liu Zhe
Date   : 2023/1/13
*/

/**********************************************************************************
https://leetcode.cn/problems/two-sum/
**********************************************************************************/

package array

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for index := range nums {
		if p, ok := hashTable[target-nums[index]]; ok {
			return []int{p, index}
		}
		hashTable[nums[index]] = index
	}
	return nil
}
