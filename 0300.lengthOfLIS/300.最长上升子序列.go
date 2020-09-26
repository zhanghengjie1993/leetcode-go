package problem0300

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	ans := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			} else {
				dp[i] = max(dp[i], 1)
			}
		}
		ans = max(ans, dp[i])
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
