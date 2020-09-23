package problem0064

/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	// dp := make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([]int, m)
	// 	for j := 0; j < m; j++ {
	// 		if i == 0 && j == 0 {
	// 			dp[i][j] = grid[0][0]
	// 		} else if i == 0 {
	// 			dp[i][j] = dp[i][j-1] + grid[i][j]
	// 		} else if j == 0 {
	// 			dp[i][j] = dp[i-1][j] + grid[i][j]
	// 		} else {
	// 			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
	// 		}
	// 	}
	// }
	// return dp[n-1][m-1]
	dp := make([]int, m)
	dp[0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == 0 {
				dp[j] = dp[j] + grid[i][0]
			} else {
				dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
