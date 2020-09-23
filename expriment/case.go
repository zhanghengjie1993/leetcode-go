package main

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
				dp[i] = 1
			} else {
				dp[i] = dp[i] + dp[i-1]
			}
		}
	}
	return dp[n-1]
}

func main() {
	uniquePaths(3, 2)
}
