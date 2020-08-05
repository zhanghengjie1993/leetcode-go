package problem0011

import "math"

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area, ans := 0, 0
	for left < right {
		area = min(height[left], height[right]) * (right - left)
		if ans < area {
			ans = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	math.Max(1, 2)
	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end
