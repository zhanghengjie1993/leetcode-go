package problem0155

import "math"

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	n := len(this.minStack)
	if this.minStack[n-1] > x {
		this.minStack = append(this.minStack, x)
	}
}

func (this *MinStack) Pop() {
	n := len(this.stack)
	m := len(this.minStack)
	if this.stack[n-1] == this.minStack[m-1] {
		this.minStack = this.minStack[:m-1]
	}
	this.stack = this.stack[:n-1]
}

func (this *MinStack) Top() int {
	n := len(this.stack)
	return this.stack[n-1]
}

func (this *MinStack) GetMin() int {
	n := len(this.minStack)
	return this.minStack[n-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
