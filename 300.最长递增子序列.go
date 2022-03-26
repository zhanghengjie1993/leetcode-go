package leetcodego

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {

	if len(nums) == 1 {
		return 1
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	temp := 0
	ans := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				temp = dp[j] + 1
			} else {
				temp = 1
			}
			dp[i] = max(temp, dp[i])
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
