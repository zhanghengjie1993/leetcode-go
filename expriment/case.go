package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	flag := true
	n, m := len(haystack), len(needle)
	for i := 0; i < n; i++ {
		if n-i < m {
			return -1
		}
		for j := 0; j < m; j++ {
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
			if j == m-1 {
				flag = true
			}
		}
		if flag {
			return i
		}
	}
	return -1
}

func main() {
	strStr("mississippi", "issipi")
}
