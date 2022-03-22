package leetcodego279

import "math"

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {

	dp := make([]int, n+1)

	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for j := 0; j <= n; j++ {
		for i := 1; i*i <= j; i++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
