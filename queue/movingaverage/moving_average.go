/*
Author : Liu Zhe
Date   : 2023/1/15
*/

/**********************************************************************************
https://leetcode.cn/problems/qIsx9U/
**********************************************************************************/

package movingaverage

type MovingAverage struct {
	width, sum int
	queue      []int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		width: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.queue) == this.width {
		this.sum -= this.queue[0]
		this.queue = this.queue[1:]
	}
	this.sum += val
	this.queue = append(this.queue, val)
	return float64(this.sum) / float64(len(this.queue))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
