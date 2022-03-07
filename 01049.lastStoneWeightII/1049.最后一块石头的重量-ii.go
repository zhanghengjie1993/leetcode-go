package laststoneweightii

/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	dp := make([]int, (sum+1)/2+1)

	for i := 0; i < len(stones); i++ {
		for j := (sum + 1) / 2; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return abs(sum - 2*dp[(sum+1)/2])
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
