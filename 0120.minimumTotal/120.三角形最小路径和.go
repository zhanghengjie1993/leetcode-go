package problem0120

import "math"

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([]int, n)
	f[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		f[i] = f[i-1] + triangle[i][i]
		for j := i - 1; j > 0; j-- {
			f[j] = min(f[j-1], f[j]) + triangle[i][j]
		}
		f[0] += triangle[i][0]
	}
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		ans = min(ans, f[i])
	}
	return ans
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
