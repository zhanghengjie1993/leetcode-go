package leetcodego

/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {

	dp := make([]int, target+1)
	dp[0] = 1
	for j := 0; j <= target; j++ {
		for _, v := range nums {
			if v <= j {
				dp[j] += dp[j-v]
			}
		}
	}
	return dp[target]

}

// @lc code=end
