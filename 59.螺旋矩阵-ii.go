package leetcodego

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	i_min, i_max, j_min, j_max := 0, n-1, 0, n-1
	i, j := i_min, j_min
	for k := 1; k <= n*n; {
		if i == i_min && j == j_min {
			for ; j <= j_max; j++ {
				dp[i][j] = k
				k++
			}
			j = j_max
			i_min++
			i = i_min
		}
		if i == i_min && j == j_max {
			for ; i <= i_max; i++ {
				dp[i][j] = k
				k++
			}
			i = i_max
			j_max--
			j = j_max
		}
		if j == j_max && i == i_max {
			for ; j >= j_min; j-- {
				dp[i][j] = k
				k++
			}
			j = j_min
			i_max--
			i = i_max
		}
		if i == i_max && j == j_min {
			for ; i >= i_min; i-- {
				dp[i][j] = k
				k++
			}
			i = i_min
			j_min++
			j = j_min
		}
	}
	return dp
}

// @lc code=end
