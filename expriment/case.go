package main

// ListNode is
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	beforeHead := new(ListNode)
	afterHead := new(ListNode)
	before := beforeHead
	after := afterHead
	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	before.Next = afterHead.Next
	after.Next = nil
	return beforeHead.Next
}
func main() {
	first := new(ListNode)
	second := new(ListNode)
	tail := new(ListNode)
	first.Val = 4
	second.Val = 3
	tail.Val = 2
	first.Next = second
	second.Next = tail
	tail.Next = nil
	partition(first, 3)
}
