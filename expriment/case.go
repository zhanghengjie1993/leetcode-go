package main

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	j := 0
	var dp = make([]int, n)
	max := 0
	dp[0] = triangle[0][0]
	for i := 1; i < n; i++ {
		if triangle[i][j] < triangle[i][j+1] && j+1 < len(triangle[i]) {
			max = triangle[i][j+1]
			j = j + 1
		} else {
			max = triangle[i][j]
		}
		dp[i] = dp[i-1] + max
	}
	return dp[n-1]
}

func main() {
	var triangle = make([][]int, 4)
	triangle[0] = []int{2}
	triangle[1] = []int{3, 4}
	triangle[2] = []int{6, 5, 7}
	triangle[3] = []int{4, 1, 8, 3}
	minimumTotal(triangle)
}
