package leetcodego

/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {

	dp := make([][]bool, len(s))

	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	num := 0

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if (j - i) <= 1 {
					num++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					num++
					dp[i][j] = true
				}

			}
		}
	}
	return num

}

// @lc code=end
