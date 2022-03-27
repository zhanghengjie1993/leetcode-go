package leetcodego

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {

	if len(s) == 0 {
		return true
	}

	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			if i == len(s)-1 {
				return true
			}
			i++
			j++
		} else {
			j++
		}
	}
	return false
}

// @lc code=end
