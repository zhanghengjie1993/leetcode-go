package main

import "fmt"

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
	fmt.Println(len(this.arr))
	if len(this.arr) == 0 {
		return true
	}
	return false
}
func main() {
	obj := Constructor()
	fmt.Println(obj.Empty())
}
