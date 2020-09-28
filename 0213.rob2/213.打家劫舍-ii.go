package problem0213

/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// n := len(nums)
	// dp := make([]int, n)
	// dp1 := make([]int, n)
	ans1, ans2 := 0, 0
	// for i := 0; i < n-1; i++ {
	// 	if i < 2 {
	// 		dp[i] = nums[i]
	// 	} else {
	// 		for j := 0; j < i-1; j++ {
	// 			dp[i] = max(dp[j]+nums[i], dp[i])
	// 		}
	// 	}
	// 	ans1 = max(ans1, dp[i])
	// }
	// for i := 1; i < n; i++ {
	// 	if i < 3 {
	// 		dp1[i] = nums[i]
	// 	} else {
	// 		for j := 1; j < i-1; j++ {
	// 			dp1[i] = max(dp1[j]+nums[i], dp1[i])
	// 		}
	// 	}
	// 	ans2 = max(ans2, dp1[i])
	// }
	ans1 = rob1(nums[1:])
	ans2 = rob1(nums[:len(nums)-1])
	return max(ans1, ans2)
}

func rob1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		if i < 2 {
			dp[i] = nums[i]
		} else {
			for j := 0; j < i-1; j++ {
				dp[i] = max(dp[j]+nums[i], dp[i])
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
