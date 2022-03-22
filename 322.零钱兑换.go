package leetcodego

import "math"

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for _, v := range coins {
		for j := 0; j <= amount; j++ {
			if j >= v {
				dp[j] = min322(dp[j], dp[j-v]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]

}

func min322(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
