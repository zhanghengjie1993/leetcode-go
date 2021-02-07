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
		v = append(v, 0)
		hmap[v[0]] = append(hmap[v[0]], v)
	}

	queue := make([][]int, 0)
	queue = append(queue, hmap[src]...)

	for K >= 0 {
		l := len(queue)
		for l > 0 {
			flight := queue[0]
			queue = queue[1:]
			if flight[1] == dst {
				ans = min(ans, flight[2]+flight[3])
				l--
				continue
			}

			for _, v := range hmap[flight[1]] {
				v[3] = flight[2] + v[3]
			}
			queue = append(queue, hmap[flight[1]]...)
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
