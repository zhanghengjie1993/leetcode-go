package main

import (
	"fmt"
	"math/rand"
)

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
	arr := this.hmap[val]
	arr = append(arr, val)
	this.hmap[val] = arr
	return true
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
	return list[index-1]
}

func main() {
	obj := Constructor()
	obj.Insert(1)
	obj.Insert(10)
	obj.Insert(10)
	obj.Insert(100)
	for i := 0; i < 20; i++ {
		fmt.Println(obj.GetRandom())
	}
	// obj.GetRandom()
}
