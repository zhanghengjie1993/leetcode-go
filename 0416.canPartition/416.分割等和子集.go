package problem0416

/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	var dp []int = make([]int, sum/2+1)

	for i := range nums {

		for j := sum / 2; j >= nums[i]; j-- {

			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	for _, v := range dp {
		if v == sum/2 {
			return true
		}
	}
	return false

}

func max(x, y int) int {

	if x >= y {
		return x
	}
	return y
}

// @lc code=end
