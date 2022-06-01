package leetcodego

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	index, dirction := 0, true
	matrix := make([][]byte, numRows)
	for i := range matrix {
		matrix[i] = make([]byte, 0)
	}
	for i := 0; i < len(s); i++ {
		if index == numRows {
			index -= 2
			dirction = !dirction
		}
		if index == -1 {
			index += 2
			dirction = !dirction
		}
		if dirction {
			matrix[index] = append(matrix[index], s[i])
			index++
		} else {
			matrix[index] = append(matrix[index], s[i])
			index--
		}
	}

	ans := ""

	for i := range matrix {
		ans += string(matrix[i])
	}

	return ans
}

// @lc code=end
