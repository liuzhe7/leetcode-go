/*
Author : Liu Zhe
Date   : 2023/1/14
*/

/**********************************************************************************
https://leetcode.cn/problems/implement-stack-using-queues/
**********************************************************************************/

package stack

type MyStack struct {
	queueOne, queueTwo []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queueTwo = append(this.queueTwo, x)
	for index := range this.queueOne {
		this.queueTwo = append(this.queueTwo, this.queueOne[index])
	}
	this.queueOne = this.queueTwo
	this.queueTwo = []int{}
}

func (this *MyStack) Pop() int {
	popInt := this.queueOne[0]
	this.queueOne = this.queueOne[1:]
	return popInt
}

func (this *MyStack) Top() int {
	return this.queueOne[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queueOne) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
