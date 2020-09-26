package main

import "fmt"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	tmp := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				tmp = dp[j] + 1
			} else {
				tmp = 1
			}
			dp[i] = max(dp[i], tmp)
		}
	}
	return dp[len(nums)-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	ans := lengthOfLIS(nums)
	fmt.Println(ans)
}
