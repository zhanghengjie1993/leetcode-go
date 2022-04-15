package leetcodego

/*
 * @lc app=leetcode.cn id=583 lang=golang
 *
 * [583] 两个字符串的删除操作
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1)+1, len(word2)+1
	dp := make([][]int, n1)
	for i := range dp {
		dp[i] = make([]int, n2)
	}

	for i, c := range word1 {
		for j, v := range word2 {
			if c == v {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return len(word1) - dp[n1-1][n2-1] + len(word2) - dp[n1-1][n2-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
