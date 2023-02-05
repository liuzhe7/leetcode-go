/*
Author : Liu Zhe
Date   : 2023/1/17
*/

/**********************************************************************************
https://leetcode.cn/problems/first-bad-version/
**********************************************************************************/

package halfintervalsearch

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	min := 1
	max := n
	for min < max {
		mid := min + (max-min)>>1
		if isBadVersion(mid) {
			max = mid
		} else {
			min = mid + 1
		}
	}
	return min
}
