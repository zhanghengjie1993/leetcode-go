package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	j := 0

	for i := 0; i < len(s); i++ {

		if j >= len(p) {
			return false
		}

		if (p[j] >= 65 && p[j] <= 90) || (p[j] >= 97 && p[j] <= 122) {
			if p[j] != s[i] {
				return false
			}
			j++
			continue
		}
	}
	return true
}
func main() {

	ans := isMatch("aa", "a")
	fmt.Printf("the ans is:%t\n", ans)
}
