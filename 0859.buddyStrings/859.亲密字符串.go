package problem0859

/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 */

// @lc code=start
func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	smap := make(map[byte]bool)
	if A == B {
		for i := 0; i < len(A); i++ {
			if _, ok := smap[A[i]]; !ok {
				smap[A[i]] = true
				continue
			} else {
				return true
			}
		}
		return false
	}
	first, second := -1, -1
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}
	return (second != -1) && (A[first] == B[second]) && (A[second] == B[first])

}

// @lc code=end
