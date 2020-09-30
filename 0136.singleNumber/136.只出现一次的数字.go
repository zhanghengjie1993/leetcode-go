package problem0136

import "sort"

/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i = i + 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
	// var ans int
	// hmap := make(map[int]bool)
	// for _, v := range nums {
	// 	if _, ok := hmap[v]; ok {
	// 		delete(hmap, v)
	// 	} else {
	// 		hmap[v] = true
	// 	}
	// }
	// for k := range hmap {
	// 	ans = k
	// }
	// return ans
}

// @lc code=end
