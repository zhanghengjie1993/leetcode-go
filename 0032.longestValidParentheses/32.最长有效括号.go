package problem0032

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	ans := 0
	stack := []byte{}
	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			stack = append(stack, s[i])
		}

		if len(stack) != 0 && stack[len(stack)-1] == '(' && s[i] == ')' {
			ans = ans + 2
			stack = stack[:len(stack)-1]
		}

	}
	return ans
}

// @lc code=end
