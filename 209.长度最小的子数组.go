package leetcodego

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {

	l, r := 0, 0
	length := len(nums) + 1
	sum := 0
	for l <= r && r < len(nums) {
		sum += nums[r]
		if sum >= target {
			sum -= (nums[l] + nums[r])
			length = min(length, r-l+1)
			l++
		} else if sum < target {
			r++
		}
	}
	if length == len(nums)+1 {
		return 0
	}
	return length
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
