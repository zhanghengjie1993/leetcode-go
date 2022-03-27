package leetcodego

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {

	sum := 0
	max_sum := nums[0]

	for _, v := range nums {
		sum += v
		if sum < v {
			sum = v
		}
		max_sum = max(max_sum, sum)
	}
	return max_sum

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
