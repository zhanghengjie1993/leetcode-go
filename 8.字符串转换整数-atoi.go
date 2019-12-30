package leetcode

import (
	"strconv"
	"strings"
)

const INT_MAX = 2 ^ 31 - 1
const INT_MIN = -2 ^ 31

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(str string) int {
	str = strings.TrimSpace(str)

	res, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return res
}

// @lc code=end
