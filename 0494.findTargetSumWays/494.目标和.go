package findtargetsumways

/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum-target < 0 || (sum-target)%2 != 0 {
		return 0
	}
	bagsize := (sum - target) / 2
	dp := make([]int, bagsize+1)
	dp[0] = 1
	for i := range nums {
		for j := bagsize; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagsize]
}

// @lc code=end
