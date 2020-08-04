package problem0026

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	left, right, size := 0, 1, len(nums)
	for ; right < size; right++ {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}

// @lc code=end
