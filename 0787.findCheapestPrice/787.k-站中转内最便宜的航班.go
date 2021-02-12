package problem0787

import "math"

/*
 * @lc app=leetcode.cn id=787 lang=golang
 *
 * [787] K 站中转内最便宜的航班
 */

// @lc code=start
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	ans := math.MaxInt32
	hmap := make(map[int][][]int)
	for _, v := range flights {
		hmap[v[0]] = append(hmap[v[0]], v)
	}

	queue := make([][]int, 0)
	queue = append(queue, []int{src, 0})

	for K >= -1 {
		l := len(queue)
		for l > 0 {
			flight := queue[0]
			queue = queue[1:]
			if flight[0] == dst {
				ans = min(ans, flight[1])
				l--
				continue
			}

			for _, v := range hmap[flight[0]] {
				if flight[1] > ans {
					continue
				}
				queue = append(queue, []int{v[1], flight[1] + v[2]})
			}
			l--
		}
		K--
	}
	if ans == math.MaxInt32 {
		ans = -1
	}
	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end
