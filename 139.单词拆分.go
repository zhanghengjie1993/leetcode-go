package leetcodego

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i <= len(s); i++ {
		for _, v := range wordDict {
			if i >= len(v) {
				if dp[i] {
					break
				}
				dp[i] = dp[i-len(v)] && (s[i-len(v):i] == v)
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end
