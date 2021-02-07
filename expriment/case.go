package main

import (
	"fmt"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	ans := math.MaxInt32
	hmap := make(map[int][][]int)
	for _, v := range flights {
		v = append(v, 0)
		hmap[v[0]] = append(hmap[v[0]], v)
	}

	queue := make([][]int, 0)
	queue = append(queue, []int{src, 0})

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

func main() {

	n := 4
	edges := [][]int{
		{0, 2, 1},
		{0, 1, 1},
		{1, 3, 1},
		{2, 1, 1},
	}
	src := 0
	dst := 3
	k := 3
	ans := findCheapestPrice(n, edges, src, dst, k)
	fmt.Printf("the ans is:%d\n", ans)
}
