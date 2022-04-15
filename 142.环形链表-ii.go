leetcodego
/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	index := 0
	lmap := make(map[*ListNode][int])
    for cur := head; cur != nil; {
		if _, ok := lmap[cur]; ok {
			return lmap[cur]
		}else{
			lmap[cur] = index
		}
		index++
		cur = cur.Next
	}
	return -1
}
// @lc code=end

