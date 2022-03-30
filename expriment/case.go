package main

type State int
type CharType int

const (
	STATE_INITIAL State = iota
	STATE_INT_SIGN
	STATE_INTEGER
	STATE_POINT
	STATE_POINT_WITHOUT_INT
	STATE_FRACTION
	STATE_EXP
	STATE_EXP_SIGN
	STATE_EXP_NUMBER
	STATE_END
)

const (
	CHAR_NUMBER CharType = iota
	CHAR_EXP
	CHAR_POINT
	CHAR_SIGN
	CHAR_SPACE
	CHAR_ILLEGAL
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER
	case 'e', 'E':
		return CHAR_EXP
	case '.':
		return CHAR_POINT
	case '+', '-':
		return CHAR_SIGN
	case ' ':
		return CHAR_SPACE
	default:
		return CHAR_ILLEGAL
	}
}

func isNumber(s string) bool {
	transfer := map[State]map[CharType]State{
		STATE_INITIAL: {
			CHAR_SPACE:  STATE_INITIAL,
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
			CHAR_SIGN:   STATE_INT_SIGN,
		},
		STATE_INT_SIGN: {
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
		},
		STATE_INTEGER: {
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_EXP:    STATE_EXP,
			CHAR_POINT:  STATE_POINT,
			CHAR_SPACE:  STATE_END,
		},
		STATE_POINT: {
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
			CHAR_SPACE:  STATE_END,
		},
		STATE_POINT_WITHOUT_INT: {
			CHAR_NUMBER: STATE_FRACTION,
		},
		STATE_FRACTION: {
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
			CHAR_SPACE:  STATE_END,
		},
		STATE_EXP: {
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SIGN:   STATE_EXP_SIGN,
		},
		STATE_EXP_SIGN: {
			CHAR_NUMBER: STATE_EXP_NUMBER,
		},
		STATE_EXP_NUMBER: {
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SPACE:  STATE_END,
		},
		STATE_END: {
			CHAR_SPACE: STATE_END,
		},
	}
	state := STATE_INITIAL
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = transfer[state][typ]
		}
	}
	return state == STATE_INTEGER || state == STATE_POINT || state == STATE_FRACTION || state == STATE_EXP_NUMBER || state == STATE_END
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	Length int
	Head   *ListNode
	Tail   *ListNode
}

func Constructor() MyLinkedList {

	return MyLinkedList{0, nil, nil}
}

func (this *MyLinkedList) Get(index int) int {
	if index > this.Length-1 {
		return -1
	}
	temp := &ListNode{}
	temp.Next = this.Head
	for ; index >= 0; index-- {
		temp = temp.Next
	}
	return temp.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{val, nil}
	node.Next = this.Head
	this.Head = node
	this.Length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &ListNode{val, nil}
	if this.Tail != nil {
		this.Tail.Next = node
	} else {
		this.Head.Next = node
	}
	this.Tail = node
	this.Length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	node := &ListNode{val, nil}
	if index <= 0 {
		node.Next = this.Head
		this.Head = node
		this.Length++
	} else if index == this.Length {
		this.Tail.Next = node
		this.Tail = node
		this.Length++
	} else if index < this.Length {
		cur := &ListNode{0, this.Head}
		for ; index > 0; index-- {
			cur = cur.Next
		}
		node.Next = cur.Next
		cur.Next = node
		this.Length++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < this.Length && index >= 0 {
		dummyHead := &ListNode{0, this.Head}
		cur := dummyHead
		for ; index > 0; index-- {
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
		this.Length--
	}
}

func main() {

	linkedList := Constructor()
	linkedList.AddAtHead(7)
	linkedList.AddAtHead(2)
	linkedList.AddAtHead(1)
	// linkedList.AddAtTail(3)
	linkedList.AddAtIndex(3, 0) //链表变为1-> 2-> 3
	linkedList.Get(1)           //返回2
	linkedList.DeleteAtIndex(1) //现在链表是1-> 3
	linkedList.Get(1)
}
