package list

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode { return &ListNode{Val: val} }

func (l *ListNode) PushBack(val int) {
	if l == nil {
		l = &ListNode{Val: val}
	}
	p := l
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &ListNode{Val: val}
}

func (l *ListNode) PopFront() int {
	if l == nil {
		return 0
	}
	d := l.Val
	if l.Next != nil {
		l.Val = l.Next.Val
		l.Next = l.Next.Next
	}
	return d
}
