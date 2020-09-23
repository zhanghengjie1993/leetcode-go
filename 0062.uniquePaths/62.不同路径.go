package problem0062

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	// var rectangle = make([][]int, m)
	// for i := 0; i < m; i++ {
	// 	rectangle[i] = make([]int, n)
	// 	for j := 0; j < n; j++ {
	// 		if i == 0 || j == 0 {
	// 			rectangle[i][j] = 1
	// 		} else {
	// 			rectangle[i][j] = rectangle[i-1][j] + rectangle[i][j-1]
	// 		}
	// 	}
	// }
	// return rectangle[m-1][n-1]
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] = 1
			} else {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}
	return dp[n-1]
}

// @lc code=end
