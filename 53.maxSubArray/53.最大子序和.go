package problem0053

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// n := len(nums)
	// var dp = make([]int, n)
	// dp[0] = nums[0]
	// ans := dp[0]
	// for i := 1; i < n; i++ {
	// 	dp[i] = max(nums[i]+dp[i-1], nums[i])
	// 	ans = max(ans, dp[i])
	// }
	// return ans
	dp := nums[0]
	ans := dp
	for i := 1; i < len(nums); i++ {
		dp = max(dp+nums[i], nums[i])
		ans = max(ans, dp)
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
