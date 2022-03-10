package leetcodego /*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {

	dp := make([]int, amount+1)

	dp[0] = 1
	for i := range coins {
		for j := 0; j <= amount; j++ {
			if j >= coins[i] {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[amount]

}

// @lc code=end
