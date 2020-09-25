package main

func lengthOfLIS(nums []int) int {
	var length = 1
	var ans = 1
	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > tmp {
				length++
				tmp = nums[j]
			}
		}
		ans = max(ans, length)
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	lengthOfLIS(nums)
}
