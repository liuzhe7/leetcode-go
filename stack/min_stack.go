/*
Author : Liu Zhe
Date   : 2023/1/30
*/

/**********************************************************************************
https://leetcode.cn/problems/bao-han-minhan-shu-de-zhan-lcof/
**********************************************************************************/

package stack

type MinStack struct {
	minStack []int
	stack    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, x)
		return
	}
	if x <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, x)
	}

}

func (this *MinStack) Pop() {
	deleteElement := this.stack[len(this.stack)-1]
	if deleteElement == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
