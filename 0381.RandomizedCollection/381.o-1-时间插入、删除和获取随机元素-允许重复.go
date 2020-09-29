package problem0381

import "math/rand"

/*
 * @lc app=leetcode.cn id=381 lang=golang
 *
 * [381] O(1) 时间插入、删除和获取随机元素 - 允许重复
 */

// @lc code=start
type RandomizedCollection struct {
	hmap map[int][]int
	list []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		hmap: make(map[int][]int),
		list: []int{},
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	flag := true
	this.list = append(this.list, val)
	index := len(this.list) - 1
	if _, ok := this.hmap[val]; ok {
		flag = false
	}
	this.hmap[val] = append(this.hmap[val], index)
	return flag
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if _, ok := this.hmap[val]; !ok {
		return false
	}
	length := len(this.list)
	remove := this.hmap[val][0]
	this.list[remove] = this.list[length-1]
	if len(this.hmap[val]) == 1 {
		delete(this.hmap, val)
	} else {
		this.hmap[val] = this.hmap[val][1:]
	}
	last := this.list[length-1]
	if _, ok := this.hmap[last]; ok {
		this.hmap[last] = append(this.hmap[last], remove)
	} else {
		this.hmap[last] = []int{remove}
	}
	if remove == length-1 {
		delete(this.hmap, val)
		return true
	}
	for i, v := range this.hmap[last] {
		if v == length-1 {
			this.hmap[last] = append(this.hmap[last][0:i], this.hmap[last][i+1:]...)
		}
	}
	this.list = this.list[:length-1]
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
