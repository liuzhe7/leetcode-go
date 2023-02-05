/*
Author : Liu Zhe
Date   : 2023/1/30
*/

/**********************************************************************************
https://leetcode.cn/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
**********************************************************************************/

package stack

type CQueue struct {
	stackOne []int
	stackTwo []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stackOne = append(this.stackOne, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stackTwo) == 0 {
		for index := len(this.stackOne) - 1; index >= 0; index-- {
			this.stackTwo = append(this.stackTwo, this.stackOne[index])
		}
		this.stackOne = this.stackOne[:0]
	}
	if len(this.stackTwo) == 0 {
		return -1
	}
	res := this.stackTwo[len(this.stackTwo)-1]
	this.stackTwo = this.stackTwo[:len(this.stackTwo)-1]
	return res
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
