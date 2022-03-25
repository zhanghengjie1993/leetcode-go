package leetcodego

/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	if len(prices) == 1 {
		return 0
	}

	dp := make([][2]int, len(prices))

	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][0]+prices[i]-fee, dp[i-1][1])
	}

	return dp[len(prices)-1][1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
