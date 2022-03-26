package leetcodego

/*
 * @lc app=leetcode.cn id=718 lang=golang
 *
 * [718] 最长重复子数组
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)

	dp := make([][]int, n1)
	for i := range dp {
		dp[i] = make([]int, n2)
	}
	ans := 0
	for i := 0; i < n2; i++ {
		if nums1[0] == nums2[i] {
			dp[0][i] = 1
		}
		ans = max(ans, dp[0][i])
	}
	for i := 0; i < n1; i++ {
		if nums1[i] == nums2[0] {
			dp[i][0] = 1
		}
		ans = max(ans, dp[i][0])
	}
	for i := 1; i < n1; i++ {
		for j := 1; j < n2; j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
