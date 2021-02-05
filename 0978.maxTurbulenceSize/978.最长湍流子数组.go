package problem0978

/*
 * @lc app=leetcode.cn id=978 lang=golang
 *
 * [978] 最长湍流子数组
 */

// @lc code=start
func maxTurbulenceSize(arr []int) int {
	if len(arr) == 0 || len(arr) == 1 {
		return len(arr)
	}
	if len(arr) == 2 {
		if arr[1] == arr[0] {
			return 1
		}
		return 2
	}
	ans := 0
	dp := make([]int, len(arr))
	dp[0] = 1
	dp[1] = 2
	if arr[1] == arr[0] {
		dp[1] = 1
	}
	for i := 2; i < len(arr); i++ {
		if (arr[i] > arr[i-1]) && (arr[i-2] > arr[i-1]) || (arr[i] < arr[i-1]) && (arr[i-2] < arr[i-1]) {
			dp[i] = dp[i-1] + 1
		} else if arr[i] == arr[i-1] {
			dp[i] = 1
		} else {
			dp[i] = 2
		}
		ans = max(ans, dp[i])
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
