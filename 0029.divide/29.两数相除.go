package divide

import "math"

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	ans, flag := 0, 0
	factor := 1
	if dividend^divisor < 0 {
		factor = -1
	}
	dividend = abs(dividend)
	divisor = abs(divisor)
	if divisor == 1 {
		return formatAns(dividend * factor)
	}
	for i := getLength(dividend); i >= 0; i-- {
		if dividend>>i == divisor {
			ans = (ans<<1 | 1) << i
			return formatAns(ans * factor)
		}

		if dividend>>i > divisor {
			flag = 1
			high := (dividend>>i - divisor) << i
			low := getLowBit(dividend, i)
			dividend = high | low
		}
		ans = ans<<1 | flag
	}
	return formatAns(ans * factor)
}

func getLowBit(num int, n int) int {
	ans, flag := 0, 0

	for i := 0; i < n; i++ {
		flag = num & 1
		ans = ans<<1 | flag
	}

	return ans
}

func getLength(num int) int {
	var i int = 0
	for num != 0 {
		num = num >> 1
		i++
	}
	return i
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func formatAns(num int) int {
	if num > math.MaxInt32 {
		return math.MaxInt32
	}
	return num
}

// @lc code=end
