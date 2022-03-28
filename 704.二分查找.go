package leetcodego

/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {

	return bins(nums, 0, len(nums)-1, target)

}

func bins(nums []int, l int, r int, target int) int {

	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return bins(nums, l, mid-1, target)
	} else {
		return bins(nums, mid+1, r, target)
	}

}

// @lc code=end
