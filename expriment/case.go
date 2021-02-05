package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	dp := make([]int, 1)
	for i := 0; i < len(s); i++ {
		for _, v := range wordDict {
			for j := range dp {
				fmt.Println("program is runing")
				if s[dp[j]:i+1] == v {
					dp = append(dp, i+1)
				}
			}
		}
	}

	return dp[len(dp)-1] == len(s)
}

func main() {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	ans := wordBreak(s, wordDict)
	fmt.Printf("the ans is:%t\n", ans)
}
