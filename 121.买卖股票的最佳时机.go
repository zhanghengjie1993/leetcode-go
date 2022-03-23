package leetcodego

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {

	ans := 0
	dp := 0

	for i := 1; i < len(prices); i++ {
		dp = dp + prices[i] - prices[i-1]
		if dp < 0 {
			dp = 0
		}
		ans = max(ans, dp)
	}
	return ans
}

// @lc code=end
