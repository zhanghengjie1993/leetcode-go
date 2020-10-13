package problem0225

/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
type MyStack struct {
	arr []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{make([]int, 0)}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.arr = append(this.arr, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	n := len(this.arr)
	ans := this.arr[n-1]
	this.arr = this.arr[:n-1]
	return ans
}

/** Get the top element. */
func (this *MyStack) Top() int {
	n := len(this.arr)
	return this.arr[n-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.arr) == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
