package problem0763

/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
func partitionLabels(S string) []int {
	var ans []int
	for i := 0; i < len(S); i++ {
		var index = 0
		for j := i + 1; j < len(S); j++ {
			if S[i] == S[j] {
				index = j
			}
		}
		if index == 0 {
			ans = append(ans, 1)
		} else {
			flag, postion := hasNext(S, i+1, index)
			if flag {
				i = 
			}
		}
	}
	return ans
}

func hasNext(S string, left int, right int) (bool, int) {
	var maxIndex int = 0
	for i := left; i < right; i++ {
		for j := i + 1; j < len(S); j++ {
			index := 0
			if S[i] == S[j] {
				index = j
			}
			if index > right {
				maxIndex = max(index, maxIndex)
			}
		}
	}
	if maxIndex != 0 {
		return true, maxIndex
	}
	return false, maxIndex
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
