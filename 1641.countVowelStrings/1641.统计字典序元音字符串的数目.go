package problem1641

/*
 * @lc app=leetcode.cn id=1641 lang=golang
 *
 * [1641] 统计字典序元音字符串的数目
 */

// @lc code=start
func countVowelStrings(n int) int {

	dp := make([][]int, 5)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	dp[0][1] = 1
	dp[1][1] = 1
	dp[2][1] = 1
	dp[3][1] = 1
	dp[4][1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < 5; j++ {
			sum := 0
			for k := j; k < 5; k++ {
				sum = sum + dp[k][i-1]
			}
			dp[j][i] = sum
		}
	}
	ans := dp[0][n] + dp[1][n] + dp[2][n] + dp[3][n] + dp[4][n]
	return ans
}

// @lc code=end
