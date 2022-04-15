leetcodego
/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)
	nmap := make(map[int] bool)
	for _, v := range nums1 {
		if _, ok := nmap[v]; ok {
			continue	
		}
		nmap[v] = false
	}
	for _, v := range nums2 {
		if _, ok := nmap[v]; ok {
			if !nmap[v]{
				ans = append(ans, v)
				nmap[v] = true
			}
		}
	}
	return ans
}

// @lc code=end

