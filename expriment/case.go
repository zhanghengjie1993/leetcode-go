package main

func climbStairs(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 2
	}
	var dp = make([]int, n+1)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	climbStairs(2)
}
