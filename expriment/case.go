package main

import "fmt"

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
	if arr[1] == arr[0] {
		dp[1] = 1
	}
	dp[1] = 2
	for i := 2; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			dp[i] = 1
		}
		if (arr[i] > arr[i-1]) && (arr[i-2] > arr[i-1]) || (arr[i] < arr[i-1]) && (arr[i-2] < arr[i-1]) {
			dp[i] = dp[i-1] + 1
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
func main() {
	arr := []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	ans := maxTurbulenceSize(arr)
	fmt.Printf("the ans is:%d\n", ans)
}
