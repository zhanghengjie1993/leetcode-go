package problem0139

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	dp := make([]int, 1)
	for i := 0; i < len(s); i++ {
		for _, v := range wordDict {
			for j := range dp {
				if s[dp[j]:i+1] == v {
					dp = append(dp, i+1)
				}
			}
		}
	}

	return dp[len(dp)-1] == len(s)
}

// @lc code=end
