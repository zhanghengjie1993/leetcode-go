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
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	var rc *RandomizedCollection = new(RandomizedCollection)
	rc.hmap = make(map[int][]int)
	return *rc
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	flag := true
	if _, ok := this.hmap[val]; ok {
		flag = false
	}
	arr := this.hmap[val]
	arr = append(arr, val)
	this.hmap[val] = arr
	return flag
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if _, ok := this.hmap[val]; !ok {
		return false
	}
	if len(this.hmap[val]) == 1 {
		delete(this.hmap, val)
		return true
	}
	arr := this.hmap[val]
	arr = arr[:len(arr)-1]
	this.hmap[val] = arr
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	rand.Seed(1)
	list := []int{}
	for _, v := range this.hmap {
		list = append(list, v...)
	}
	index := rand.Intn(len(list))
	return list[index]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
