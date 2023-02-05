/*
Author : Liu Zhe
Date   : 2023/1/29
*/

/**********************************************************************************
https://leetcode.cn/problems/palindrome-number/
**********************************************************************************/

package number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var reverseInt int
	var xCopy = x
	for {
		lastNumber := x % 10
		reverseInt = reverseInt*10 + lastNumber
		if x < 10 {
			break
		}
		x = x / 10
	}
	return reverseInt == xCopy
}
