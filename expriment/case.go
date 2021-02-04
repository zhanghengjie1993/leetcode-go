package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if i == 0 && j == 0 {
				dp[0] = 1
				continue
			}
			if j != 0 {
				dp[j] = dp[j-1] + dp[j]
			}
		}
	}
	return dp[n-1]
}

func main() {
	var m1 map[string]string

	m1 = make(map[string]string, 1)

	m1["ds"] = "sd"
	m1["as"] = "ad"

	m2 := map[string]string{
		"zzz": "ffff",
		"ttt": "zzzz",
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	type member struct {
		id     int
		name   string
		email  string
		gender int
	}

	mb1 := member{1, "tom", "dfdf@qq.com", 1}
	mb1.id = 3
	mb2 := member{id: 2, name: "jack"}
	mb2.email = "sqdq@qq.com"

}
