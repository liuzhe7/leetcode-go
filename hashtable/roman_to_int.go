/*
Author : Liu Zhe
Date   : 2023/1/17
*/

/**********************************************************************************
https://leetcode.cn/problems/roman-to-integer/
**********************************************************************************/

package hashtable

func romanToInt(s string) int {
	mapping := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	var res int
	for index := 0; index < len(s); {
		valueLeft := s[index]
		if index == len(s)-1 {
			res += mapping[string(valueLeft)]
			return res
		}
		valueRight := s[index+1]
		v, ok := mapping[string(valueLeft)+string(valueRight)]
		if ok {
			res += v
			index += 2
		} else {
			index++
			res += mapping[string(valueLeft)]
		}
	}
	return res
}
