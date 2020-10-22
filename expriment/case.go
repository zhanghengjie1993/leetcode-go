package main

import "fmt"

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
			postion := hasNext(S, i+1, index)
			ans = append(ans, postion-i+1)
			i = postion
		}
	}
	return ans
}

func hasNext(S string, left int, right int) int {
	var maxIndex int = right
	for i := left; i < right; i++ {
		index := 0
		for j := i + 1; j < len(S); j++ {
			if S[i] == S[j] {
				index = j
			}
			maxIndex = max(index, maxIndex)
		}
	}
	return maxIndex
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	S := "ababcbacadefegdehijhklij"
	ans := partitionLabels(S)
	fmt.Println(ans)
}
