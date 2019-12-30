package problem0006

import "bytes"

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
	buffers := make([]bytes.Buffer, numRows)
	var res bytes.Buffer
	flag := true
	index := 0
	for i := 0; i < len(s); i++ {
		if flag == true {
			if index == numRows-1 {
				buffers[index].WriteByte(s[i])
				index--
				flag = false
				continue
			}
			buffers[index].WriteByte(s[i])
			index++
		} else {
			if index == 0 {
				buffers[index].WriteByte(s[i])
				index++
				flag = true
				continue
			}
			buffers[index].WriteByte(s[i])
			index--
		}
	}
	for i := 0; i < numRows; i++ {
		res.WriteString(buffers[i].String())
	}
	return res.String()
}

// @lc code=end
