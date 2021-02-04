package problem0063

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if i == 0 && j == 0 {
				dp[0] = 1
				continue
			}
			if j != 0 {
				dp[j] = dp[j-1] + dp[j]
			}
		}
	}
	return dp[n-1]
}

// @lc code=end
