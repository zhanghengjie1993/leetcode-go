package problem4

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	nums3 := make([]int, n1+n2)
	i, j, k := 0, 0, 0
	for ; k < n1+n2; k++ {
		if i >= n1 {
			nums3[k] = nums2[j]
			j++
			continue
		}
		if j >= n2 {
			nums3[k] = nums1[i]
			i++
			continue
		}
		if nums1[i] <= nums2[j] {
			nums3[k] = nums1[i]
			i++
		} else {
			nums3[k] = nums2[j]
			j++
		}
	}
	len := n1 + n2
	if len%2 == 0 {
		return float64((nums3[(len/2)-1] + nums3[(len/2)])) / 2
	}
	return float64(nums3[((len+1)/2)-1])
}

// @lc code=end
