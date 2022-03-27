package leetcodego

/*
 * @lc app=leetcode.cn id=1035 lang=golang
 *
 * [1035] 不相交的线
 */

// @lc code=start
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1)+1, len(nums2)+1
	dp := make([][]int, n1)
	for i := range dp {
		dp[i] = make([]int, n2)
	}

	for i, c := range nums1 {
		for j, v := range nums2 {
			if c == v {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[n1-1][n2-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
