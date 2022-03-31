package leetcodego

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode1, k int) *ListNode1 {
	n := 0
	// lHead := head
	var lTail *ListNode1
	pre := &ListNode1{0, nil}
	cur := head

	for cur != nil {
		lTail = cur
		for n < k && cur.Next != nil {
			nextNode := cur.Next
			cur.Next = pre
			pre = cur
			cur = nextNode
			n++
		}
		lTail.Next = cur
		if n < k {

		}
	}
	return pre

}

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

// @lc code=end
