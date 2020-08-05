package problem0027

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	n := 0
	for index := 0; index < len(nums); index++ {
		if nums[index] != val {
			nums[n] = nums[index]
			n++
		}
	}
	return n
}

// @lc code=end
